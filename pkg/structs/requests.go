package structs

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// All request structs
// Request struct have Req prefix

type ReqRegisterUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Roles     string `json:"roles" validate:"required"`
}

type ReqLoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ReqCreateMovie struct {
	Title              string    `json:"title,omitempty" validate:"required,max=100"`
	Genre 			   string    `json:"genre,omitempty" validate:"required,min=3"`
	Year			   int       `json:"year,omitempty" validate:"required,gte=1900,year_lt_now"` //
	Rating 			   float32   `json:"rating" validate:"gte=0,lte=5"`
}

type ReqGetMovieList struct {
	Page  int64  `json:"page" query:"page"`
	Limit int64  `json:"limit" query:"limit"`
	ReqCreateMovie
}

type ReqUpdateMovie struct {
	Title  string  `json:"title,omitempty" validate:"max=100" faker:"sentence"`
	Genre  string  `json:"genre,omitempty" validate:"min=3" faker:"paragraph"`
	Year   int     `json:"year,omitempty" validate:"gte=1900,year_lt_now"` 
	Rating float32 `json:"rating,omitempty" validate:"gte=0,lte=5"`
}

func YearLessThanNow(fl validator.FieldLevel) bool {
	year := fl.Field().Int()
	currentYear := int64(time.Now().Year())
	return year <= currentYear
}