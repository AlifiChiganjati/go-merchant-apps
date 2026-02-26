package jwttoken

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/pkg/response"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	issuer string
	key    []byte
}
type JwtClaim struct {
	DataClaims JwtClaims `json:"data"`
	jwt.RegisteredClaims
}

type JwtClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var (
	appName          = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
)

func NewJWTService(issuer string, key []byte) *JWTService {
	return &JWTService{
		issuer: issuer,
		key:    key,
	}
}

func getJWTKey() []byte {
	return []byte(os.Getenv("TOKEN_KEY"))
}

func (j *JWTService) GenerateToken(id, name, role string, expiredAt int64) (string, error) {
	if len(j.key) == 0 {
		return "", fmt.Errorf("jwt key empty")
	}

	claims := JwtClaim{
		DataClaims: JwtClaims{
			ID:   id,
			Name: name,
			Role: role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiredAt, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)
	return token.SignedString(j.key)
}

func (j *JWTService) JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			response.SendErrorResponse(c, http.StatusUnauthorized, "Authorization header required")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.SendErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &JwtClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return j.key, nil
		})

		if err != nil || !token.Valid {
			response.SendErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Role check
		if len(roles) > 0 {
			validRole := false
			for _, role := range roles {
				if role == claims.DataClaims.Role {
					validRole = true
					break
				}
			}
			if !validRole {
				response.SendErrorResponse(c, http.StatusForbidden, "You don't have permission")
				c.Abort()
				return
			}
		}

		c.Set("claims", claims)
		c.Next()
	}
}
