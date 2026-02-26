package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	AuthController struct {
		ac usecase.AuthUsecase
		rg *gin.RouterGroup
	}
)

func NewAuthController(ac usecase.AuthUsecase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		ac: ac,
		rg: rg,
	}
}

func (acon *AuthController) registerHandler(c *gin.Context) {
	var payload dto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := acon.ac.CreateUser(payload)
	resp := dto.UserResponseDto{
		Fullname:    user.Fullname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SendCreateResponse(c, "User Berhasil dibuat!", resp)
}

func (acon *AuthController) Route() {
	acon.rg.POST("/register", acon.registerHandler)
}
