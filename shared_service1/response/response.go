package response

import (
	"github.com/labstack/echo/v4"
)

type response struct {
	ctx echo.Context
}

func Response(ctx echo.Context) HttpSuccess {
	return &response{
		ctx: ctx,
	}
}
