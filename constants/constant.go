package constants

// variables
const (
	CookieUser   = "user"
	KratosCookie = "ory_kratos_session"
)

// fiber contexts
const (
	ContextUid = "userId"
	ParamCid   = "companyId"
	ParamJid   = "jobId"
	ParamMid   = "movieId"
)

// kratos
const (
	KratosID          = "kratosId"
	KratosUserDetails = "kratosUserDetails"
)

// params
const (
	ParamUid = "userId"
)

// Success messages
// ...

// Fail messages
// ...
const (
	Unauthenticated    = "unauthenticated to access resource"
	Unauthorized       = "unauthorized to access resource"
	InvalidCredentials = "invalid credenticals"
	UserNotExist       = "user does not exists"
	CompanyNotExist    = "company does not exists"
	InvalidCompanyId   = "invalid companyId"
	InvalidMovieId     = "invalid movieId"
	InvalidJobId       = "invalid jobId"
	JobNotExist        = "job does not exists"
	MovieNotExist      = "movie does not exists"
)

// Error messages
const (
	ErrGetUser             = "error while get user"
	ErrLoginUser           = "error while login user"
	ErrInsertUser          = "error while creating user, please try after sometime"
	ErrHealthCheckDb       = "error while checking health of database"
	ErrUnauthenticated     = "error verifing user identity"
	ErrKratosAuth          = "error while fetching user from kratos"
	ErrKratosDataInsertion = "error while inserting user data came from kratos"
	ErrKratosIDEmpty       = "error no session_id found in kratos cookie"
	ErrKratosCookieTime    = "error while parsing the expiration time of the cookie"
	ErrGetCompany          = "error while get company"
	ErrGetJob              = "error while get job"
	ErrDeleteJob           = "error while delete job"
	ErrGetMovie			   = "error while get movie"
	ErrDeleteMovie		   = "error while delete movie"
)

// Events
const (
	EventUserRegistered = "event:userRegistered"
)
