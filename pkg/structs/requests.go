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

type ReqCreateJob struct {
	Title              string    `json:"title,omitempty" validate:"required,min=5,max=100" faker:"sentence"`
	Description        string    `json:"description,omitempty" validate:"required,min=20" faker:"paragraph"`
	MaxSalary          int64     `json:"max_salary,omitempty" validate:"required,gt=0,gtefield=MinSalary" faker:"boundary_start=50000,boundary_end=200000"`
	MinSalary          int64     `json:"min_salary,omitempty" validate:"required,gt=0" faker:"boundary_start=30000,boundary_end=50000"`
	PayPeriod          string    `json:"pay_period,omitempty" validate:"required,oneof=HOURLY DAILY BIWEEKLY WEEKLY MONTHLY YEARLY" faker:"oneof:HOURLY,DAILY,BIWEEKLY,WEEKLY,MONTHLY,YEARLY"`
	Location           string    `json:"location,omitempty" validate:"required" faker:"word"`
	CompanyId          int64     `json:"company_id,omitempty" validate:"required,gt=0" faker:"boundary_start=1,boundary_end=10000"`
	Views              int64     `json:"views,omitempty" validate:"gte=0" faker:"boundary_start=0,boundary_end=10000"`
	ClosedTime         time.Time `json:"closed_time,omitempty" validate:"omitempty,gtfield=OriginalListedTime" `
	ExpiryTime         time.Time `json:"expiry_time,omitempty" validate:"omitempty,gtfield=OriginalListedTime"`
	OriginalListedTime time.Time `json:"original_listed_time,omitempty" validate:"omitempty"`
}

type ReqUpdateJob struct {
	Title              string    `json:"title,omitempty" validate:"omitempty,min=5,max=100" faker:"sentence"`
	Description        string    `json:"description,omitempty" validate:"omitempty,min=20" faker:"paragraph"`
	MaxSalary          int64     `json:"max_salary,omitempty" validate:"omitempty,gt=0,gtefield=MinSalary" faker:"boundary_start=50000,boundary_end=200000"`
	MinSalary          int64     `json:"min_salary,omitempty" validate:"omitempty,gt=0" faker:"boundary_start=30000,boundary_end=50000"`
	PayPeriod          string    `json:"pay_period,omitempty" validate:"omitempty,oneof=HOURLY DAILY BIWEEKLY WEEKLY MONTHLY YEARLY" faker:"oneof:HOURLY,DAILY,BIWEEKLY,WEEKLY,MONTHLY,YEARLY"`
	Location           string    `json:"location,omitempty" validate:"omitempty" faker:"word"`
	CompanyId          int64     `json:"company_id,omitempty" validate:"omitempty,gt=0" faker:"boundary_start=1,boundary_end=10000"`
	Views              int64     `json:"views,omitempty" validate:"omitempty,gte=0" faker:"boundary_start=0,boundary_end=10000"`
	ClosedTime         time.Time `json:"closed_time,omitempty" validate:"omitempty,gtfield=OriginalListedTime"`
	ExpiryTime         time.Time `json:"expiry_time,omitempty" validate:"omitempty,gtfield=OriginalListedTime"`
	OriginalListedTime time.Time `json:"original_listed_time,omitempty" validate:"omitempty"`
}

type ReqLoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ReqGetCompanyList struct {
	Page  int64  `json:"page" query:"page"`
	Limit int64  `json:"limit" query:"limit"`
	Name  string `json:"name" query:"name"`
}

type ReqGetJobList struct {
	Page      uint  `json:"page" query:"page"`
	Limit     uint  `json:"limit" query:"limit"`
	MinSalary int64 `json:"min_salary"`
	MaxSalary int64 `json:"max_salary"`
}

type ReqCreateMovie struct {
	Title              string    `json:"title,omitempty" validate:"required,max=100"`
	Genre 			   string    `json:"genre,omitempty" validate:"required,min=3"`
	Year			   int       `json:"year,omitempty" validate:"required,gte=1900,year_lt_now"` //
	Rating 			   float32   `json:"rating" validate:"required,gte=0,lte=5"`
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
	return year < currentYear
}


