package auth_middleware

import (
	auth_dto "{{ .ModPath }}/src/module/auth_service/dto"

	"git.sriss.uz/shared/shared_service/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware = middleware.AuthEchoMiddleware[*auth_dto.AuthUser, echo.Context, *gorm.DB, struct{}]
