package request

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	UnAuthorizedErrorMessage = "unauthorized"
	Authorization            = "Authorization"
	Bearer                   = "Bearer "
	AuthUser                 = "__@@AUTH_USER@@__"
	RequestData              = "__@@REQUEST_DATA@@__"
)

func Request(ctx echo.Context) *request[any] {
	return RequestWithData[any](ctx)
}

func RequestWithData[T any](ctx echo.Context) *request[T] {
	return &request[T]{
		ctx:    ctx,
		binder: echo.DefaultBinder{},
	}
}

type request[T any] struct {
	ctx    echo.Context
	binder echo.DefaultBinder
}

func (r *request[T]) SetUser(user *T) {
	r.ctx.Set(AuthUser, user)
}

func (r *request[T]) AuthUser() *T {
	u, _ := r.ctx.Get(AuthUser).(*T)

	return u
}

func (r *request[T]) SetData(data any) {
	r.ctx.Set(RequestData, data)
}

func (r *request[T]) Data() any {
	return r.ctx.Get(RequestData)
}

func (r *request[T]) EchoContext() echo.Context {
	return r.ctx
}

func (r *request[T]) Context() context.Context {
	return r.ctx.Request().Context()
}

func (r *request[T]) BindParam(in any) error {
	return r.binder.BindPathParams(r.ctx, in)
}

func (r *request[T]) BindQuery(in any) error {
	return r.binder.BindQueryParams(r.ctx, in)
}

func (r *request[T]) BindHeader(in any) error {
	return r.binder.BindHeaders(r.ctx, in)
}

func (r *request[T]) BindBody(in any) error {
	return r.binder.BindBody(r.ctx, in)
}

func (r *request[T]) Bind(in any) error {
	return r.binder.Bind(in, r.ctx)
}

func (r *request[T]) Param(name string) string {
	return r.ctx.Param(name)
}

func (r *request[T]) Query(name string) string {
	return r.ctx.QueryParam(name)
}

func (r *request[T]) Header(name string) string {
	return r.ctx.Request().Header.Get(name)
}

func (r *request[T]) ParamToInt(name string) (int64, error) {
	return strconv.ParseInt(r.Param(name), 10, 64)
}

func (r *request[T]) QueryToInt(name string) (int64, error) {
	return strconv.ParseInt(r.Query(name), 10, 64)
}

func (r *request[T]) HeaderToInt(name string) (int64, error) {
	return strconv.ParseInt(r.Header(name), 10, 64)
}

func (r *request[T]) ParamToFloat(name string) (float64, error) {
	return strconv.ParseFloat(r.Param(name), 64)
}

func (r *request[T]) QueryToFloat(name string) (float64, error) {
	return strconv.ParseFloat(r.Query(name), 64)
}

func (r *request[T]) HeaderToFloat(name string) (float64, error) {
	return strconv.ParseFloat(r.Header(name), 64)
}

func (r *request[T]) ParamToBool(name string) (bool, error) {
	return strconv.ParseBool(r.Param(name))
}

func (r *request[T]) QueryToBool(name string) (bool, error) {
	return strconv.ParseBool(r.Query(name))
}

func (r *request[T]) HeaderToBool(name string) (bool, error) {
	return strconv.ParseBool(r.Header(name))
}

func (r *request[T]) NewPaginate() *Paginate {
	return NewPaginateEchoWithContext(r.ctx)
}

func (r *request[T]) AuthorizationToken() string {
	return r.Header(Authorization)
}

func (r *request[T]) AuthorizationTokenWithBearer() (string, error) {

	token := r.AuthorizationToken()
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
