package services

import (
	"github.com/dhruvvadoliya1/movie-app-backend/constants"
	"github.com/dhruvvadoliya1/movie-app-backend/models"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/events"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/structs"
)

type UserService struct {
	userModel *models.UserModel
}

func NewUserService(userModel *models.UserModel) *UserService {
	return &UserService{
		userModel: userModel,
	}
}

func (userSvc *UserService) RegisterUser(user models.User, event events.IEvents) (models.User, error) {
	user, err := userSvc.userModel.InsertUser(user)
	if err != nil {
		return user, err
	}

	event.Publish(constants.EventUserRegistered, structs.EventUserRegistered{Email: user.Email})
	return user, err
}

func (userSvc *UserService) GetUser(userId string) (models.User, error) {
	user, err := userSvc.userModel.GetById(userId)
	return user, err
}

// Authenticate verify identity using email, and password.
// On successful validtion it'll return the user
func (userSvc *UserService) Authenticate(email, password string) (models.User, error) {
	return userSvc.userModel.GetUserByEmailAndPassword(email, password)
}
