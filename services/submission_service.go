package services

import (
	"errors"

	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
)

type SubmissionService struct {
	SubmissionRepository repository.SubmissionRepository
	UserRepository       repository.UserRepository
}

func NewSubmissionService(submissionRepo repository.SubmissionRepository, userRepo repository.UserRepository) *SubmissionService {
	return &SubmissionService{
		SubmissionRepository: submissionRepo,
		UserRepository:       userRepo,
	}
}

func (s *SubmissionService) CreateSubmission(submission *models.Submission) error {
	err := s.SubmissionRepository.CreateSubmission(submission)
	if err != nil {
		return err
	}
	return nil
}

func (s *SubmissionService) GetSubmission(id uint) (*models.Submission, error) {
	return s.SubmissionRepository.GetSubmission(id)
}

func (s *SubmissionService) GetAllSubmissionByStudentId(studentId uint) ([]models.Submission, error) {
	user, err := s.UserRepository.GetUser(studentId)
	if err != nil {
		return nil, err
	}
	if user.RoleID != 2 {
		return nil, errors.New("user is not a student")
	}
	return s.SubmissionRepository.GetAllSubmissionByStudentId(studentId)
}

func (s *SubmissionService) GetSubmissionsByLecturer(lecturerID uint) ([]models.Submission, error) {
	user, err := s.UserRepository.GetUser(lecturerID)
	if err != nil {
		return nil, err
	}
	if user.RoleID != 1 {
		return nil, errors.New("user is not a lecturer")
	}
	return s.SubmissionRepository.GetSubmissionsByLecturer(lecturerID)
}

func (s *SubmissionService) UpdateSubmission(submission *models.Submission, id uint) error {
	return s.SubmissionRepository.UpdateSubmission(submission, id)
}

func (s *SubmissionService) DeleteSubmission(id uint) error {
	return s.SubmissionRepository.DeleteSubmission(id)
}

func (s *SubmissionService) GetAllSubmissions() ([]models.Submission, error) {
	return s.SubmissionRepository.GetAllSubmissions()
}
