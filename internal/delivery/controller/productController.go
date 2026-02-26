package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	pu         usecase.ProductUsecase
	rg         *gin.RouterGroup
	jwtService *jwttoken.JWTService
}

func NewProductController(
	pu usecase.ProductUsecase,
	rg *gin.RouterGroup,
	jwtService *jwttoken.JWTService,
) *ProductController {
	return &ProductController{pu: pu, rg: rg, jwtService: jwtService}
}

func (con *ProductController) Route() {
	con.rg.POST("/product", con.jwtService.JWTAuth("merchant"), con.addProductHandler)
	con.rg.POST("/products", con.productGetByNameHandler)
}

func (con *ProductController) addProductHandler(c *gin.Context) {
	var payload dto.ProductRequestDto

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

	product, err := con.pu.CreateProduct(userID, payload)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SendSingleResponse(c, "success", product)
}

func (con *ProductController) productGetByNameHandler(c *gin.Context) {
	var payload dto.ProductSearchRequestDto

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// default limit biar aman
	if payload.Limit <= 0 {
		payload.Limit = 10
	}

	products, total, err := con.pu.ProductGetByName(
		payload.Name,
		payload.Limit,
		payload.Offset,
	)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// convert []models.Product -> []any
	var result []any
	for _, p := range products {
		result = append(result, p)
	}

	response.SendPagedResponse(
		c,
		"success",
		result,
		map[string]any{
			"limit":  payload.Limit,
			"offset": payload.Offset,
			"total":  total,
		},
	)
}
