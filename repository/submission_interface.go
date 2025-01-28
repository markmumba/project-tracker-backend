package repository

import "github.com/markmumba/project-tracker/models"

type SubmissionRepository interface {
	CreateSubmission(submission *models.Submission) error
	GetSubmission(id uint) (*models.Submission, error)
	GetAllSubmissionByStudentId(studentId uint) ([]models.Submission, error)
	GetSubmissionsByLecturer(lecturerID uint) ([]models.Submission, error)
	UpdateSubmission(submission *models.Submission, id uint) error
	DeleteSubmission(id uint) error
	GetAllSubmissions() ([]models.Submission, error)
}
