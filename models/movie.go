package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/dhruvvadoliya1/movie-app-backend/pkg/structs"
	"github.com/dhruvvadoliya1/movie-app-backend/utils"
	"github.com/doug-martin/goqu/v9"
)

// Movie model represents the movie data
type Movie struct {
	Id   		int64        `db:"id" json:"id"`
	Title 		string       ` db:"title" json:"title"`
	Genre     	string       `db:"genre" json:"genre"`
	Year        int       `db:"year" json:"year"`
	Rating     	float32       `csv:"rating" db:"rating" json:"rating"`
	CreatedAt   sql.NullTime `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at,omitempty" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at,omitempty" json:"deleted_at"`
}

type ResCount struct {
	Count int64
	Err   error
}

const MovieTable = "movies"

type MovieModel struct {
	db *goqu.Database
}

type ChanMovies struct {
	Movies []Movie
	Err  error
}

func InitMovieModel(goqu *goqu.Database) (MovieModel, error) {
	return MovieModel{
		db: goqu,
	}, nil
}

func (model *MovieModel) InsertFromMap(data map[string]string) error {

	movie, err := model.MapToStruct(data)
	if err != nil {
		return err
	}

	_, err = model.InsertMovie(movie)
	return err
}

func (model *MovieModel) MapToStruct(data map[string]string) (Movie, error) {
	movie := Movie{}

	if value, ok := data["id"]; ok && value != "" {
		movieId, err := utils.ParseInt64(value)
		if err != nil {
			return movie, fmt.Errorf("error parsing job_id: %v", err)
		}
		movie.Id = movieId
	}

	if value, ok := data["title"]; ok && value != "" {
		movie.Title = value
	}

	if value, ok := data["genre"]; ok && value != "" {
		movie.Genre = value
	}

	if value, ok := data["year"]; ok && value != "" {
		yearInt, err := utils.ParseInt64(value)
		if err != nil {
			return movie, fmt.Errorf("error parsing year: %v", err)
		}
		movie.Year = int(yearInt)
	}

	if value, ok := data["rating"]; ok && value != "" {
		ratingFloat, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return movie, fmt.Errorf("error parsing rating: %v", err)
		}
		movie.Rating = float32(ratingFloat)
	}
	
	return movie, nil
}

func (model *MovieModel) InsertMovie(movie Movie) (Movie, error) {

	_, err := model.GetByTitle(movie.Title)
	if err == nil {
		// Duplicate found
		return movie, fmt.Errorf("movie with title '%s' already exists", movie.Title)
	} else if err != sql.ErrNoRows {
		// Some other error occurred
		return movie, err
	}

	record := goqu.Record{
		"title": 		movie.Title,
		"genre":      	movie.Genre,
		"year":         movie.Year,
		"rating":     	movie.Rating,
	}

	_, err = model.db.Insert(MovieTable).
		Rows(record).
		Returning("id", "title", "genre", "year", "rating").
		Executor().
		ScanStruct(&movie)
	return movie, err
}

// GetMovieByID retrieves a movie by its ID
func (model *MovieModel) GetById(id int64) (Movie, error) {
	var movie Movie
	found, err := model.db.From(MovieTable).Where(goqu.Ex{
		"id": id,
		"deleted_at": nil,
	}).ScanStruct(&movie)

	if err != nil {
		return movie, err
	}

	if !found {
		return movie, sql.ErrNoRows
	}

	return movie, nil
}

func (model *MovieModel) GetByTitle(title string) (Movie, error) {
	var movie Movie
	found, err := model.db.From(MovieTable).Where(
		goqu.L("LOWER(title) = LOWER(?)", title),
		goqu.C("deleted_at").IsNull(),
	).ScanStruct(&movie)

	if err != nil {
		return movie, err
	}

	if !found {
		return movie, sql.ErrNoRows
	}

	return movie, nil
}

func (model *MovieModel) getMoviesQuery(param structs.ReqGetMovieList, query *goqu.SelectDataset) (*goqu.SelectDataset) {
	
	if param.Title != ""{
		query = query.Where(
			goqu.L("LOWER(title) LIKE LOWER(?)", "%"+param.Title+"%"),
		)
	}

	if param.Genre != ""{
		query = query.Where(
			goqu.L("LOWER(genre) LIKE LOWER(?)", "%"+param.Genre+"%"),
		)
	}

	if param.Year != 0 {
		query = query.Where(
			goqu.C("year").Eq(param.Year),
		)
	}

	return query
}

func (model *MovieModel) GetMovies(param structs.ReqGetMovieList) ([]Movie, error) {

	var movies []Movie
	
	offset := uint((param.Page - 1) * param.Limit)
	query := model.db.From(MovieTable).Offset(offset).Limit(uint(param.Limit)).Where(goqu.C("deleted_at").IsNull())
	query = model.getMoviesQuery(param, query)

	if err := query.ScanStructs(&movies); err != nil {
		return nil, err
	}

	// if len(movies) == 0 {
	// 	return movies, sql.ErrNoRows
	// }

	return movies, nil

}

func (model *MovieModel) GetMovieFromChan(param structs.ReqGetMovieList, channel chan ChanMovies) {
	jobs, err := model.GetMovies(param)
	channel <- ChanMovies{Movies: jobs, Err: err}
}

func (model *MovieModel) GetMovieCount(param structs.ReqGetMovieList) (int64, error) {

	query := model.db.From(MovieTable).Where(goqu.C("deleted_at").IsNull())
	query = model.getMoviesQuery(param,query)

	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (model *MovieModel) GetMovieCountFromChan(param structs.ReqGetMovieList, channel chan ResCount) {
	count, err := model.GetMovieCount(param)
	channel <- ResCount{Count: count, Err: err}
}

func (model *MovieModel) UpdateMovie(movie Movie) (Movie, error) {

	record := goqu.Record{}

	if movie.Title != "" {
		record["title"] = movie.Title
	}

	if movie.Genre != "" {
		record["genre"] = movie.Genre
	}
	
	if movie.Year != 0 {
		record["year"] = movie.Year
	}

	if movie.Rating != 0 {
		record["rating"] = movie.Rating
	}

	found, err := model.db.Update(MovieTable).
		Set(record).Where(goqu.Ex{
		"id": movie.Id,
	}, goqu.C("deleted_at").IsNull()).
		Returning("id", "title", "genre", "year", "rating").
		Executor().
		ScanStruct(&movie)

	if !found {
		return movie, sql.ErrNoRows
	}

	return movie, err
}

func (model *MovieModel) DeleteMovie(id int64) (int64, error) {

	record := goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	found, err := model.db.Update(MovieTable).
		Set(record).
		Where(goqu.C("id").Eq(id),
			goqu.C("deleted_at").IsNull()).
		Returning("id").
		Executor().
		ScanVal(&id)

	if err != nil {
		fmt.Println("error while delete===>", err)
		return id, err
	}
	if !found {
		return id, sql.ErrNoRows
	}
	return id, err
}


