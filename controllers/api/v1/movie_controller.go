package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhruvvadoliya1/movie-app-backend/constants"
	"github.com/dhruvvadoliya1/movie-app-backend/models"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/structs"
	"github.com/dhruvvadoliya1/movie-app-backend/services"
	"github.com/dhruvvadoliya1/movie-app-backend/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

type MovieController struct {
	movieService     *services.MovieService
	logger         *zap.Logger
}

func NewMovieController(goqu *goqu.Database, logger *zap.Logger) (*MovieController, error) {
	movieModel, err := models.InitMovieModel(goqu)
	if err != nil {
		return nil, err
	}
	movieSvc := services.NewMovieServices(&movieModel)

	return &MovieController{
		movieService:     movieSvc,
		logger:         logger,
	}, nil

}

// swagger:route POST /movies Movies RequestCreateMovie
//
// Create a Movie.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  201: ResponseCreateMovie
//	   400: GenericResFailBadRequest
//		  500: GenericResError

func (ctrl *MovieController) CreateMovie(c *fiber.Ctx) error {
	var movieReq structs.ReqCreateMovie

	err := json.Unmarshal(c.Body(), &movieReq)

	if err != nil {
		ctrl.logger.Error("error unmarshaling request body", zap.Error(err))
		return utils.JSONError(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	validate.RegisterValidation("year_lt_now", structs.YearLessThanNow)
	err = validate.Struct(movieReq)
	if err != nil {
		ctrl.logger.Error("error while validating body", zap.Any("body", movieReq), zap.Error(err))
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}


	movie, err := ctrl.movieService.CreateMovie(
		models.Movie{
			Title:              movieReq.Title,
			Genre: movieReq.Genre,
			Year: movieReq.Year,
			Rating: float32(movieReq.Rating),
		})

	if err != nil {
		return utils.JSONFail(c, http.StatusInternalServerError, err.Error())
	}

	//preparing response movie
	resMovie := structs.ResMovie{
		Id:       	movie.Id,
		Title:      movie.Title,
		Genre: 		movie.Genre,
		Year: 		movie.Year,
		Rating: 	float32(movie.Rating),
	}

	return utils.JSONSuccess(c, http.StatusCreated, resMovie)
}

// swagger:route GET /movies/{movieId} Movies RequestGetMovie
//
// Get a Movie Details.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  201: ResponseGetMovie
//	   400: GenericResFailBadRequest
//		  500: GenericResError

func (ctrl *MovieController) GetMovie(c *fiber.Ctx) error {

	movieID, err := utils.ParseInt64(c.Params(constants.ParamMid))
	if err != nil {
		ctrl.logger.Error("error while parsing id", zap.Any("id", movieID), zap.Error(err))
		return utils.JSONError(c, http.StatusBadRequest, constants.InvalidMovieId)
	}
	movie, err := ctrl.movieService.GetMovie(movieID)

	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.MovieNotExist)
		}
		ctrl.logger.Error("error while get movie by id", zap.Any("id", movieID), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetMovie)
	}

	//preparing response movie
	resMovie := structs.ResMovie{
		Id:       	movie.Id,
		Title:      movie.Title,
		Genre: 		movie.Genre,
		Year: 		movie.Year,
		Rating: 	float32(movie.Rating),
	}

	return utils.JSONSuccess(c, http.StatusOK, resMovie)

}

// swagger:route GET /movies Movies RequestGetMovieList
//
// Get a Movie List.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  201: ResponseGetMovieList
//	   400: GenericResFailBadRequest
//		  500: GenericResError

