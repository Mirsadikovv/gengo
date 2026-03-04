package {{ toSnake .ModuleName }}_cmd

import (
	auth_middleware "{{ .ModPath }}/src/module/auth_service/middleware"
	{{ toSnake .EntityName }}_handler "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}/handler"

	"git.sriss.uz/shared/shared_service/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title         {{ toCamel .ModuleName }} API
// @version       1.0
// @description   {{ toCamel .ModuleName }} API
// @BasePath      /api/v1
// @Schemes       http https
// @securityDefinitions.apikey ApiKeyAuth
// @in            header
// @name          Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	routerGroup := router.Group("/api/v1")
	{
		{{ toSnake .EntityName }}_handler.New{{ toCamel .EntityName }}Handler(routerGroup, db, log, authMiddleware)
	}
}
