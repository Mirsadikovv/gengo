package {{ toSnake .EntityName }}_handler

import (
	auth_middleware "{{ .ModPath }}/src/module/auth_service/middleware"
	{{ toSnake .EntityName }}_dto "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}/dto"
	{{ toSnake .EntityName }}_service "{{ .ModPath }}/src/module/{{ toSnake .ModuleName }}/service"

	"github.com/Mirsadikovv/gengo/shared_service/logger"
	"github.com/Mirsadikovv/gengo/shared_service/request"
	"github.com/Mirsadikovv/gengo/shared_service/response"
	"github.com/Mirsadikovv/gengo/shared_service/sharedutil"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type {{ toLowerCamel .EntityName }}Handler struct {
	db                      *gorm.DB
	log                     logger.Logger
	authMiddleware          *auth_middleware.AuthMiddleware
	{{ toLowerCamel .EntityName }}Service {{ toSnake .EntityName }}_service.{{ toCamel .EntityName }}Service
}

func New{{ toCamel .EntityName }}Handler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &{{ toLowerCamel .EntityName }}Handler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		{{ toLowerCamel .EntityName }}Service: {{ toSnake .EntityName }}_service.New{{ toCamel .EntityName }}Service(db),
	}

	g := group.Group("/{{ toSnake .EntityName }}")
	{
		g.POST("/create", handler.Create)
		g.PATCH("/:id", handler.Update)
		g.DELETE("/:id", handler.Delete)
		g.GET("/page", handler.Page)
		g.GET("/search", handler.Search)
		g.GET("/:id", handler.GetByID)
	}
}

// Create
// @Summary      Create {{ toCamel .EntityName }}
// @Tags         {{ toCamel .EntityName }}
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body {{ toSnake .EntityName }}_dto.Create{{ toCamel .EntityName }} true "Create"
// @Success      201 {object} response.ID64
// @Failure      400 {object} response.HttpSuccess
// @Failure      500 {object} response.HttpSuccess
// @Router       /{{ toSnake .EntityName }}/create [post]
func (h *{{ toLowerCamel .EntityName }}Handler) Create(ctx echo.Context) error {

	req := request.Request(ctx)

	var dto {{ toSnake .EntityName }}_dto.Create{{ toCamel .EntityName }}
	{
		if err := req.BindBody(&dto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.{{ toLowerCamel .EntityName }}Service.Create(&dto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update
// @Summary      Update {{ toCamel .EntityName }}
// @Tags         {{ toCamel .EntityName }}
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id    path int                                       true "ID"
// @Param        input body {{ toSnake .EntityName }}_dto.Update{{ toCamel .EntityName }} true "Update"
// @Success      204 {object} response.HttpSuccess
// @Failure      400 {object} response.HttpSuccess
// @Failure      500 {object} response.HttpSuccess
// @Router       /{{ toSnake .EntityName }}/{id} [patch]
func (h *{{ toLowerCamel .EntityName }}Handler) Update(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var dto {{ toSnake .EntityName }}_dto.Update{{ toCamel .EntityName }}
	{
		if err := req.BindBody(&dto); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := h.{{ toLowerCamel .EntityName }}Service.Update(&dto, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Delete
// @Summary      Delete {{ toCamel .EntityName }}
// @Tags         {{ toCamel .EntityName }}
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path int true "ID"
// @Success      204 {object} response.HttpSuccess
// @Failure      400 {object} response.HttpSuccess
// @Failure      500 {object} response.HttpSuccess
// @Router       /{{ toSnake .EntityName }}/{id} [delete]
func (h *{{ toLowerCamel .EntityName }}Handler) Delete(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := h.{{ toLowerCamel .EntityName }}Service.Delete(filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Page
// @Summary      {{ toCamel .EntityName }} Page
// @Tags         {{ toCamel .EntityName }}
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page  query int                                                  false "Page"
// @Param        size  query int                                                  false "Size"
// @Param        query query {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}QueryParams false "Query"
// @Success      200 {object} {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}Page
// @Failure      400 {object} response.HttpSuccess
// @Failure      500 {object} response.HttpSuccess
// @Router       /{{ toSnake .EntityName }}/page [get]
func (h *{{ toLowerCamel .EntityName }}Handler) Page(ctx echo.Context) error {

	req := request.Request(ctx)

	var queryParams {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}QueryParams
	{
		if err := req.BindQuery(&queryParams); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if queryParams.Name != "" {
			tx = tx.Where("name ILIKE ?", sharedutil.Join("%", queryParams.Name, "%"))
		}
		return tx
	}

	page, err := h.{{ toLowerCamel .EntityName }}Service.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(page)
}

// Search
// @Summary      {{ toCamel .EntityName }} Search
// @Tags         {{ toCamel .EntityName }}
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        query query {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}QueryParams false "Query"
// @Success      200 {object} []{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}
// @Failure      400 {object} response.HttpSuccess
// @Failure      500 {object} response.HttpSuccess
// @Router       /{{ toSnake .EntityName }}/search [get]
func (h *{{ toLowerCamel .EntityName }}Handler) Search(ctx echo.Context) error {

	req := request.Request(ctx)

	var queryParams {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}QueryParams
	{
		if err := req.BindQuery(&queryParams); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if queryParams.Name != "" {
			tx = tx.Where("name ILIKE ?", sharedutil.Join("%", queryParams.Name, "%"))
		}
		return tx
	}

	items, err := h.{{ toLowerCamel .EntityName }}Service.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}

		if len(items) == 0 {
			return req.OK([]{{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}{})
		}
	}

	return req.OK(items)
}

// GetByID
// @Summary      Get {{ toCamel .EntityName }} by id
// @Tags         {{ toCamel .EntityName }}
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path int true "ID"
// @Success      200 {object} {{ toSnake .EntityName }}_dto.{{ toCamel .EntityName }}
// @Failure      400 {object} response.HttpSuccess
// @Failure      500 {object} response.HttpSuccess
// @Router       /{{ toSnake .EntityName }}/{id} [get]
func (h *{{ toLowerCamel .EntityName }}Handler) GetByID(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	item, err := h.{{ toLowerCamel .EntityName }}Service.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(item)
}
