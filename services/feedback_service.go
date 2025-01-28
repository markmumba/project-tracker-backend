package services

import (

	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
)

type FeedbackService struct {
	FeedbackRepository repository.FeedbackRepository
}

func NewFeedbackService(feedbackRepo repository.FeedbackRepository) *FeedbackService {
	return &FeedbackService{
		FeedbackRepository: feedbackRepo,
	}
}

func (s *FeedbackService) CreateFeedback(feedback *models.Feedback) error {
	return s.FeedbackRepository.CreateFeedback(feedback)
}

func (s *FeedbackService) GetFeedback(id uint) (*models.Feedback, error) {
	return s.FeedbackRepository.GetFeedback(id)
}

func (s *FeedbackService) GetFeedbackByStudent(studentID uint) (*[]models.Feedback, error) {
	return s.FeedbackRepository.GetFeedbackByStudent(studentID)
}

func (s *FeedbackService) GetFeedbackByLecturer (LecturerID uint ) (*[]models.Feedback, error) {
    return s.FeedbackRepository.GetFeedbackByLecturer(LecturerID)
}


func (s *FeedbackService) GetAllFeedback() ([]models.Feedback, error) {
	return s.FeedbackRepository.GetAllFeedback()
}

func (s *FeedbackService) GetFeedbackBySubmissionId(submissionId uint) ([]models.Feedback, error) {
	return s.FeedbackRepository.GetFeedbackBySubmissionId(submissionId)
}

func (s *FeedbackService) UpdateFeedback(id uint, feedbackData *models.Feedback) (*models.Feedback, error) {
    existingFeedback, err := s.FeedbackRepository.GetFeedback(id)
    if err != nil {
        return nil, err
    }

    // Update only the fields that are allowed to be updated
    existingFeedback.Comment = feedbackData.Comment
    existingFeedback.FeedbackDate = feedbackData.FeedbackDate

    err = s.FeedbackRepository.UpdateFeedback(existingFeedback)
    if err != nil {
        return nil, err
    }

    return existingFeedback, nil
}

func (s *FeedbackService) GetFeedbackForSubmission(submissionID uint) (*models.Feedback, error) {
    return s.FeedbackRepository.GetFeedbackForSubmission(submissionID)
}

func (s *FeedbackService) DeleteFeedback(id uint) error {
	return s.FeedbackRepository.DeleteFeedback(id)
}
