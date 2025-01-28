package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/auth"
	"github.com/markmumba/project-tracker/custommiddleware"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}
)

type WebsocketController struct {
	CommunicationService services.CommunicationService
	ProjectService       services.ProjectService
	clients              map[uint]*websocket.Conn
}

func NewWebsocketController(commService services.CommunicationService, projService services.ProjectService) *WebsocketController {
	return &WebsocketController{
		CommunicationService: commService,
		ProjectService:       projService,
		clients:              make(map[uint]*websocket.Conn),
	}
}
func (wsc *WebsocketController) HandleWebSocket(c echo.Context) error {
	// Extract token from cookie
	cookie, err := c.Cookie("token")
	if err != nil {
		log.Println("could not get cookie man ")
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	// Parse and validate JWT token
	token, err := jwt.ParseWithClaims(cookie.Value, &auth.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return custommiddleware.JwtSecret, nil
	})
	if err != nil {
		log.Println("Parsing the cookie was crazy man ")
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	// Extract user ID from token claims
	claims, ok := token.Claims.(*auth.JwtCustomClaims)
	if !ok || !token.Valid {
		log.Println("here we are trying to get the user id ")
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	// Upgrade the HTTP request to a WebSocket connection
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("At this point we are trying to upgrade to the websocket")
		return err
	}
	defer ws.Close()

	// Register client
	userID := claims.UserId
	wsc.clients[userID] = ws

	// Handle messages
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(wsc.clients, userID)
			break
		}

		var chatMessage struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(msg, &chatMessage); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		// Get the project associated with the student
		project, err := wsc.ProjectService.GetProject(userID)
		if err != nil {
			log.Printf("error getting project: %v", err)
			continue
		}

		// Create CommunicationHistory entry
		communication := models.CommunicationHistory{
			SenderID:   userID,
			ReceiverID: project.LecturerID,
			Message:    chatMessage.Message,
			Timestamp:  time.Now(),
		}

		// Save message to database
		err = wsc.CommunicationService.SaveMessage(&communication)
		if err != nil {
			log.Printf("error saving message: %v", err)
		}

		// Send message to recipient (lecturer)
		if client, ok := wsc.clients[project.LecturerID]; ok {
			err = client.WriteJSON(communication)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(wsc.clients, project.LecturerID)
			}
		}
	}

	return nil
}
