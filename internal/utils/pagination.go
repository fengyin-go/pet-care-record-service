// Package pagination 提供分页相关工具函数。
package utils

import (
	"fmt"
	"math"
)

// PageInfo 分页信息结构。
type PageInfo struct {
	Page       int  `json:"page"`
	Size       int  `json:"size"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// NewPageInfo 创建分页信息。
func NewPageInfo(page, size, total int) *PageInfo {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	totalPages := int(math.Ceil(float64(total) / float64(size)))
	if totalPages < 1 {
		totalPages = 1
	}
	return &PageInfo{
		Page:       page,
		Size:       size,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}

// Offset 返回数据库偏移量。
func (p *PageInfo) Offset() int {
	return (p.Page - 1) * p.Size
}

// Limit 返回每页数量。
func (p *PageInfo) Limit() int {
	return p.Size
}

// NextPage 返回下一页页码，如果没有下一页返回当前页。
func (p *PageInfo) NextPage() int {
	if p.HasNext {
		return p.Page + 1
	}
	return p.Page
}

// PrevPage 返回上一页页码，如果没有上一页返回当前页。
func (p *PageInfo) PrevPage() int {
	if p.HasPrev {
		return p.Page - 1
	}
	return p.Page
}

// PageRange 返回当前页显示的项目索引范围（从1开始）。
func (p *PageInfo) PageRange() (start, end int) {
	start = p.Offset() + 1
	end = p.Offset() + p.Size
	if end > p.Total {
		end = p.Total
	}
	if start > end {
		start = end
	}
	return start, end
}

// PaginatedResult 带分页的结果结构。
type PaginatedResult[T any] struct {
	Items    []T       `json:"items"`
	PageInfo *PageInfo `json:"page_info"`
}

// NewPaginatedResult 创建带分页的结果。
func NewPaginatedResult[T any](items []T, page, size, total int) *PaginatedResult[T] {
	return &PaginatedResult[T]{
		Items:    items,
		PageInfo: NewPageInfo(page, size, total),
	}
}

// EmptyPaginatedResult 创建空的分页结果。
func EmptyPaginatedResult[T any](page, size int) *PaginatedResult[T] {
	return NewPaginatedResult([]T{}, page, size, 0)
}

// SlicePage 对切片进行分页。
func SlicePage[T any](items []T, page, size int) ([]T, int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	total := len(items)
	start := (page - 1) * size
	if start >= total {
		return []T{}, total
	}
	end := start + size
	if end > total {
		end = total
	}
	return items[start:end], total
}

// NormalizePageParams 规范化分页参数。
func NormalizePageParams(page, size, defaultSize, maxSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = defaultSize
	}
	if size > maxSize {
		size = maxSize
	}
	return page, size
}

// PageNumbers 返回页码列表用于导航。
func PageNumbers(currentPage, totalPages, maxVisible int) []int {
	if totalPages <= 0 {
		return []int{}
	}
	if maxVisible < 1 {
		maxVisible = 5
	}
	if totalPages <= maxVisible {
		result := make([]int, totalPages)
		for i := 0; i < totalPages; i++ {
			result[i] = i + 1
		}
		return result
	}
	half := maxVisible / 2
	start := currentPage - half
	end := currentPage + half
	if maxVisible%2 == 0 {
		end--
	}
	if start < 1 {
		start = 1
		end = maxVisible
	}
	if end > totalPages {
		end = totalPages
		start = totalPages - maxVisible + 1
	}
	result := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

// FormatPaginationSummary 格式化分页摘要。
func FormatPaginationSummary(page, size, total int) string {
	start := (page-1)*size + 1
	end := start + size - 1
	if end > total {
		end = total
	}
	if total == 0 {
		return "无记录"
	}
	return fmt.Sprintf("第 %d-%d 条，共 %d 条", start, end, total)
}
