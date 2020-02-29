package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/x-color/vue-trello/model"
)

// UserUsecase is interface. It defines to control a User authentication.
type UserUsecase interface {
	SignUp(user model.User) (model.User, error)
	SignIn(user model.User) (model.User, error)
}

// UserInteractor includes repogitories and a logger.
type UserInteractor struct {
	userRepo UserRepository
	logger   Logger
}

// NewUserInteractor generates new interactor for a User.
func NewUserInteractor(
	userRepo UserRepository,
	logger Logger,
) (UserInteractor, error) {
	i := UserInteractor{
		userRepo: userRepo,
		logger:   logger,
	}
	return i, nil
}

// SignUp registers new User to repository.
func (i *UserInteractor) SignUp(user model.User) (model.User, error) {
	u, err := i.userRepo.FindByName(model.User{Name: user.Name})
	if err != nil && !errors.Is(err, model.NotFoundError{}) {
		logError(i.logger, err)
		return model.User{}, err
	}
	if user.Name == u.Name {
		i.logger.Info(formatLogMsg(user.ID, "New user name conflicts. '"+user.Name+"' already exists"))
		return model.User{}, model.ConflictError{
			UserID: user.ID,
			Err:    nil,
			ID:     user.Name,
			Act:    "validate name",
		}
	}

	user.ID = uuid.New().String()
	if err := i.userRepo.Create(user); err != nil {
		logError(i.logger, err)
		return model.User{}, err
	}
	i.logger.Info(formatLogMsg(user.ID, "Create user("+user.ID+")"))
	return user, nil
}

// SignIn returns User data if a user succeds at authentication.
func (i *UserInteractor) SignIn(user model.User) (model.User, error) {
	u, err := i.userRepo.FindByName(model.User{Name: user.Name})
	if err != nil {
		logError(i.logger, err)
		return model.User{}, err
	}
	if user.Password != u.Password {
		i.logger.Info(formatLogMsg(u.ID, "Invalid password. '"+u.ID+"' Fails to sign in"))
		return model.User{}, model.NotFoundError{
			UserID: u.ID,
			Err:    nil,
			ID:     u.ID,
			Act:    "validate password",
		}
	}

	i.logger.Info(formatLogMsg(u.ID, "Sign in user("+u.ID+")"))
	return u, nil
}
