package pg

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Filter = func(tx *gorm.DB) *gorm.DB
)

type pageEntity[T any] struct {
	Total int64
	Data  T `gorm:"embedded"`
}

type Model interface {
	TableName() string
}

var ErrRowsAffected = errors.New("not rows affected")

func NewReturning(columns ...string) clause.Returning {
	var clauseReturning clause.Returning
	{
		for _, column := range columns {
			clauseReturning.Columns = append(
				clauseReturning.Columns, clause.Column{
					Name: column,
				},
			)
		}
	}
	return clauseReturning
}

func IsTx(db *gorm.DB) bool {
	return db.Statement != nil && db.Statement.ConnPool != db.ConnPool
}

func Transaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db
	{
		if !IsTx(tx) {
			tx = tx.Begin()
		}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

type Point struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (Point) GormDataType() string {
	return "POINT"
}

func (p Point) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return gorm.Expr("POINT(?, ?)", p.Longitude, p.Latitude)
}

func (p *Point) Scan(v any) error {
	switch val := v.(type) {
	case string:
		_, err := fmt.Sscanf(val, "(%f,%f)", &p.Longitude, &p.Latitude)
		return err
	case []byte:
		_, err := fmt.Sscanf(string(val), "(%f,%f)", &p.Longitude, &p.Latitude)
		return err
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
