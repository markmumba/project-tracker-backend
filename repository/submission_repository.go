package repository

import (
	"errors"

	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

// repositories/submission_repository_impl.go

type SubmissionRepositoryImpl struct{}

func NewSubmissionRepository() SubmissionRepository {
	return &SubmissionRepositoryImpl{}
}

func (r *SubmissionRepositoryImpl) CreateSubmission(submission *models.Submission) error {
	result := database.DB.Create(submission)
	if result.Error != nil {
		return result.Error
	}
	err := database.DB.Preload("Project").First(submission, submission.ID).Error
	return err
}

func (r *SubmissionRepositoryImpl) GetSubmission(id uint) (*models.Submission, error) {
	var submission models.Submission
	result := database.DB.Preload("Project").Preload("Student").First(&submission, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &submission, nil
}

func (r *SubmissionRepositoryImpl) GetAllSubmissionByStudentId(studentId uint) ([]models.Submission, error) {
	var submissions []models.Submission
	result := database.DB.Where("student_id = ?", studentId).Find(&submissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return submissions, nil
}
func (r *SubmissionRepositoryImpl) GetSubmissionsByLecturer(lecturerID uint) ([]models.Submission, error) {
	var submissions []models.Submission

	result := database.DB.
		Joins("JOIN projects ON projects.id = submissions.project_id").
		Where("projects.lecturer_id = ?", lecturerID).
		Preload("Project").
		Preload("Student").
		Find(&submissions)

	if result.Error != nil {
		return nil, result.Error
	}
	return submissions, nil
}

func (r *SubmissionRepositoryImpl) UpdateSubmission(submission *models.Submission, id uint) error {
	// Find the submission by ID
	var existingSubmission models.Submission
	
	result := database.DB.First(&existingSubmission, id)
	if result.Error != nil {
		return result.Error
	}

	existingSubmission.Description = submission.Description
	existingSubmission.DocumentPath = submission.DocumentPath
	existingSubmission.SubmissionDate = submission.SubmissionDate
	existingSubmission.Reviewed = submission.Reviewed


	result = database.DB.Save(&existingSubmission)
	return result.Error
}

func (r *SubmissionRepositoryImpl) DeleteSubmission(id uint) error {
    result := database.DB.Delete(&models.Submission{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("submission not found")
    }
    return nil
}
func (r *SubmissionRepositoryImpl) GetAllSubmissions() ([]models.Submission, error) {
	var submissions []models.Submission
	result := database.DB.Preload("Project").Preload("Student").Find(&submissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return submissions, nil
}
