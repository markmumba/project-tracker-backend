package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

type CommunicationController struct {
	CommunicationService *services.CommunicationService
}

func NewCommunicationContoller(commService *services.CommunicationService) *CommunicationController {
	return &CommunicationController{
		CommunicationService: commService,
	}
}

func (comm *CommunicationController) SaveMessage(c echo.Context) error {
	message := new(models.CommunicationHistory)
	if err := c.Bind(message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err := comm.CommunicationService.SaveMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, message)
}

func (comm *CommunicationController) GetMessagesBetweenUsers(c echo.Context) error {
	senderID, _ := strconv.Atoi(c.QueryParam("sender_id"))
	receiverID, _ := strconv.Atoi(c.QueryParam("receiver_id"))

	messages, err := comm.CommunicationService.GetMessageBetweenUsers(uint(senderID), uint(receiverID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve messages"})
	}

	return c.JSON(http.StatusOK, messages)
}
