package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dhruvvadoliya1/movie-app-backend/cli/workers"
	"github.com/dhruvvadoliya1/movie-app-backend/constants"
	"github.com/dhruvvadoliya1/movie-app-backend/models"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/events"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/structs"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/watermill"
	"github.com/dhruvvadoliya1/movie-app-backend/services"
	"github.com/dhruvvadoliya1/movie-app-backend/utils"

	"github.com/doug-martin/goqu/v9"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gofiber/fiber/v2"
)

// UserController for user controllers
type UserController struct {
	userService *services.UserService
	logger      *zap.Logger
	event       *events.Events
	pub         *watermill.WatermillPublisher
}

// NewUserController returns a user
func NewUserController(goqu *goqu.Database, logger *zap.Logger, event *events.Events, pub *watermill.WatermillPublisher) (*UserController, error) {
	userModel, err := models.InitUserModel(goqu)
	if err != nil {
		return nil, err
	}

	userSvc := services.NewUserService(&userModel)

	return &UserController{
		userService: userSvc,
		logger:      logger,
		event:       event,
		pub:         pub,
	}, nil
}

// UserGet get the user by id
// swagger:route GET /users/{userId} Users RequestGetUser
//
// Get a user.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  200: ResponseGetUser
//	   400: GenericResFailNotFound
//		  500: GenericResError
func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	userID := c.Params(constants.ParamUid)
	user, err := ctrl.userService.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.JSONFail(c, http.StatusNotFound, constants.UserNotExist)
		}
		ctrl.logger.Error("error while get user by id", zap.Any("id", userID), zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetUser)
	}
	return utils.JSONSuccess(c, http.StatusOK, user)
}

// CreateUser registers a user
// swagger:route POST /users Users RequestCreateUser
//
// Register a user.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  201: ResponseCreateUser
//	   400: GenericResFailBadRequest
//		  500: GenericResError
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {

	var userReq structs.ReqRegisterUser

	err := json.Unmarshal(c.Body(), &userReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(userReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	user, err := ctrl.userService.RegisterUser(models.User{FirstName: userReq.FirstName, LastName: userReq.LastName, Email: userReq.Email, Password: userReq.Password, Roles: userReq.Roles}, ctrl.event)
	if err != nil {
		ctrl.logger.Error("error while insert user", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrInsertUser)
	}

	// publish message to queue
	welcomeMail := workers.WelcomeMail{FirstName: userReq.FirstName, LastName: userReq.LastName, Email: userReq.Email, Roles: userReq.Roles}
	err = ctrl.pub.Publish("user", welcomeMail)
	if err != nil {
		ctrl.logger.Error("error while publish message", zap.Error(err))
	}

	return utils.JSONSuccess(c, http.StatusCreated, user)
}
