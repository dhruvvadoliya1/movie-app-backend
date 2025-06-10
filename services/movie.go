package services

import (
	"github.com/dhruvvadoliya1/movie-app-backend/models"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/structs"
)

type MovieService struct {
	movieModel *models.MovieModel
}

func NewMovieServices(movieModel *models.MovieModel) *MovieService {
	return &MovieService{
		movieModel: movieModel,
	}
}

func (movieSrv *MovieService) CreateMovie(movie models.Movie) (models.Movie, error) {
	movie, err := movieSrv.movieModel.InsertMovie(movie)
	return movie, err
}

func (movieSrv *MovieService) UpdateMovie(movie models.Movie) (models.Movie, error) {
	movie, err := movieSrv.movieModel.UpdateMovie(movie)
	return movie, err
}

func (movieSrv *MovieService) GetMovie(movieId int64) (models.Movie, error) {
	movie, err := movieSrv.movieModel.GetById(movieId)
	return movie, err
}

func (movieSvc *MovieService) GetMovies(params structs.ReqGetMovieList) ([]models.Movie, int64, error) {

	chanMovies := make(chan models.ChanMovies)
	chanCount := make(chan models.ResCount)

	go func() {
		movieSvc.movieModel.GetMovieFromChan(params, chanMovies)
		close(chanMovies)
	}()

	go func() {
		movieSvc.movieModel.GetMovieCountFromChan(params, chanCount)
		close(chanCount)
	}()

	movieResult := <-chanMovies
	if movieResult.Err != nil {
		return nil, 0, movieResult.Err
	}

	countResult := <-chanCount
	if countResult.Err != nil {
		return nil, 0, countResult.Err
	}

	return movieResult.Movies, countResult.Count, nil
}

func (movieSrv *MovieService) DeleteMovie(movieId int64) (int64, error) {
	movieId, err := movieSrv.movieModel.DeleteMovie(movieId)
	return movieId, err
}
