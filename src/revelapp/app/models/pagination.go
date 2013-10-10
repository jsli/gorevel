package models

import (
	"fmt"
	"html/template"
)

const (
	PagesPerView = 11 //每次显示的页码数
	ItemsPerPage = 10 //每页几条记录
)

type Pagination struct {
	page      int //当前页码
	total     int //记录总数
	url       string
	pageTotal int //总页数
}

func NewPagination(page int, total int, url string) *Pagination {
	p := Pagination{}
	p.page = page
	p.total = total
	p.url = url
	return &p
}

func (p *Pagination) Html() template.HTML {
	html := ""
	p.pageTotal = p.total / ItemsPerPage
	if p.pageTotal*ItemsPerPage < p.total {
		p.pageTotal += 1
	}
	if p.pageTotal == 1 {
		return template.HTML(html)
	}

	page := p.page
	page -= PagesPerView / 2
	if page < 0 {
		page = 0
	}

	count := page + PagesPerView
	if count > p.pageTotal {
		count = p.pageTotal
	}

	pageNum := 0
	for ; page < count; page++ {
		pageNum = page + 1
		if page != p.page {
			html += fmt.Sprintf(`<li><a href="%s%d">%d</a></li>`, p.url, pageNum, pageNum)
		} else {
			html += fmt.Sprintf(`<li class="active"><a href="#">%d</a></li>`, pageNum)
		}
	}

	return template.HTML(html)
}
