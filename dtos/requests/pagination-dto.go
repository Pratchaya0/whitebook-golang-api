package requests

type PaginationDto struct {
	Limit      int         `form:"limit,omitempty;query:limit"`
	Page       int         `form:"page,omitempty;query:page"`
	Sort       string      `form:"sort,omitempty;query:sort"`
	TotalRows  int64       `form:"total_rows"`
	TotalPages int         `form:"total_pages"`
	Rows       interface{} `form:"rows"`
}

func (p *PaginationDto) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationDto) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PaginationDto) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginationDto) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
