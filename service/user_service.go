package service

import (
	"errors"
	"project-akhir-exam/entity"
	"project-akhir-exam/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Address  string `form:"address" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService interface {
	Register(input RegisterInput, fileLocation string) (entity.User, error)
	Login(input LoginInput) (entity.User, error)
	// UploadPhoto(ID int, fileLocation string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	GetUserByRoleID(RoleID int) (entity.User, error)
	IsEmailAvailable(email string) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) Register(input RegisterInput, fileLocation string) (entity.User, error) {
	user := entity.User{}
	user.Name = input.Name
	user.Email = input.Email

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.RoleID = 2
	user.Phone = input.Phone
	user.Address = input.Address
	user.Photo = fileLocation

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input LoginInput) (entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (s *userService) UploadPhoto(ID int, fileLocation string) (entity.User, error) {
// 	user, err := s.userRepository.FindByID(ID)
// 	if err != nil {
// 		return user, err
// 	}

// 	user.Photo = fileLocation

// 	updatedUser, err := s.userRepository.Update(user)
// 	if err != nil {
// 		return updatedUser, err
// 	}

// 	return updatedUser, nil
// }

func (s *userService) GetUserByID(ID int) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that id")
	}

	return user, nil
}

func (s *userService) GetUserByRoleID(RoleID int) (entity.User, error) {
	user, err := s.userRepository.FindByRoleID(RoleID)
	if err != nil {
		return user, err
	}

	if user.RoleID == 0 {
		return user, errors.New("no role found on that id")
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(email string) (bool, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}
