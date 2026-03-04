package {{ toSnake .EntityName }}_dto

import (
	"git.sriss.uz/shared/shared_service/response"
	"time"
)

type {{ toCamel .EntityName }}Page = response.PageData[{{ toCamel .EntityName }}]

type {{ toCamel .EntityName }} struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
} // @name {{ toCamel .EntityName }}

type Create{{ toCamel .EntityName }} struct {
	Name string `json:"name" validate:"required"`
} // @name Create{{ toCamel .EntityName }}

type Update{{ toCamel .EntityName }} struct {
	Name *string `json:"name"`
} // @name Update{{ toCamel .EntityName }}

type {{ toCamel .EntityName }}QueryParams struct {
	Name string `json:"name" query:"name"`
} // @name {{ toCamel .EntityName }}QueryParams
