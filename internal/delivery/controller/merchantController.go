package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	mc usecase.MerchantUsecase
	rg *gin.RouterGroup
}

func NewMerchantController(
	mc usecase.MerchantUsecase,
	rg *gin.RouterGroup,
) *MerchantController {
	return &MerchantController{
		mc: mc,
		rg: rg,
	}
}

func (con *MerchantController) Route() {
	con.rg.POST(
		"/merchants",
		jwttoken.JWTAuth("customer"),
		con.createMerchantHandler,
	)
}

func (con *MerchantController) createMerchantHandler(c *gin.Context) {
	var payload dto.MerchantRequestDto

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		response.SendErrorResponse(c, http.StatusInternalServerError, "claims not found")
		return
	}

	userID := claims.(*jwttoken.JwtClaim).DataClaims.ID

	merchant, err := con.mc.RegisterMerchant(userID, payload)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp := dto.MerchantResponseDto{
		MerchantName:    merchant.MerchantName,
		MerchantAddress: merchant.MerchantAddress,
	}

	response.SendSingleResponse(c, "success", resp)
}
