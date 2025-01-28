package repository

import "github.com/markmumba/project-tracker/models"

type ProjectRepository interface {
	CreateProject(project *models.Project) error
	GetProject(id uint) (*models.Project, error)
	GetProjectsByLecturerId(lecturerId uint) ([]models.Project, error)
	UpdateProject(project *models.Project) error
	DeleteProject(id uint) error
}
