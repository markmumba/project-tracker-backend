package repository

import "github.com/markmumba/project-tracker/models"

type CommunicationRepository interface {
	SaveMessage(message *models.CommunicationHistory) error
	GetMessagesBetweenUsers(senderID, receiverID uint) ([]models.CommunicationHistory, error)
}
