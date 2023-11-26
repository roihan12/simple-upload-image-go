package services

import (
	"errors"


	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/app/request"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/helpers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/models"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/repository"
	"golang.org/x/crypto/bcrypt"
)


type UserService interface {
	RegisterUser(input request.RegisterInput) (models.User, error)
	LoginUser(input request.LoginInput) (models.User, error)
	IsEmailAvailable(input string) (bool, error)
	IsUsernameAvailable(input string) (bool, error)
	GetUserByID(ID int) (models.User, error)
	UpdateUser(userID int, inputData request.UpdateInput) (models.User, error)
	DeleteUser(ID int) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (s *UserServiceImpl) RegisterUser(input request.RegisterInput) (models.User, error) {

	user := models.User{}
	user.Username = input.Username
	user.Email = input.Email

	passwordHash := helpers.HashPassword(input.Password)

	user.Password = string(passwordHash)

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *UserServiceImpl) LoginUser(input request.LoginInput) (models.User, error) {
	
	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	ok := helpers.ComparePassword(user.Password, input.Password)
	if !ok {
		return models.User{}, errors.New("password not match")
	}
	return user, nil
}

func (s *UserServiceImpl) IsEmailAvailable(input string) (bool, error) {
	user, err := s.userRepository.FindByEmail(input)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *UserServiceImpl) IsUsernameAvailable(input string) (bool, error) {
	user, err := s.userRepository.FindByUsername(input)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *UserServiceImpl) GetUserByID(ID int) (models.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found on that ID")
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUser(userID int, inputData request.UpdateInput) (models.User, error) {


	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return user, err
	}

	if user.ID != inputData.User.ID {
		return user, err
	}

	if user.Username == inputData.Username {
		user.Username = inputData.Username
	}

	var isUsernameAvailable string
	isUsernameAvailable = inputData.Username

	checkUsername, _ := s.IsUsernameAvailable(isUsernameAvailable)
	if checkUsername {
		user.Username = inputData.Username
	}

	if user.Email == inputData.Email {
		user.Email = inputData.Email
	}

	var isEmailAvailable string
	isEmailAvailable = inputData.Email

	checkEmail, _ := s.IsEmailAvailable(isEmailAvailable)
	if checkEmail {
		user.Email = inputData.Email
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	updatedUser, err := s.userRepository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *UserServiceImpl) DeleteUser(ID int) error {
	err := s.userRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepository repository.UserRepository,) UserService {
	return &UserServiceImpl{userRepository: userRepository,}
}
