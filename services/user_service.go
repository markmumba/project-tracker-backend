package services

import (
	"errors"

	"github.com/markmumba/project-tracker/auth"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.UserRepository.CreateUser(user)
}

func (u *UserService) LoginUser(email, password string) (string, error) {
	var user models.User
	err := u.UserRepository.FindByEmail(email, &user)
	if err != nil {
		return "", err
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(&user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) GetUser(id uint) (*models.User, error) {
	return u.UserRepository.GetUser(id)
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.UserRepository.GetAllUsers()
}

func (u *UserService) GetStudentsByLecturer(lecturerID uint) ([]models.User, error) {
	return u.UserRepository.GetStudentsByLecturer(lecturerID)
}

func (u *UserService) GetLecturers() ([]models.User, error) {
	return u.UserRepository.GetLecturers()
}

func (u *UserService) UpdateUser(id uint, user *models.User) error {
	return u.UserRepository.UpdateUser(id, user)
}

func (u *UserService) UpdateUserProfileImage(id uint, profileImage string) error {
	return u.UserRepository.UpdateUserProfileImage(id, profileImage)
}

func (u *UserService) DeleteUser(id uint) error {
	return u.UserRepository.DeleteUser(id)
}
