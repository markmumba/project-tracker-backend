package repository

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

type CommunicationRepositoryImpl struct{}

func NewCommunicationRepository() CommunicationRepository {
	return &CommunicationRepositoryImpl{}
}

func (comm *CommunicationRepositoryImpl) SaveMessage(message *models.CommunicationHistory) error {

	result := database.DB.Create(message)
	return result.Error
}

func (comm *CommunicationRepositoryImpl) GetMessagesBetweenUsers(senderID, receiverID uint) ([]models.CommunicationHistory, error) {

	var messages []models.CommunicationHistory
	err := database.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		senderID, receiverID, receiverID, senderID).
		Order("timestamp asc").
		Find(&messages).Error
	return messages, err
}
