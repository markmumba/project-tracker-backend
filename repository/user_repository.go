package repository

import (
	"errors"
	"fmt"

	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}


func (r *UserRepositoryImpl) CreateUser(user *models.User) error {
    return database.DB.Transaction(func(tx *gorm.DB) error {
        var existingUser models.User
        result := tx.Where("email = ?", user.Email).First(&existingUser)

        // If the user already exists, return an error
        if result.Error == nil {
            return fmt.Errorf("user with email %s already exists", user.Email)
        }

        // If the error is not RecordNotFound, return the error
        if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return result.Error
        }

        // Proceed with user creation if no existing user is found
        result = tx.Create(user)
        return result.Error
    })
}

func (r *UserRepositoryImpl) FindByEmail(email string, user *models.User) error {
	return database.DB.Where("email = ?", email).First(user).Error
}

func (r *UserRepositoryImpl) GetUser(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.Preload("Role").First(&user, id)
	return &user, result.Error
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepositoryImpl) GetStudentsByLecturer(lecturerID uint) ([]models.User, error) {
	var projects []models.Project
	err := database.DB.Preload("Student").Where("lecturer_id = ?", lecturerID).Find(&projects).Error
	if err != nil {
		return nil, err
	}

	var students []models.User
	for _, project := range projects {
		students = append(students, project.Student)
	}

	return students, nil
}

func (r *UserRepositoryImpl) GetLecturers() ([]models.User, error) {
	var lecturers []models.User
	result := database.DB.Preload("Role").Where("role_id = 1").Find(&lecturers)
	return lecturers, result.Error
}

func (r *UserRepositoryImpl) UpdateUser(id uint, user *models.User) error {
	result := database.DB.Save(user).Where("id = ?", id)
	return result.Error
}

func (r *UserRepositoryImpl) UpdateUserProfileImage(id uint, profileImage string) error {
	err := database.DB.Model(&models.User{}).Where("id = ?", id).Update("profile_image", profileImage).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	var user models.User
	result := database.DB.Delete(&user, id)
	return result.Error
}
