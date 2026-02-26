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

func (con *AuthController) Route() {
	con.rg.POST("/register", con.registerHandler)
	con.rg.POST("/login", con.loginHandler)
}

func (con *AuthController) registerHandler(c *gin.Context) {
	var payload dto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := con.ac.CreateUser(payload)
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

func (con *AuthController) loginHandler(c *gin.Context) {
	var payload dto.LoginRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := con.ac.LoginUser(payload)
	if err != nil {
		if err.Error() == "1" {
			response.SendErrorResponse(c, http.StatusForbidden, "Password salah")
			return
		}
		response.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	response.SendSingleResponse(c, "success", loginData)
}
