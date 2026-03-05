package middleware

import (
	"errors"
	"strings"

	"github.com/Mirsadikovv/gengo/shared_service/jwt"
	"github.com/Mirsadikovv/gengo/shared_service/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware[U jwt.IUser[T, E, K], T, E, K any] struct {
	jwt          jwt.JwtService[U, T, E, K]
	redisService redis.Client
	db           *gorm.DB
}

func NewAuthMiddleware[U jwt.IUser[T, E, K], T, E, K any](
	jwt jwt.JwtService[U, T, E, K],
	redisService redis.Client,
	db *gorm.DB,
) *AuthMiddleware[U, T, E, K] {
	return &AuthMiddleware[U, T, E, K]{
		jwt:          jwt,
		redisService: redisService,
		db:           db,
	}
}

func AuthorizationToken(c echo.Context) (string, error) {
	const (
		UnAuthorizedErrorMessage = "unauthorized"
		Authorization            = "Authorization"
		Bearer                   = "Bearer "
	)

	token := c.Request().Header.Get(Authorization)
	{
		if token == "" || !strings.HasPrefix(token, Bearer) {
			return "", errors.New(UnAuthorizedErrorMessage)
		}

		if strings.Count(token, ".") != 2 {
			return "", errors.New(UnAuthorizedErrorMessage)
		}

		token = strings.TrimPrefix(token, Bearer)
	}

	return token, nil
}
