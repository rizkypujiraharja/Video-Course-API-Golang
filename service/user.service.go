package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerRequest request.RegisterRequest) (*entity.User, error)
	UpdateUser(updateUserLoginRequest request.UpdateUserLoginRequest, userId string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByID(userID string) (*entity.User, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) UpdateUser(updateUserLoginRequest request.UpdateUserLoginRequest, userId string) (*entity.User, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&updateUserLoginRequest))

	if err != nil {
		return nil, err
	}

	intUserId, _ := strconv.Atoi(userId)

	user.ID = int64(intUserId)

	user, err = c.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (c *userService) CreateUser(registerRequest request.RegisterRequest) (*entity.User, error) {
	user, err := c.userRepo.FindByEmail(registerRequest.Email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}
	user.Role = "user"
	user, _ = c.userRepo.InsertUser(user)
	return &user, nil
}

func (c *userService) FindUserByEmail(email string) (*entity.User, error) {
	user, err := c.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *userService) FindUserByID(userID string) (*entity.User, error) {
	user, err := c.userRepo.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	usr := entity.User{}
	err = smapping.FillStruct(&usr, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
