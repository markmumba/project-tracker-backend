package models

type Feedback struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Comment      string     `gorm:"not null" json:"comment"`
	FeedbackDate string     `json:"feedback_date"`
	SubmissionID uint       `gorm:"not null" json:"submission_id"`
	LecturerID   uint       `gorm:"not null"`
	Submission   Submission `gorm:"foreignKey:SubmissionID"`
	Lecturer     User       `gorm:"foreignKey:LecturerID"`
}

type FeedbackDTO struct {
	ID           uint          `json:"id"`
	Comment      string        `json:"comment"`
	FeedbackDate string        `json:"feedback_date"`
	Submission   SubmissionDTO `json:"submission"`
	Lecturer     UserDTO       `json:"lecturer"`
}

func FeedbackToDTO(f *Feedback) FeedbackDTO {
	return FeedbackDTO{
		ID:           f.ID,
		Comment:      f.Comment,
		FeedbackDate: f.FeedbackDate,
		Submission:   SubmissionToDTO(&f.Submission),
		Lecturer:     UserToDTO(&f.Lecturer),
	}
}

func FeedbackToDTOs(feedbacks []Feedback) []FeedbackDTO {
	feedbackDTOs := make([]FeedbackDTO, len(feedbacks))
	for i, feedback := range feedbacks {
		feedbackDTOs[i] = FeedbackToDTO(&feedback)
	}
	return feedbackDTOs
}
