package request

import "github.com/Mirsadikovv/gengo/shared_service/response"

func (r *request[T]) BadRequest(err ...any) error {
	return response.HTTPError(err...).BadRequest()
}

func (r *request[T]) Unauthorized(err ...any) error {
	return response.HTTPError(err...).Unauthorized()
}

func (r *request[T]) PaymentRequired(err ...any) error {
	return response.HTTPError(err...).PaymentRequired()
}

func (r *request[T]) Forbidden(err ...any) error {
	return response.HTTPError(err...).Forbidden()
}

func (r *request[T]) NotFound(err ...any) error {
	return response.HTTPError(err...).NotFound()
}

func (r *request[T]) MethodNotAllowed(err ...any) error {
	return response.HTTPError(err...).MethodNotAllowed()
}

func (r *request[T]) NotAcceptable(err ...any) error {
	return response.HTTPError(err...).NotAcceptable()
}

func (r *request[T]) ProxyAuthRequired(err ...any) error {
	return response.HTTPError(err...).ProxyAuthRequired()
}

func (r *request[T]) RequestTimeout(err ...any) error {
	return response.HTTPError(err...).RequestTimeout()
}

func (r *request[T]) Conflict(err ...any) error {
	return response.HTTPError(err...).Conflict()
}

func (r *request[T]) Gone(err ...any) error {
	return response.HTTPError(err...).Gone()
}

func (r *request[T]) LengthRequired(err ...any) error {
	return response.HTTPError(err...).LengthRequired()
}

func (r *request[T]) PreconditionFailed(err ...any) error {
	return response.HTTPError(err...).PreconditionFailed()
}

func (r *request[T]) RequestEntityTooLarge(err ...any) error {
	return response.HTTPError(err...).RequestEntityTooLarge()
}

func (r *request[T]) RequestURITooLong(err ...any) error {
	return response.HTTPError(err...).RequestURITooLong()
}

func (r *request[T]) UnsupportedMediaType(err ...any) error {
	return response.HTTPError(err...).UnsupportedMediaType()
}

func (r *request[T]) RequestedRangeNotSatisfiable(err ...any) error {
	return response.HTTPError(err...).RequestedRangeNotSatisfiable()
}

func (r *request[T]) SessionExpired(err ...any) error {
	return response.HTTPError(err...).SessionExpired()
}

func (r *request[T]) ExpectationFailed(err ...any) error {
	return response.HTTPError(err...).ExpectationFailed()
}

func (r *request[T]) MisdirectedRequest(err ...any) error {
	return response.HTTPError(err...).MisdirectedRequest()
}

func (r *request[T]) UnprocessableEntity(err ...any) error {
	return response.HTTPError(err...).UnprocessableEntity()
}

func (r *request[T]) Locked(err ...any) error {
	return response.HTTPError(err...).Locked()
}

func (r *request[T]) FailedDependency(err ...any) error {
	return response.HTTPError(err...).FailedDependency()
}

func (r *request[T]) TooEarly(err ...any) error {
	return response.HTTPError(err...).TooEarly()
}

func (r *request[T]) UpgradeRequired(err ...any) error {
	return response.HTTPError(err...).UpgradeRequired()
}

func (r *request[T]) PreconditionRequired(err ...any) error {
	return response.HTTPError(err...).PreconditionRequired()
}

func (r *request[T]) TooManyRequests(err ...any) error {
	return response.HTTPError(err...).TooManyRequests()
}

func (r *request[T]) RequestHeaderFieldsTooLarge(err ...any) error {
	return response.HTTPError(err...).RequestHeaderFieldsTooLarge()
}

func (r *request[T]) UnavailableForLegalReasons(err ...any) error {
	return response.HTTPError(err...).UnavailableForLegalReasons()
}

func (r *request[T]) InternalServerError(err ...any) error {
	return response.HTTPError(err...).InternalServerError()
}

func (r *request[T]) NotImplemented(err ...any) error {
	return response.HTTPError(err...).NotImplemented()
}

func (r *request[T]) BadGateway(err ...any) error {
	return response.HTTPError(err...).BadGateway()
}

func (r *request[T]) ServiceUnavailable(err ...any) error {
	return response.HTTPError(err...).ServiceUnavailable()
}

func (r *request[T]) GatewayTimeout(err ...any) error {
	return response.HTTPError(err...).GatewayTimeout()
}

func (r *request[T]) HTTPVersionNotSupported(err ...any) error {
	return response.HTTPError(err...).HTTPVersionNotSupported()
}

func (r *request[T]) VariantAlsoNegotiates(err ...any) error {
	return response.HTTPError(err...).VariantAlsoNegotiates()
}

func (r *request[T]) InsufficientStorage(err ...any) error {
	return response.HTTPError(err...).InsufficientStorage()
}

func (r *request[T]) LoopDetected(err ...any) error {
	return response.HTTPError(err...).LoopDetected()
}

func (r *request[T]) NotExtended(err ...any) error {
	return response.HTTPError(err...).NotExtended()
}

func (r *request[T]) NetworkAuthenticationRequired(err ...any) error {
	return response.HTTPError(err...).NetworkAuthenticationRequired()
}
