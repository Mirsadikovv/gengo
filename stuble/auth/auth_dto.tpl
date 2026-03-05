package auth_dto

import (
	"github.com/Mirsadikovv/gengo/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthUser struct {
	UserId int64 `json:"userId"`
	RoleId int64 `json:"roleId"`
}

func (u *AuthUser) ID() int64 {
	return u.UserId
}

func (u *AuthUser) Pre(ctx echo.Context, db *gorm.DB, _ ...struct{}) (permission403 bool, _ error) {
	req := request.RequestWithData[AuthUser](ctx)
	req.SetUser(u)
	return false, nil
}
