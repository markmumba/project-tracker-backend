package repository

import "github.com/markmumba/project-tracker/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string, user *models.User) error
	GetUser(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetStudentsByLecturer(lecturerID uint) ([]models.User, error)
	GetLecturers() ([]models.User, error)
	UpdateUser(id uint, user *models.User) error
	UpdateUserProfileImage(id uint, profileImage string) error
	DeleteUser(id uint) error
}
