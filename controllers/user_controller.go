// CreateUser godoc
// @Summary Create a new user
// @Description This API creates a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Create User"
// @Success 201 {object} models.User
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /users [post]

package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// TODO : Get all the lecturers that is get all users where role id is 1
// @Summary Login User
// @Description Logs in a user and returns a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} echo.Map{"token":string}
// @Failure 400 {object} echo.Map{"error":string}
// @Failure 401 {object} echo.Map{"error":string}
// @Router /login [post]


func (uc *UserController) Login(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := uc.UserService.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
func (uc *UserController) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logout successful",
	})
}

func (uc *UserController) CreateUser(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    err := uc.UserService.CreateUser(&user)
    if err != nil {
        if strings.Contains(err.Error(), "already exists") {
            return c.JSON(http.StatusConflict, map[string]string{"message": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, models.UserToDTO(&user))
}

func (uc *UserController) GetUser(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user, err := uc.UserService.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.UserToDTO(user))
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(users))
}

func (uc *UserController) GetLecturers(c echo.Context) error {
	lecturers, err := uc.UserService.GetLecturers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(lecturers))
}

func (uc *UserController) GetStudentsByLecturerId(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	students, err := uc.UserService.GetStudentsByLecturer(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(students))
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var updateUser models.User
	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := uc.UserService.UpdateUser(userID, &updateUser); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.UserToDTO(&updateUser))
}

func (uc *UserController) UpdateUserProfileImage(c echo.Context) error {
	var image struct {
		ProfileImage string `json:"profile_image"`
	}
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Bind(&image); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uc.UserService.UpdateUserProfileImage(userID, image.ProfileImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Profile image updated successfully")
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = uc.UserService.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}
