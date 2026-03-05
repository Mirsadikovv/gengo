package pg

import (
	"github.com/Mirsadikovv/gengo/shared_service/request"
	"github.com/Mirsadikovv/gengo/shared_service/response"
	"gorm.io/gorm"
)

var _ Crud[any] = (*curd[any])(nil)

type CreateInterface[T any] interface {
	Create(data *T, columns ...string) (*T, error)
}

type DeleteInterface[T any] interface {
	Delete(filter Filter, columns ...string) (*T, error)
}

type FindInterface[T any] interface {
	Find(filter ...Filter) ([]T, error)
	FindOne(filter ...Filter) (T, error)
	Page(paginate *request.Paginate, filter ...Filter) (*response.PageData[T], error)
}

type UpdateInterface[T any] interface {
	Update(dto any, filter Filter, columns ...string) (*T, error)
}

type CreateWithScanInterface[T, E any] interface {
	CreateWithScan(data *T, columns ...string) (*T, error)
}

type UpdateWithScanInterface[T, E any] interface {
	UpdateWithScan(dto any, filter Filter, columns ...string) (*T, error)
}

type FindWithScanInterface[T, E any] interface {
	FindWithScan(filter ...Filter) ([]T, error)
	FindOneWithScan(filter ...Filter) (T, error)
	PageWithScan(paginate *request.Paginate, filter ...Filter) (*response.PageData[T], error)
}

type Crud[T any] interface {
	CreateInterface[T]
	UpdateInterface[T]
	FindInterface[T]
	DeleteInterface[T]
	DB() *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Crud[T] {
	return &curd[T]{db}
}

type curd[T any] struct {
	db *gorm.DB
}

func (r *curd[T]) Create(data *T, columns ...string) (*T, error) {
	if err := Create[T](r.db, data, columns...); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *curd[T]) Delete(filter Filter, columns ...string) (*T, error) {
	model := new(T)
	{
		if err := Delete[T](r.db, model, filter, columns...); err != nil {
			return nil, err
		}
	}

	return model, nil
}

func (r *curd[T]) Find(filter ...Filter) ([]T, error) {
	return Find[T](r.db, filter...)
}

func (r *curd[T]) FindOne(filter ...Filter) (T, error) {
	return FindOne[T](r.db, filter...)
}

func (r *curd[T]) Page(paginate *request.Paginate, filter ...Filter) (*response.PageData[T], error) {
	return Page[T](r.db, paginate, filter...)
}

func (r *curd[T]) Update(dto any, filter Filter, columns ...string) (*T, error) {
	return Update[T](r.db, dto, filter, columns...)
}

func (r *curd[T]) DB() *gorm.DB {
	return r.db
}
