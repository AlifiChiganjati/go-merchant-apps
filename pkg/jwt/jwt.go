package common

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtClaim struct {
	jwt.StandardClaims
	DataClaims JwtClaims `json:"data"`
}
type JwtClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var (
	appName          = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("TOKEN_KEY"))
)

func GenerateTokenJwt(id, name, role string, expiredAt int64) (string, error) {
	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: expiredAt, // expayet waktu login
		},
		DataClaims: JwtClaims{
			ID:   id,
			Name: name,
			Role: role,
		},
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)
	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println(authHeader)
		if !strings.Contains(authHeader, "Bearer") {
			response.SendErrorResponse(c, http.StatusForbidden, "Invalid Token")
			c.Abort()
			return
		}

		// jwtSignatureKey := []byte(os.Getenv("SIGNATURE_KEY"))
		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
		claims := &JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return jwtSignatureKey, nil
		})
		if err != nil {
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		if !token.Valid {
			response.SendErrorResponse(c, http.StatusUnauthorized, "Unaunthorized user")
			c.Abort()
			return
		}

		expiredAt := claims.ExpiresAt
		if time.Now().Unix() > expiredAt {
			response.SendErrorResponse(c, http.StatusUnauthorized, "Expired Token")
			c.Abort()
			return
		}

		// validation role

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.DataClaims.Role {
					validRole = true
					break
				}
			}
		}
		if !validRole {
			response.SendErrorResponse(c, http.StatusForbidden, "You dont have permission")
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
