package repository

import "github.com/markmumba/project-tracker/models"

type FeedbackRepository interface {
	CreateFeedback(feedback *models.Feedback) error
	GetFeedback(id uint) (*models.Feedback, error)
	GetFeedbackByStudent(studentID uint) (*[]models.Feedback, error)
	GetFeedbackByLecturer(lecturerID uint) (*[]models.Feedback, error)
	GetAllFeedback() ([]models.Feedback, error)
	GetFeedbackBySubmissionId(submissionId uint) ([]models.Feedback, error)
	UpdateFeedback(feedback *models.Feedback) error
	GetFeedbackForSubmission(id uint)(*models.Feedback ,error)
	DeleteFeedback(id uint) error
}
