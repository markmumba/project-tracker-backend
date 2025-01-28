package services

import (
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
)

type CommunicationService struct {
	CommunicationRepository repository.CommunicationRepository
}

func NewCommunicationService(commRepo repository.CommunicationRepository) *CommunicationService {

	return &CommunicationService{
		CommunicationRepository: commRepo,
	}
}

func (comm *CommunicationService) SaveMessage(message *models.CommunicationHistory) error {
	err := comm.CommunicationRepository.SaveMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (comm *CommunicationService) GetMessageBetweenUsers(senderID, receiverID uint) ([]models.CommunicationHistory, error) {
	messages, err := comm.CommunicationRepository.GetMessagesBetweenUsers(senderID, receiverID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
