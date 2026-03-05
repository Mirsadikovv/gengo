package request

import "github.com/Mirsadikovv/gengo/shared_service/response"

func (r *request[T]) OK(data ...any) error {
	return response.Response(r.ctx).OK(data...)
}

func (r *request[T]) Created(data ...any) error {
	return response.Response(r.ctx).Created(data...)
}

func (r *request[T]) Accepted(data ...any) error {
	return response.Response(r.ctx).Accepted(data...)
}

func (r *request[T]) NonAuthoritativeInfo(data ...any) error {
	return response.Response(r.ctx).NonAuthoritativeInfo(data...)
}

func (r *request[T]) NoContent() error {
	return response.Response(r.ctx).NoContent()
}

func (r *request[T]) ResetContent(data ...any) error {
	return response.Response(r.ctx).ResetContent(data...)
}

func (r *request[T]) PartialContent(data ...any) error {
	return response.Response(r.ctx).PartialContent(data...)
}

func (r *request[T]) MultiStatus(data ...any) error {
	return response.Response(r.ctx).MultiStatus(data...)
}

func (r *request[T]) AlreadyReported(data ...any) error {
	return response.Response(r.ctx).AlreadyReported(data...)
}

func (r *request[T]) IMUsed(data ...any) error {
	return response.Response(r.ctx).IMUsed(data...)
}
