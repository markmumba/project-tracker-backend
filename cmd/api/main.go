package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
	"github.com/markmumba/project-tracker/routes"
	"github.com/markmumba/project-tracker/services"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Project{},
		&models.Submission{},
		&models.Feedback{},
		&models.CommunicationHistory{},
	)

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)

	projectRepository := repository.NewProjectRepository()
	projectService := services.NewProjectService(projectRepository, userRepository)

	submissionRepository := repository.NewSubmissionRepository()
	submissionService := services.NewSubmissionService(submissionRepository, userRepository)

	feedbackRepository := repository.NewFeedbackRepository()
	feedbackService := services.NewFeedbackService(feedbackRepository)

	communicationRepository := repository.NewCommunicationRepository()
	communicationService := services.NewCommunicationService(communicationRepository)

	handler := routes.SetupRouter(userService, projectService, submissionService, feedbackService, communicationService)
	
	// Get port from environment variable with fallback to 8080
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
		log.Println("No BACKEND_PORT specified, defaulting to 8080")
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", portNum),
		Handler:     handler,
		ReadTimeout: time.Second * 10,
	}
	
	log.Printf("Server starting on port: %v\n", portNum)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}