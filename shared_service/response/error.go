package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpError interface {
	Code(int) HttpError
	Send() error
	send() error
	// 400 - Bad Request
	BadRequest() error
	// 401 - Unauthorized
	Unauthorized() error
	// 402 - Payment Required
	PaymentRequired() error
	// 403 - Forbidden
	Forbidden() error
	// 404 - Not Found
	NotFound() error
	// 405 - Method Not Allowed
	MethodNotAllowed() error
	// 406 - Not Acceptable
	NotAcceptable() error
	// 407 - Proxy Authentication Required
	ProxyAuthRequired() error
	// 408 - Request Timeout
	RequestTimeout() error
	// 409 - Conflict
	Conflict() error
	// 410 - Gone
	Gone() error
	// 411 - Length Required
	LengthRequired() error
	// 412 - Precondition Failed
	PreconditionFailed() error
	// 413 - Request Entity Too Large
	RequestEntityTooLarge() error
	// 414 - Request URI Too Long
	RequestURITooLong() error
	// 415 - Unsupported Media Type
	UnsupportedMediaType() error
	// 416 - Requested Range Not Satisfiable
	RequestedRangeNotSatisfiable() error
	// 417 - Expectation Failed
	// 419 - Session Expired
	SessionExpired() error
	ExpectationFailed() error
	// 421 - Misdirected Request
	MisdirectedRequest() error
	// 422 - Unprocessable Entity
	UnprocessableEntity() error
	// 423 - Locked
	Locked() error
	// 424 - Failed Dependency
	FailedDependency() error
	// 425 - Too Early
	TooEarly() error
	// 426 - Upgrade Required
	UpgradeRequired() error
	// 428 - Precondition Required
	PreconditionRequired() error
	// 429 - Too Many Requests
	TooManyRequests() error
	// 431 - Request Header Fields Too Large
	RequestHeaderFieldsTooLarge() error
	// 451 - Unavailable For Legal Reasons
	UnavailableForLegalReasons() error
	// 500 - Internal Server Error
	InternalServerError() error
	// 501 - Not Implemented
	NotImplemented() error
	// 502 - Bad Gateway
	BadGateway() error
	// 503 - Service Unavailable
	ServiceUnavailable() error
	// 504 - Gateway Timeout
	GatewayTimeout() error
	// 505 - HTTP Version Not Supported
	HTTPVersionNotSupported() error
	// 506 - Variant Also Negotiates
	VariantAlsoNegotiates() error
	// 507 - Insufficient Storage
	InsufficientStorage() error
	// 508 - Loop Detected
	LoopDetected() error
	// 510 - Not Extended
	NotExtended() error
	// 511 - Network Authentication Required
	NetworkAuthenticationRequired() error
}

func HTTPError(msg ...any) HttpError {
	h := httpError{
		status: http.StatusInternalServerError,
	}

	if len(msg) > 0 {
		h.err = msg[0]
	}

	return &h
}

type httpError struct {
	status int
	err    any
}

func (h *httpError) Code(status int) HttpError {
	h.status = status
	return h
}

func (h *httpError) Send() error {
	return h.send()
}

func (h *httpError) send() error {
	err := h.err
	{
		if err == nil {
			return echo.NewHTTPError(h.status)
		}
	}
	switch e := err.(type) {
	case error:
		return echo.NewHTTPError(h.status, ErrorResponse(e.Error()))
	default:
		return echo.NewHTTPError(h.status, ErrorResponse(err))
	}
}

func (h *httpError) BadRequest() error {
	return h.Code(http.StatusBadRequest).send()
}

func (h *httpError) Unauthorized() error {
	return h.Code(http.StatusUnauthorized).send()
}

func (h *httpError) PaymentRequired() error {
	return h.Code(http.StatusPaymentRequired).send()
}

func (h *httpError) Forbidden() error {
	return h.Code(http.StatusForbidden).send()
}

func (h *httpError) NotFound() error {
	return h.Code(http.StatusNotFound).send()
}

func (h *httpError) MethodNotAllowed() error {
	return h.Code(http.StatusMethodNotAllowed).send()
}

