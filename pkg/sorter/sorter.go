// Package sorter 提供通用排序工具。
package sorter

import (
	"sort"
	"strings"
)

// StringSorter 字符串排序器。
type StringSorter struct {
	items []string
	asc   bool
}

// NewStringSorter 创建字符串排序器。
func NewStringSorter(items []string, asc bool) *StringSorter {
	copied := make([]string, len(items))
	copy(copied, items)
	return &StringSorter{items: copied, asc: asc}
}

// Sort 执行排序。
func (ss *StringSorter) Sort() []string {
	if ss.asc {
		sort.Strings(ss.items)
	} else {
		sort.Slice(ss.items, func(i, j int) bool {
			return ss.items[i] > ss.items[j]
		})
	}
	return ss.items
}

// IntSorter 整数排序器。
type IntSorter struct {
	items []int
	asc   bool
}

// NewIntSorter 创建整数排序器。
func NewIntSorter(items []int, asc bool) *IntSorter {
	copied := make([]int, len(items))
	copy(copied, items)
	return &IntSorter{items: copied, asc: asc}
}

// Sort 执行排序。
func (is *IntSorter) Sort() []int {
	if is.asc {
		sort.Ints(is.items)
	} else {
		sort.Slice(is.items, func(i, j int) bool {
			return is.items[i] > is.items[j]
		})
	}
	return is.items
}

// ByLength 按长度排序字符串。
func ByLength(items []string, asc bool) []string {
	result := make([]string, len(items))
	copy(result, items)
	sort.Slice(result, func(i, j int) bool {
		if asc {
			return len(result[i]) < len(result[j])
		}
		return len(result[i]) > len(result[j])
	})
	return result
}

// ByField 按字段排序结构体（泛型支持）。
type ByField[T any] struct {
	items []T
	less  func(i, j int) bool
}

// NewByField 创建字段排序器。
func NewByField[T any](items []T, less func(i, j int) bool) *ByField[T] {
	copied := make([]T, len(items))
	copy(copied, items)
	return &ByField[T]{items: copied, less: less}
}

// Sort 执行排序。
func (bf *ByField[T]) Sort() []T {
	sort.Slice(bf.items, bf.less)
	return bf.items
}

// StableSort 稳定排序。
func (bf *ByField[T]) StableSort() []T {
	sort.SliceStable(bf.items, bf.less)
	return bf.items
}

// Reverse 反转切片。
func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	copy(result, items)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// Unique 去重（要求已排序）。
func Unique(items []string) []string {
	if len(items) == 0 {
		return []string{}
	}
	result := make([]string, 0, len(items))
	prev := ""
	for _, s := range items {
		if s != prev {
			result = append(result, s)
			prev = s
		}
	}
	return result
}

// UniqueSorted 排序并去重。
func UniqueSorted(items []string) []string {
	sorted := make([]string, len(items))
	copy(sorted, items)
	sort.Strings(sorted)
	return Unique(sorted)
}

// TopN 取前 N 个元素。
func TopN[T any](items []T, n int, less func(a, b T) bool) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(items) {
		result := make([]T, len(items))
		copy(result, items)
		return result
	}
	copied := make([]T, len(items))
	copy(copied, items)
	sort.Slice(copied, func(i, j int) bool {
		return less(copied[i], copied[j])
	})
	return copied[:n]
}

// GroupBy 按分组函数分组。
func GroupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range items {
		key := keyFn(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Filter 过滤切片。
func Filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 映射切片。
func Map[T any, R any](items []T, mapper func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

// Any 检查是否有任意元素满足条件。
func Any[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All 检查是否所有元素满足条件。
func All[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// IndexOf 查找元素索引。
func IndexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

// Contains 检查是否包含元素。
func Contains[T comparable](items []T, target T) bool {
	return IndexOf(items, target) >= 0
}

// CaseInsensitiveSort 不区分大小写排序。
func CaseInsensitiveSort(items []string) []string {
	result := make([]string, len(items))
	copy(result, items)
	sort.Slice(result, func(i, j int) bool {
		return strings.ToLower(result[i]) < strings.ToLower(result[j])
	})
	return result
}
