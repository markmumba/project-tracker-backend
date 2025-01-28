package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : get the latest feedback for the student

type FeedbackController struct {
	feedbackService *services.FeedbackService
}

func NewFeedbackController(feedbackService *services.FeedbackService) *FeedbackController {
	return &FeedbackController{
		feedbackService: feedbackService,
	}
}

func (fc *FeedbackController) CreateFeedback(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	feedback.LecturerID = userID
	if err := fc.feedbackService.CreateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.FeedbackToDTO(&feedback))
}

func (fc *FeedbackController) GetFeedbackByStudent(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	feedback, err := fc.feedbackService.GetFeedbackByStudent(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.FeedbackToDTOs(*feedback))
}
func (fc *FeedbackController) GetFeedbackByLecturer (c echo.Context)error {

	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	feedback, err := fc.feedbackService.GetFeedbackByLecturer(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.FeedbackToDTOs(*feedback))

}

func (fc *FeedbackController) GetAllFeedback(c echo.Context) error {
	feedbacks, err := fc.feedbackService.GetAllFeedback()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.FeedbackToDTOs(feedbacks))
}

func (fc *FeedbackController) UpdateFeedback(c echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid feedback ID")
    }

    var feedbackData models.Feedback
    if err := c.Bind(&feedbackData); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request body")
    }

    updatedFeedback, err := fc.feedbackService.UpdateFeedback(uint(id), &feedbackData)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update feedback")
    }

    return c.JSON(http.StatusOK, updatedFeedback)
}



func (fc *FeedbackController) GetFeedbackForSubmission(c echo.Context) error {
	submissionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid submission ID")
	}

	feedback, err := fc.feedbackService.GetFeedbackForSubmission(uint(submissionID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if feedback == nil {
		return c.JSON(http.StatusNotFound, "No feedback found for this submission")
	}

	return c.JSON(http.StatusOK, feedback)
}

func (fc *FeedbackController) DeleteFeedback(c echo.Context) error {
	var feedbackParams models.Feedback
	err := c.Bind(&feedbackParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := feedbackParams.ID
	err = fc.feedbackService.DeleteFeedback(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, "Feedback deleted successfully")
}
