package repository

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

type ProjectRepositoryImpl struct{}

func NewProjectRepository() ProjectRepository {
	return &ProjectRepositoryImpl{}
}

func (r *ProjectRepositoryImpl) CreateProject(project *models.Project) error {
	result := database.DB.Create(project)
	return result.Error
}

func (r *ProjectRepositoryImpl) GetProject(id uint) (*models.Project, error) {
	var project models.Project
	result := database.DB.
		Preload("Lecturer.Role").
		Preload("Student.Role").
		Where("student_id = ?", id).
		First(&project)
	return &project, result.Error
}

func (r *ProjectRepositoryImpl) GetProjectsByLecturerId(lecturerId uint) ([]models.Project, error) {
	var projects []models.Project
	result := database.DB.Where("lecturer_id = ?", lecturerId).Find(&projects)
	return projects, result.Error
}

func (r *ProjectRepositoryImpl) UpdateProject(project *models.Project) error {
	result := database.DB.Save(project)
	return result.Error
}

func (r *ProjectRepositoryImpl) DeleteProject(id uint) error {
	var project models.Project
	result := database.DB.Delete(&project, id)
	return result.Error
}