func (ctrl *MovieController) GetMovies(c *fiber.Ctx) error {
	query := structs.ReqGetMovieList{}
	
	page, err := strconv.ParseUint(c.Query("page", "1"), 10, 64)
	if err != nil {
		return err
	}

	if page < 1 {
		page = 1
	}

	query.Page = int64(page)

	limit, err := strconv.ParseUint(c.Query("limit", "10"), 10, 64)
	if err != nil {
		return err
	}

	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	query.Limit = int64(limit)
	// offset := ((pageInt - 1) * limitInt)

	query.Title = c.Query("title", "")
	query.Genre = c.Query("genre", "")
	query.Year = c.QueryInt("year", 0)

	movies, count, err := ctrl.movieService.GetMovies(query)

	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.MovieNotExist)
		}
		ctrl.logger.Error("error while get movie list", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetMovie)
	}

	movielist := structs.ResGetMovieList{}

	resMovies := make([]structs.ResMovie, len(movies))

	for i, movie := range movies {
		resMovies[i] = structs.ResMovie{
			Id:       	movie.Id,
			Title:      movie.Title,
			Genre: 		movie.Genre,
			Year: 		movie.Year,
			Rating: 	float32(movie.Rating),
		}

	}

	movielist.Data.Movies = resMovies
	movielist.Data.Meta.Count = count

	return utils.JSONSuccess(c, http.StatusOK, movielist.Data)

}

// swagger:route PUT /movies/{movieId} Movies RequestUpdateMovie
//
// Update a Movie.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  200: ResponseUpdateMovie
//	   400: GenericResFailBadRequest
//		  500: GenericResError

func (ctrl *MovieController) UpdateMovie(c *fiber.Ctx) error {
	var movieReq structs.ReqUpdateMovie

	movieId := c.Params(constants.ParamMid)
	MovieIdInt, err := strconv.ParseInt(movieId, 10, 64)

	if err != nil {
		ctrl.logger.Error("error while parsing movie Id", zap.Any("movie Id", movieId), zap.Error(err))
		return utils.JSONFail(c, http.StatusBadRequest, constants.InvalidMovieId)
	}

	err = json.Unmarshal(c.Body(), &movieReq)

	if err != nil {
		ctrl.logger.Error("error while parsing body", zap.Any("body", movieReq), zap.Error(err))
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	validate.RegisterValidation("year_lt_now", structs.YearLessThanNow)
	err = validate.Struct(movieReq)
	if err != nil {
		ctrl.logger.Error("error while validating body", zap.Any("body", movieReq), zap.Error(err))
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	_, err = ctrl.movieService.GetMovie(MovieIdInt)

	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.MovieNotExist)
		}
		ctrl.logger.Error("error while get movie by id", zap.Any("body", movieReq), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetMovie)

	}

	movie, err := ctrl.movieService.UpdateMovie(
		models.Movie{
			Id:       	MovieIdInt,
		Title:      movieReq.Title,
		Genre: 		movieReq.Genre,
		Year: 		movieReq.Year,
		Rating: 	float32(movieReq.Rating),
		})

	if err != nil {
		return utils.JSONFail(c, http.StatusInternalServerError, err.Error())
	}

	//preparing response
	resMovie := structs.ResMovie{
		Id:       	movie.Id,
		Title:      movie.Title,
		Genre: 		movie.Genre,
		Year: 		movie.Year,
		Rating: 	float32(movie.Rating),
	}

	return utils.JSONSuccess(c, http.StatusOK, resMovie)
}

// swagger:route DELETE /movies/{movieId} Movies RequestDeleteMovie
//
// Delete Movie Details.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  200: ResponseDeleteMovie
//	   400: GenericResFailBadRequest
//		  500: GenericResError

func (ctrl *MovieController) DeleteMovie(c *fiber.Ctx) error {

	movieID, err := utils.ParseInt64(c.Params(constants.ParamMid))
	if err != nil {
		ctrl.logger.Error("error while parsing id", zap.Any("id", movieID), zap.Error(err))
		return utils.JSONError(c, http.StatusBadRequest, constants.InvalidMovieId)
	}

	movieID, err = ctrl.movieService.DeleteMovie(movieID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.MovieNotExist)
		}
		ctrl.logger.Error("error while delete movie", zap.Any("id", movieID), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrDeleteMovie)
	}

	return utils.JSONSuccess(c, http.StatusOK, movieID)

}
