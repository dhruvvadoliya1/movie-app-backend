package structs

import (
	"time"
)

// All response sturcts
// Response struct have Res prefix

type ResGetJobDetails struct {
	Status string `json:"status"`
	Data   ResJob `json:"data"`
}

type Meta struct {
	Count int64 `json:"count"`
}

type ResGetJobList struct {
	Status string `json:"status"`
	Data   struct {
		Jobs []ResJob `json:"jobs"`
		Meta Meta     `json:"meta"`
	}
}

type ResDeleteJob struct {
	Status string `json:"status"`
	JobId  int64  `json:"jobId"`
}

type ResJob struct {
	JobId              int64     `json:"id"`
	CompanyName        string    `json:"company_name"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	MaxSalary          int64     `json:"max_salary"`
	MinSalary          int64     `json:"min_salary"`
	PayPeriod          string    `json:"pay_period"`
	Location           string    `json:"location"`
	CompanyId          int64     `json:"company_id"`
	Views              int64     `json:"views"`
	ClosedTime         time.Time `json:"closed_time"`
	ExpiryTime         time.Time `json:"expiry_time"`
	OriginalListedTime time.Time `json:"original_listed_time"`
}

type ResGetCompanyDetails struct {
	Status string     `json:"status"`
	Data   ResCompany `json:"data"`
}

type ResGetCompanyList struct {
	Status string `json:"status"`
	Data   struct {
		Company []ResCompany `json:"company"`
		Meta    Meta         `json:"meta"`
	} `json:"data"`
}

type ResCompany struct {
	CompanyId   int64  `json:"id"`
	CompanyName string `json:"company_name"`
	Country     string `json:"country"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	Address     string `json:"address"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Industry    string `json:"industry"`
}

type ResMovie struct {
	Id     int64   `json:"id"`
	Title  string  `json:"title,omitempty"`
	Genre  string  `json:"genre,omitempty"`
	Year   int     `json:"year,omitempty"`
	Rating float32 `json:"rating,omitempty"`
}

type ResGetMovieList struct {
	Status string `json:"status"`
	Data   struct {
		Movies []ResMovie `json:"movies"`
		Meta    Meta         `json:"meta"`
	} `json:"data"`
}