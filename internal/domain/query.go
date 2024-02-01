package domain

type GetsQuery struct {
	SelectQuery
	FiltersQuery
	PaginationQuery
}

type SelectQuery struct {
	Selects []string
}

type FiltersQuery struct {
	Filters map[string][]string
}

type PaginationQuery struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
}
