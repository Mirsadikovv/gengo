package {{ toSnake .EntityName }}_service

import (
	"context"

	{{ toSnake .EntityName }}_dto "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}/dto"
	{{ toSnake .EntityName }}_model "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"gorm.io/gorm"
)

type {{ toCamel .EntityName }}Service interface {
	Create(dto *{{ toSnake .EntityName }}_dto.Create{{ toCamel .EntityName }}) (int64, error)
	Update(dto *{{ toSnake .EntityName }}_dto.Update{{ toCamel .EntityName }}, filter pg.Filter) error
	Delete(filter pg.Filter) error
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}Page, error)
	Find(ctx context.Context, filter pg.Filter) ([]{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}, error)
	FindOne(ctx context.Context, filter pg.Filter) (*{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}, error)
}

func New{{ toCamel .EntityName }}Service(db *gorm.DB) {{ toCamel .EntityName }}Service {
	return &{{ toLowerCamel .EntityName }}Service{db: db}
}

type {{ toLowerCamel .EntityName }}Service struct {
	db *gorm.DB
}

func (s *{{ toLowerCamel .EntityName }}Service) Create(dto *{{ toSnake .EntityName }}_dto.Create{{ toCamel .EntityName }}) (int64, error) {

	model := &{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}{
		Name: dto.Name,
	}

	if err := pg.Create(s.db, model, "id"); err != nil {
		return 0, err
	}

	return model.Id, nil
}

func (s *{{ toLowerCamel .EntityName }}Service) Update(dto *{{ toSnake .EntityName }}_dto.Update{{ toCamel .EntityName }}, filter pg.Filter) error {

	updateDto := map[string]any{}
	{
		if dto.Name != nil {
			updateDto["name"] = dto.Name
		}

		if len(updateDto) == 0 {
			return nil
		}
	}

	_, err := pg.Update[{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}](s.db, updateDto, filter, "id")
	return err
}

func (s *{{ toLowerCamel .EntityName }}Service) Delete(filter pg.Filter) error {
	return pg.Delete[{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}](s.db, nil, filter)
}

func (s *{{ toLowerCamel .EntityName }}Service) Find(ctx context.Context, filter pg.Filter) ([]{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}, error) {
	return pg.FindWithScan[{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}, {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}](s.db, filter)
}

func (s *{{ toLowerCamel .EntityName }}Service) FindOne(ctx context.Context, filter pg.Filter) (*{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}, error) {
	return pg.FindOneWithScan[{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}, {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}](s.db, filter)
}

func (s *{{ toLowerCamel .EntityName }}Service) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}Page, error) {
	return pg.PageWithScan[{{ toSnake .EntityName }}_model.{{ toCamel .EntityName }}, {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}](s.db, paginate, filter)
}
