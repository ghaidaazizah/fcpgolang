package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var loginData model.UserLogin

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("email or password is empty"))
		return
	}

	user, err := u.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") {
			c.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid credentials"))
		} else {
			c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		}
		return
	}

	token, err := u.userService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("could not generate token"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponseWithData("login success", gin.H{
		"token": token,
	}))
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
		return
	}

	tasks, err := u.userService.GetUserTasksByCategory(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("could not retrieve tasks"))
		return
	}

	c.JSON(http.StatusOK, tasks)
}