func (h *httpError) NotAcceptable() error {
	return h.Code(http.StatusNotAcceptable).send()
}

func (h *httpError) ProxyAuthRequired() error {
	return h.Code(http.StatusProxyAuthRequired).send()
}

func (h *httpError) RequestTimeout() error {
	return h.Code(http.StatusRequestTimeout).send()
}

func (h *httpError) Conflict() error {
	return h.Code(http.StatusConflict).send()
}

func (h *httpError) Gone() error {
	return h.Code(http.StatusGone).send()
}

func (h *httpError) LengthRequired() error {
	return h.Code(http.StatusLengthRequired).send()
}

func (h *httpError) PreconditionFailed() error {
	return h.Code(http.StatusPreconditionFailed).send()
}

func (h *httpError) RequestEntityTooLarge() error {
	return h.Code(http.StatusRequestEntityTooLarge).send()
}

func (h *httpError) RequestURITooLong() error {
	return h.Code(http.StatusRequestURITooLong).send()
}

func (h *httpError) UnsupportedMediaType() error {
	return h.Code(http.StatusUnsupportedMediaType).send()
}

func (h *httpError) RequestedRangeNotSatisfiable() error {
	return h.Code(http.StatusRequestedRangeNotSatisfiable).send()
}

func (h *httpError) ExpectationFailed() error {
	return h.Code(http.StatusExpectationFailed).send()
}

func (h *httpError) Teapot() error {
	return h.Code(http.StatusTeapot).send()
}

func (h *httpError) SessionExpired() error {
	return h.Code(419).send()
}

func (h *httpError) MisdirectedRequest() error {
	return h.Code(http.StatusMisdirectedRequest).send()
}

func (h *httpError) UnprocessableEntity() error {
	return h.Code(http.StatusUnprocessableEntity).send()
}

func (h *httpError) Locked() error {
	return h.Code(http.StatusLocked).send()
}

func (h *httpError) FailedDependency() error {
	return h.Code(http.StatusFailedDependency).send()
}

func (h *httpError) TooEarly() error {
	return h.Code(http.StatusTooEarly).send()
}

func (h *httpError) UpgradeRequired() error {
	return h.Code(http.StatusUpgradeRequired).send()
}

func (h *httpError) PreconditionRequired() error {
	return h.Code(http.StatusPreconditionRequired).send()
}

func (h *httpError) TooManyRequests() error {
	return h.Code(http.StatusTooManyRequests).send()
}

func (h *httpError) RequestHeaderFieldsTooLarge() error {
	return h.Code(http.StatusRequestHeaderFieldsTooLarge).send()
}

func (h *httpError) UnavailableForLegalReasons() error {
	return h.Code(http.StatusUnavailableForLegalReasons).send()
}

func (h *httpError) InternalServerError() error {
	return h.Code(http.StatusInternalServerError).send()
}

func (h *httpError) NotImplemented() error {
	return h.Code(http.StatusNotImplemented).send()
}

func (h *httpError) BadGateway() error {
	return h.Code(http.StatusBadGateway).send()
}

func (h *httpError) ServiceUnavailable() error {
	return h.Code(http.StatusServiceUnavailable).send()
}

func (h *httpError) GatewayTimeout() error {
	return h.Code(http.StatusGatewayTimeout).send()
}

func (h *httpError) HTTPVersionNotSupported() error {
	return h.Code(http.StatusHTTPVersionNotSupported).send()
}

func (h *httpError) VariantAlsoNegotiates() error {
	return h.Code(http.StatusVariantAlsoNegotiates).send()
}

func (h *httpError) InsufficientStorage() error {
	return h.Code(http.StatusInsufficientStorage).send()
}

func (h *httpError) LoopDetected() error {
	return h.Code(http.StatusLoopDetected).send()
}

func (h *httpError) NotExtended() error {
	return h.Code(http.StatusNotExtended).send()
}

func (h *httpError) NetworkAuthenticationRequired() error {
	return h.Code(http.StatusNetworkAuthenticationRequired).send()
}
