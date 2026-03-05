package response

type (
	ID64          = ID[int64]       //@name ID64
	IDStr         = ID[string]      //@name IDStr
	MessageString = Message[string] //@name MessageString
)

type HttpErrorResponse[T any] struct {
	Message T `json:"message"`
}

func ErrorResponse[T any](message T) *HttpErrorResponse[T] {
	return &HttpErrorResponse[T]{message}
}

type ID[T any] struct {
	ID T `json:"id" xml:"id"`
} //@name ID

func NewID[T any](id T) ID[T] {
	return ID[T]{
		ID: id,
	}
}

type Message[T any] struct {
	Message T `json:"message" xml:"message"`
} //@name Message

func NewMessage[T any](message T) Message[T] {
	return Message[T]{
		Message: message,
	}
}
