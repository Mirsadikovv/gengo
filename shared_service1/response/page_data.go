package response

type PageData[T any] struct {
	Total       int64 `json:"total" xml:"total"`
	TotalPages  int64 `json:"totalPages" xml:"totalPages"`
	PageSize    int   `json:"pageSize" xml:"pageSize"`
	CurrentPage int   `json:"currentPage" xml:"currentPage"`
	Data        []T   `json:"data" xml:"data"`
}
