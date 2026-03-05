package request

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	min         = 10
	max         = 100
	pageName    = "page"
	perpageName = "perpage"
)

func SetPageName(v string) {
	pageName = v
}

func SetPerPageName(v string) {
	perpageName = v
}

func SetMin(v int) {
	min = v
}

func SetMax(v int) {
	max = v
}

type Paginate struct{ page, perpage, min, max int }

func (p *Paginate) SetMin(v int) {
	p.min = v
}

func (p *Paginate) SetMax(v int) {
	p.max = v
}

func (p *Paginate) Offset() int {
	return p.PerPage() * (p.Page() - 1)
}

func (p *Paginate) Limit() int {
	return p.PerPage()
}

func (p *Paginate) Page() int {

	if p.page < 1 {
		p.page = 1
	}

	return p.page
}

func (p *Paginate) PerPage() int {

	if p.perpage < 1 {
		p.perpage = min
	} else if p.perpage > max {
		p.perpage = max
	}

	return p.perpage
}

func NewPaginate(page, perPage int) *Paginate {
	return &Paginate{
		page:    page,
		perpage: perPage,
		min:     min,
		max:     max,
	}
}

func NewPaginateEchoWithContext(c echo.Context) *Paginate {

	page, _ := strconv.Atoi(c.QueryParam(pageName))

	perPage, _ := strconv.Atoi(c.QueryParam(perpageName))

	return NewPaginate(page, perPage)
}
