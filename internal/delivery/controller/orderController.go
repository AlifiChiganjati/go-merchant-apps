package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	uc         usecase.OrderUsecase
	rg         *gin.RouterGroup
	jwtService *jwttoken.JWTService
}

func NewOrderController(
	uc usecase.OrderUsecase,
	rg *gin.RouterGroup,
	jwtService *jwttoken.JWTService,
) *OrderController {
	return &OrderController{uc: uc, rg: rg, jwtService: jwtService}
}

func (con *OrderController) Route() {
	con.rg.POST("/order", con.jwtService.JWTAuth("merchant", "customer"), con.createOrderHandler)
}

func (con *OrderController) createOrderHandler(c *gin.Context) {
	var payload dto.OrderRequestDto

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		response.SendErrorResponse(c, http.StatusUnauthorized, "unauthorized: claims not found")
		return
	}

	jwtClaim, ok := claims.(*jwttoken.JwtClaim)
	if !ok {
		response.SendErrorResponse(c, http.StatusInternalServerError, "invalid token claims")
		return
	}

	userID := jwtClaim.DataClaims.ID

	result, err := con.uc.CreateOrder(userID, payload)
	if err != nil {
		response.SendErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.SendSingleResponse(c, "order created successfully", result)
}
