package entity

type Pagination struct {
	PageSize  int32
	Page      int32
	Total     int64
	TotalPage int32
}

func (p *Pagination) ValidatePagination() {
	if p.Total <= 0 || p.PageSize <= 0 {
		return
	}

	p.TotalPage = int32(p.Total) / p.PageSize
	if int32(p.Total)%p.PageSize != 0 {
		p.TotalPage++
	}

	if p.Page <= 0 || p.Page > p.TotalPage {
		return
	}
}
