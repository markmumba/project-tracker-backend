package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : update the update project function to use the id and the new data
// TODO : update the delete project function to use the id
// TODO : update all remaining functions to use the helper function

type ProjectController struct {
	projectService *services.ProjectService
}

func NewProjectController(projectService *services.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (pc *ProjectController) CreateProject(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var project models.Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	project.StudentID= userID // Set the user ID to the project
	if err := pc.projectService.CreateProject(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.ProjectToDTO(&project))
}

func (pc *ProjectController) GetProject(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	project, err := pc.projectService.GetProject(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.ProjectToDTO(project))
}

func (pc *ProjectController) GetAllProjectByLecturerId(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projects, err := pc.projectService.GetProjectsByLecturerId(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.ProjectToDTOs(projects))
}

func (pc *ProjectController) UpdateProject(c echo.Context) error {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := pc.projectService.UpdateProject(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.ProjectToDTO(&project))
}

func (pc *ProjectController) DeleteProject(c echo.Context) error {
	var projectParams models.Project
	if err := c.Bind(&projectParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := projectParams.ID
	if err := pc.projectService.DeleteProject(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Project deleted successfully")
}