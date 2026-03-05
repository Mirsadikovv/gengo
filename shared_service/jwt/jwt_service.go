package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IUser[T, E, K any] interface {
	ID() int64
	Pre(T, E, ...K) (bool, error)
}

type Payload[U IUser[T, E, K], T, E, K any] struct {
	User U `json:"user" xml:"user"`
	jwt.RegisteredClaims
}

type JwtService[U IUser[T, E, K], T, E, K any] interface {
	ParseToken(string) (U, error)
	ParseTokenWithExpired(string) (U, error)
	ParseTokenWithGracePeriod(token string, grace int64) (U, error)
	Token(U) (string, error)
	Config() *JwtConfig
}

type jwtService[U IUser[T, E, K], T, E, K any] struct {
	config *JwtConfig
}

func (j *jwtService[U, T, E, K]) Token(user U) (string, error) {
	return Encode[U](user, j.config.Secret, j.config.Expired)
}

func (j *jwtService[U, T, E, K]) ParseToken(token string) (U, error) {
	payload, err := Decode[U](token, j.config.Secret, false)
	{
		if err != nil {
			var none U
			return none, err
		}
	}
	return payload.User, nil
}

func (j *jwtService[U, T, E, K]) ParseTokenWithExpired(token string) (U, error) {
	payload, err := Decode[U](token, j.config.Secret, true)
	{
		if err != nil {
			var none U
			return none, err
		}
	}
	return payload.User, nil
}

func (j *jwtService[U, T, E, K]) ParseTokenWithGracePeriod(token string, grace int64) (U, error) {
	payload, err := Decode[U](token, j.config.Secret, true)
	{
		if err != nil {
			var none U
			return none, err
		}
	}

	if payload.ExpiresAt.Add(time.Duration(grace) * time.Minute).Before(time.Now()) {
		var none U
		return none, errors.New("token expired")
	}

	return payload.User, nil
}

func (j *jwtService[U, T, E, K]) Config() *JwtConfig {
	return j.config
}

func NewJwtService[U IUser[T, E, K], T, E, K any](config *JwtConfig) JwtService[U, T, E, K] {
	return &jwtService[U, T, E, K]{
		config: config,
	}
}

func Encode[U IUser[T, E, K], T, E, K any](user U, secret string, expired int64) (string, error) {
	payload := NewPayload(user, expired)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(secret))
	{
		if err != nil {
			return "", err
		}
	}

	return tokenString, nil
}

func Decode[U IUser[T, E, K], T, E, K any](tokStr, secret string, withExpired bool) (*Payload[U, T, E, K], error) {
	keyfunc := func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}

	token, err := jwt.ParseWithClaims(tokStr, &Payload[U, T, E, K]{}, keyfunc)
	{
		if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}

		if withExpired && !token.Valid {
			return nil, errors.New("token is not valid")
		}
	}

	payload, ok := token.Claims.(*Payload[U, T, E, K])
	{
		if !ok {
			return nil, errors.New("unknown error")
		}
	}

	return payload, nil
}

func NewPayload[U IUser[T, E, K], T, E, K any](user U, expired int64) *Payload[U, T, E, K] {
	now := time.Now()

	exp := now.Add(time.Minute * time.Duration(expired))

	return &Payload[U, T, E, K]{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatInt(user.ID(), 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}
