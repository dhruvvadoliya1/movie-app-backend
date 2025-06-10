package utils

import (
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/structs"
)

// swagger:parameters RequestCreateUser
type RequestCreateUser struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqRegisterUser
	}
}

// swagger:parameters RequestCreateJob
type RequestCreateJob struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqCreateJob
	}
}

// swagger:parameters RequestGetJob
type RequestGetJob struct {
	// in:path
	JobId int64 `json:"jobId"`
}

// swagger:parameters RequestGetCompany
type RequestGetCompany struct {
	// in:path
	CompanyId int64 `json:"companyId"`
}

// swagger:parameters RequestGetCompanyList
// type RequestGetCompanyList struct {
// 	// Page number for pagination (optional)
// 	// in:query
// 	Page int64 `json:"page"`
// 	// Number of items to return (optional)
// 	// in:query
// 	Limit int64 `json:"limit"`
// 	// filter company by name (optional)
// 	// in:query
// 	Name string `json:"name"`
// }

// swagger:parameters RequestGetCompanyList
type RequestGetCompanyList struct {
	// in:query
	structs.ReqGetCompanyList
}

// swagger:parameters RequestGetJobList
// type RequestGetJobList struct {
// 	// Page number for pagination (optional)
// 	// in:query
// 	Page string `json:"page"`
// 	// Number of items to return (optional)
// 	// in:query
// 	Limit string `json:"limit"`
// 	// Minimum salary filter (optional)
// 	// in:query
// 	MinSalary string `json:"min_salary"`
// 	// Maximum salary filter (optional)
// 	// in:query
// 	MaxSalary string `json:"max_salary"`
// }

// swagger:parameters RequestGetJobList
type RequestGetJobList struct {
	// in:query
	structs.ReqGetJobList
}

// swagger:parameters RequestUpdateJob
type RequestUpdateJob struct {
	// in:path
	JobId int64 `json:"jobId"`
	// in:body
	// required: true
	Body struct {
		structs.ReqUpdateJob
	}
}

// swagger:parameters RequestDeleteJob
type RequestDeleteJob struct {
	// in:path
	JobId int64 `json:"jobId"`
}

// swagger:response ResponseCreateUser
type ResponseCreateUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			// models.User
		} `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestGetUser
type RequestGetUser struct {
	// in:path
	UserId string `json:"userId"`
}

// swagger:response ResponseGetUser
// type ResponseGetUser struct {
// 	// in:body
// 	Body struct {
// 		// enum: success
// 		Status string `json:"status"`
// 		Data   struct {
// 			models.User
// 		} `json:"data"`
// 	} `json:"body"`
// }

// swagger:parameters RequestAuthnUser
// type RequestAuthnUser struct {
// 	// in:body
// 	// required: true
// 	Body struct {
// 		structs.ReqLoginUser
// 	}
// }

// swagger:response ResponseAuthnUser
// type ResponseAuthnUser struct {
// 	// in:body
// 	Body struct {
// 		// enum: success
// 		Status string `json:"status"`
// 		Data   struct {
// 			models.User
// 		} `json:"data"`
// 	} `json:"body"`
// }

// swagger:response ResponseCreateJob
type ResponseCreateJob struct {
	// in:body
	Body struct {
		// enum: success
		Status string         `json:"status"`
		Data   structs.ResJob `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseGetJob
type ResponseGetJob struct {
	// in:body
	Body struct {
		// enum: success
		Status string         `json:"status"`
		Data   structs.ResJob `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseGetJobList
type ResponseGetJobList struct {
	// in:body
	Body structs.ResGetJobList `json:"body"`
}

// swagger:response ResponseUpdateJob
type ResponseUpdateJob struct {
	// in:body
	Body struct {
		// enum: success
		Status string         `json:"status"`
		Data   structs.ResJob `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseDeleteJob
type ResponseDeleteJob struct {
	// in:body
	Body struct {
		// enum: success
		Status string               `json:"status"`
		Data   structs.ResDeleteJob `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseGetCompany
type ResponseGetCompany struct {
	// in:body
	Body struct {
		Status string             `json:"status"`
		Data   structs.ResCompany `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseGetCompanyList
type ResponseGetCompanyList struct {
	// in:body
	Body struct {
		Status string                    `json:"status"`
		Data   structs.ResGetCompanyList `json:"data"`
	} `json:"body"`
}

////////////////////
// --- GENERIC ---//
////////////////////

// Response is okay
// swagger:response GenericResOk
type ResOK struct {
	// in:body
	Body struct {
		// enum:success
		Status string `json:"status"`
	}
}

// Fail due to user invalid input
// swagger:response GenericResFailBadRequest
type ResFailBadRequest struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to user invalid input
// swagger:response ResForbiddenRequest
type ResForbiddenRequest struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Server understand request but refuse to authorize it
// swagger:response GenericResFailConflict
type ResFailConflict struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to server understand request but unable to process
// swagger:response GenericResFailUnprocessableEntity
type ResFailUnprocessableEntity struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to resource not exists
// swagger:response GenericResFailNotFound
type ResFailNotFound struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Unexpected error occurred
// swagger:response GenericResError
type ResError struct {
	// in: body
	Body struct {
		// enum: error
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	} `json:"body"`
}
