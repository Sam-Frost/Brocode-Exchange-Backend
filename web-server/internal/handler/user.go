package handler

import (
	"net/http"

	"github.com/Sam-Frost/web-server/internal/dto"
	"github.com/Sam-Frost/web-server/internal/service"
	"github.com/Sam-Frost/web-server/internal/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	server  *util.Server
	service service.UserService
}

func NewUserController(server *util.Server, service service.UserService) UserController {
	return UserController{
		server:  server,
		service: service,
	}
}

func (u *UserController) RegisterUser(c *gin.Context) {

	var requestBody dto.CreateUserRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		if validatonErrors, ok := util.CreateValidationErrorResponse(err); ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.APIResponse{
				Success: false,
				Error:   validatonErrors,
			})
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	responseDto, err := u.service.CreateUser(requestBody)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    responseDto,
	})
}

func (u *UserController) LoginUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}
