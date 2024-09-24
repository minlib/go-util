package slicex

import (
	"github.com/minlib/go-util/core"
	"strconv"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Slice return a slice type
func Slice[E any](s ...E) []E {
	return []E(s)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	return slices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[E any](s []E, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return slices.Contains(s, v)
}

// ContainsAny 数组中包含任意一个元素
func ContainsAny[S ~[]E, E comparable](s S, e ...E) bool {
	for _, v := range e {
		if slices.Contains(s, v) {
			return true
		}
	}
	return false
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[E comparable](s1, s2 []E) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Delete removes the elements s[i:i+1] from s, returning the modified slice
func Delete[S ~[]E, E any](s S, i int) S {
	return slices.Delete(s, i, i+1)
}

// Distinct 返回去重后的元素
func Distinct[S ~[]E, E comparable](s S) S {
	if len(s) == 0 {
		return s
	}
	var res S
	m := make(map[E]struct{})
	for _, v := range s {
		if _, found := m[v]; !found {
			m[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Duplicate 返回重复的元素
func Duplicate[S ~[]E, E comparable](s S) S {
	if len(s) == 0 {
		return s
	}
	var res S
	m := make(map[E]int)
	for _, v := range s {
		m[v]++
		if m[v] == 2 {
			res = append(res, v)
		}
	}
	return res
}

// Subtract returns the elements in `s1` that aren't in `s2`
func Subtract[S ~[]E, E comparable](s1, s2 S) S {
	var s S
	if len(s1) > 0 {
		var mp = make(map[E]struct{}, len(s2))
		for _, v := range s2 {
			mp[v] = struct{}{}
		}
		for _, v := range s1 {
			if _, found := mp[v]; !found {
				s = append(s, v)
			}
		}
	}
	return s
}

// SubtractDistinct 返回差集并去重
func SubtractDistinct[S ~[]E, E comparable](s1, s2 S) S {
	s := Subtract(s1, s2)
	return Distinct(s)
}

// Intersect 返回交集并去重
func Intersect[S ~[]E, E comparable](s1, s2 S) S {
	var s S
	if len(s1) > 0 && len(s2) > 0 {
		var mp = make(map[E]struct{}, 0)
		for _, v := range s2 {
			mp[v] = struct{}{}
		}
		for _, v := range s1 {
			if _, found := mp[v]; found {
				s = append(s, v)
			}
		}
	}
	return Distinct(s)
}

// Union 返回并集并去重
func Union[S ~[]E, E comparable](s1, s2 S) S {
	s := append(s1, s2...)
	return Distinct(s)
}

// Sum 求和
func Sum[S ~[]E, E constraints.Ordered](s S) E {
	var e E
	for _, v := range s {
		e += v
	}
	return e
}

// EqualIgnoreOrder 先排序再比较，不同排序的切片对比返回true
func EqualIgnoreOrder[E constraints.Ordered](s1, s2 []E) bool {
	slices.Sort(s1)
	slices.Sort(s2)
	return Equal(s1, s2)
}

// IntToString 整型切片转为字符串切片
func IntToString[S ~[]E, E constraints.Integer](s S) []string {
	var res []string
	for i := range s {
		res = append(res, strconv.FormatInt(int64(s[i]), 10))
	}
	return res
}

// StringToInt 字符串切片转为整型切片，E为转换后的整型类型
func StringToInt[E constraints.Integer](s []string) ([]E, error) {
	var res []E
	for i := range s {
		if val, err := strconv.Atoi(s[i]); err != nil {
			return nil, err
		} else {
			res = append(res, E(val))
		}
	}
	return res, nil
}

// LongToInt64 convert long slice to int64 slice, ignore nil value
func LongToInt64(s []core.Long) []int64 {
	var res []int64
	for _, v := range s {
		if v.Int64 != nil {
			res = append(res, *v.Int64)
		}
	}
	return res
}

// Int64ToLong convert int64 slice to long slice
func Int64ToLong(s []int64) []core.Long {
	var res []core.Long
	for _, v := range s {
		res = append(res, core.NewLong(v))
	}
	return res
}

// Extract 提取切片中某个属性的集合
func Extract[T any, U any](slice []T, extractor func(T) U) []U {
	var result []U
	for _, item := range slice {
		result = append(result, extractor(item))
	}
	return result
}
