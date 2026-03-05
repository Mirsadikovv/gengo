package pg

import (
	"math"

	"github.com/Mirsadikovv/gengo/shared_service/request"
	"github.com/Mirsadikovv/gengo/shared_service/response"
	"gorm.io/gorm"
)

func query[T any](db *gorm.DB, filter ...Filter) *gorm.DB {
	return db.Model(new(T)).Scopes(filter...)
}

func page[T any, E any](db *gorm.DB, paginate *request.Paginate, filter ...Filter) *gorm.DB {
	totalFilter := func(tx *gorm.DB) *gorm.DB {
		selects := tx.Statement.Selects
		{
			if len(selects) == 0 {
				selects = []string{
					"*",
					"COUNT(1) OVER() AS total",
				}
			} else {
				selects = append(
					selects,
					"COUNT(1) OVER() AS total",
				)
			}
		}

		tx.Statement.Selects = selects

		return tx
	}

	tx := query[T](db, append(filter, totalFilter)...)

	return tx.Offset(paginate.Offset()).
		Limit(paginate.Limit())
}

func pageResult[T any](pageEntities []pageEntity[T], paginate *request.Paginate) *response.PageData[T] {
	var (
		total      int64
		totalPages int64
	)

	if len(pageEntities) > 0 {
		pageEntity := pageEntities[0]
		total = pageEntity.Total
		totalPages = int64(math.Ceil(float64(total) / float64(paginate.Limit())))
	}

	entities := make([]T, 0, len(pageEntities))
	{
		for _, pageEntity := range pageEntities {
			entities = append(entities, pageEntity.Data)
		}
	}

	return &response.PageData[T]{
		Total:       total,
		TotalPages:  totalPages,
		PageSize:    paginate.Limit(),
		CurrentPage: paginate.Page(),
		Data:        entities,
	}
}

func FindOne[T any](db *gorm.DB, filter ...Filter) (T, error) {
	var entity T
	{
		if err := query[T](db, filter...).
			First(&entity).Error; err != nil {
			return entity, err
		}
	}

	return entity, nil
}

func Find[T any](db *gorm.DB, filter ...Filter) ([]T, error) {
	var entites []T
	{
		if err := query[T](db, filter...).
			Find(&entites).Error; err != nil {
			return nil, err
		}
	}

	if entites == nil {
		entites = []T{}
	}

	return entites, nil
}

func Page[T any](db *gorm.DB, paginate *request.Paginate, filter ...Filter) (*response.PageData[T], error) {
	var pageEntities []pageEntity[T]
	{
		result := page[T, T](db, paginate, filter...).
			Scan(&pageEntities)

		if err := result.Error; err != nil {
			return nil, err
		}
	}

	return pageResult(pageEntities, paginate), nil
}

func Create[T any](db *gorm.DB, entity *T, columns ...string) error {
	var tx *gorm.DB
	{
		if len(columns) > 0 {
			tx = db.Clauses(NewReturning(columns...))
		} else {
			tx = db
		}
	}

	result := query[T](tx).Create(entity)
	{
		if err := result.Error; err != nil {
			return err
		}
	}

	return nil
}

func Update[T any, E any](db *gorm.DB, dto E, filter Filter, columns ...string) (*T, error) {
	var tx *gorm.DB
	{
		if len(columns) > 0 {
			tx = db.Clauses(NewReturning(columns...))
		} else {
			tx = db
		}
	}

	model := new(T)
	{

		result := tx.Model(model).Scopes(filter).Updates(dto)
		{
			if err := result.Error; err != nil {
				return nil, err
			}

			if result.RowsAffected == 0 {
				return nil, gorm.ErrRecordNotFound
			}
		}

	}

	return model, nil
}

func Delete[T any](db *gorm.DB, entity *T, filter Filter, columns ...string) error {
	var tx *gorm.DB
	{
		if len(columns) > 0 {
			tx = db.Clauses(NewReturning(columns...))
		} else {
			tx = db
		}
	}

	if entity == nil {
		entity = new(T)
	}

	result := query[T](tx, filter).Delete(entity)
	{
		if err := result.Error; err != nil {
			return err
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}

	return nil
}

func FindOneWithScan[T any, E any](db *gorm.DB, filter ...Filter) (*E, error) {
	var entity = new(E)
	{
		result := query[T](db, filter...).Scan(entity)
		{
			if err := result.Error; err != nil {
				return nil, err
			}

			if result.RowsAffected == 0 {
				return nil, gorm.ErrRecordNotFound
			}
		}

	}

	return entity, nil
}

func FindWithScan[T any, E any](db *gorm.DB, filter ...Filter) ([]E, error) {
	var entites []E
	{
		result := query[T](db, filter...).Scan(&entites)
		{
			if err := result.Error; err != nil {
				return nil, err
			}
		}

	}

	if entites == nil {
		entites = []E{}
	}

	return entites, nil
}

func PageWithScan[T any, E any](db *gorm.DB, paginate *request.Paginate, filter ...Filter) (*response.PageData[E], error) {
	var pageEntities []pageEntity[E]
	{
		result := page[T, E](db, paginate, filter...).
			Scan(&pageEntities)

		if err := result.Error; err != nil {
			return nil, err
		}
	}

	return pageResult(pageEntities, paginate), nil
}

func Count[T any](db *gorm.DB, filter ...Filter) (int64, error) {
	var count int64
	{
		tx := query[T](db, filter...)

		if err := tx.Count(&count).Error; err != nil {
			return 0, err
		}
	}

	return count, nil
}

func Exists[T any](db *gorm.DB, filter ...Filter) (bool, error) {
	count, err := Count[T](db, filter...)
	{
		if err != nil {
			return false, err
		}
	}

	return count > 0, err
}
