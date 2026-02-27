package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	cu         usecase.CartUsecase
	rg         *gin.RouterGroup
	jwtService *jwttoken.JWTService
}

func NewCartController(cu usecase.CartUsecase, rg *gin.RouterGroup, jwtService *jwttoken.JWTService) *CartController {
	return &CartController{
		cu:         cu,
		rg:         rg,
		jwtService: jwtService,
	}
}

func (con *CartController) Route() {
	con.rg.POST("/cart", con.jwtService.JWTAuth("merchant", "customer"), con.createCartHandler)
}

func (con *CartController) createCartHandler(c *gin.Context) {
	var payload dto.CartRequestDto

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		response.SendErrorResponse(c, http.StatusInternalServerError, "claims not found")
		return
	}

	payload.UserID = claims.(*jwttoken.JwtClaim).DataClaims.ID
	cart, err := con.cu.CreateCart(payload)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendSingleResponse(c, "success", cart)
}
