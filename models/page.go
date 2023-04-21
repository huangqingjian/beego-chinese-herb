package models

// 分页查询
type PageQuery struct {
	PageNum   int32		// 当前页
	PageSize  int32		// 每页的数量
}

// 分页
type Page struct {
	PageNum  int32      		`json:"pageNum"`   // 当前页
	PageSize int32				`json:"pageSize"`  // 每页的数量
	Size 	 int32				`json:"size"`      // 当前页的数量
	Total    int32			    `json:"total"`     // 总记录数
	Pages    int32				`json:"pages"`     // 总页数
	List     interface{}   	    `json:"list"`      // 结果集
}

// 创建分页
func NewPage(pageNum int32, pageSize int32, size int32, total int32, list interface{}) *Page {
	page := Page{}
	page.PageNum = pageNum
	page.PageSize = pageSize
	page.Size = size
	if size == 0 {
		page.Total = 0
		page.Pages = 0
	} else {
		page.Total = total
		if total % pageSize == 0 {
			page.Pages = total / pageSize
		} else {
			page.Pages = total / pageSize + 1
		}
		page.List = list
	}
	return &page
}
