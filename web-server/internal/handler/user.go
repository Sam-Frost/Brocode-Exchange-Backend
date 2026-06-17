package handler

import (
	"net/http"

	"github.com/Sam-Frost/web-server/internal/dto"
	"github.com/Sam-Frost/web-server/internal/service"
	"github.com/Sam-Frost/web-server/internal/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	server  *util.Server
	service service.UserService
}

func NewUserHandler(server *util.Server, service service.UserService) UserHandler {
	return UserHandler{
		server:  server,
		service: service,
	}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {

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

func (u *UserHandler) LoginUser(c *gin.Context) {

	var requestBody dto.LoginUserRequest

	if err := c.ShouldBindJSON(requestBody); err != nil {
		if validationErrors, ok := util.CreateValidationErrorResponse(err); ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.APIResponse{
				Success: false,
				Error:   validationErrors,
			})
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err,
		})
	}

	loginResponse, err := u.service.LoginUser(c.Request.Context(), requestBody)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    loginResponse,
	})
}

func (u *UserHandler) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}
