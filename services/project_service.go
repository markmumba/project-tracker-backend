package services

import (
	"errors"

	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
)

type ProjectService struct {
	ProjectRepository repository.ProjectRepository
	UserRepository    repository.UserRepository
}

func NewProjectService(projectRepo repository.ProjectRepository, userRepo repository.UserRepository) *ProjectService {
	return &ProjectService{
		ProjectRepository: projectRepo,
	}
}

func (p *ProjectService) CreateProject(project *models.Project) error {
	return p.ProjectRepository.CreateProject(project)
}

func (p *ProjectService) GetProject(id uint) (*models.Project, error) {
	return p.ProjectRepository.GetProject(id)
}

func (p *ProjectService) GetProjectsByLecturerId(lecturerId uint) ([]models.Project, error) {
	// Ensure the lecturer exists and is actually a lecturer
	user, err := p.UserRepository.GetUser(lecturerId)
	if err != nil {
		return nil, err
	}
	if user.Role.ID != 1 {
		return nil, errors.New("user is not a lecturer")
	}
	return p.ProjectRepository.GetProjectsByLecturerId(lecturerId)
}

func (p *ProjectService) UpdateProject(project *models.Project) error {
	return p.ProjectRepository.UpdateProject(project)
}

func (p *ProjectService) DeleteProject(id uint) error {
	return p.ProjectRepository.DeleteProject(id)
}
