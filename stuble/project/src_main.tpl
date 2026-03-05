package src

import (
	"fmt"
	"log"

	{{ toSnake .ModuleName }}_cmd "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}"
	{{ toSnake .EntityName }}_model "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}/model"
	auth_dto "{{ .ModPath }}/src/module/auth_service/dto"

	"github.com/Mirsadikovv/gengo/shared_service/jwt"
	"github.com/Mirsadikovv/gengo/shared_service/logger"
	"github.com/Mirsadikovv/gengo/shared_service/middleware"
	"github.com/Mirsadikovv/gengo/shared_service/pg"
	"github.com/Mirsadikovv/gengo/shared_service/redis"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Env struct {
	HTTP_Host          string `env:"HTTP_HOST"           default:"localhost"`
	HTTP_Port          int    `env:"HTTP_PORT"           default:"8080"`
	DB_Host            string `env:"DB_HOST"`
	DB_User            string `env:"DB_USER"`
	DB_Password        string `env:"DB_PASSWORD"`
	DB_DBName          string `env:"DB_NAME"`
	DB_Port            int    `env:"DB_PORT"`
	DB_SSLMode         string `env:"DB_SSL_MODE"         default:"disable"`
	DB_TimeZone        string `env:"DB_TIME_ZONE"        default:"UTC"`
	JWT_Secret         string `env:"JWT_SECRET"`
	JWT_Expired        int64  `env:"JWT_EXPIRED"`
	JWT_RefreshExpired int64  `env:"JWT_REFRESH_EXPIRED"`
	REDIS_Addr         string `env:"REDIS_ADDR"`
	SwaggerUser        string `env:"SWAGGER_USER"        default:""`
	SwaggerPassword    string `env:"SWAGGER_PASSWORD"    default:""`
}

func Exec(env *Env) {

	redisConfig := redis.Config{
		Addr: env.REDIS_Addr,
	}

	jwtConfig := &jwt.JwtConfig{
		Secret:         env.JWT_Secret,
		Expired:        env.JWT_Expired,
		RefreshExpired: env.JWT_RefreshExpired,
	}

	gormConfig := &pg.GormConfig{
		SkipDefaultTransaction: true,
	}

	pgConfig := pg.ConnectionConfig{
		Host:     env.DB_Host,
		User:     env.DB_User,
		Password: env.DB_Password,
		DBName:   env.DB_DBName,
		Port:     env.DB_Port,
		SSLMode:  env.DB_SSLMode,
		TimeZone: env.DB_TimeZone,
	}

	db := pg.Primary(gormConfig, pgConfig)

	if err := migration(db); err != nil {
		log.Println(err)
	}

	router := echo.New()

	router.Use(echo_middleware.CORS())

	log := logger.New()

	memoryCache := redis.Instance(redisConfig)

	authMiddleware := middleware.NewAuthEchoMiddleware[*auth_dto.AuthUser](jwtConfig, memoryCache, db)
	{
		// @gengo:modules
		{{ toSnake .ModuleName }}_cmd.Cmd(router, db, log, authMiddleware)
	}

	router.Start(fmt.Sprintf(":%d", env.HTTP_Port))
}

func migration(db *gorm.DB) error {
	return db.AutoMigrate(
		// @gengo:migration
		&{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}{},
	)
}
