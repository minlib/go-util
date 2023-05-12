package slicex

import (
	"github.com/minlib/go-util/jsonx"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"strings"
)

// Distinct2D remove duplicate two-dimensional arrays
func Distinct2D[E constraints.Ordered](s [][]E) [][]E {
	var res [][]E
	m := make(map[string]struct{}, len(s))
	for i := range s {
		key := jsonx.MarshalString(s[i])
		if _, found := m[key]; !found {
			m[key] = struct{}{}
			res = append(res, s[i])
		}
	}
	return res
}

// Equal2D reports whether two slices are equal: the same length and all
func Equal2D[E comparable](s1, s2 [][]E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !Equal(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

// EqualIgnoreOrder2D reports whether two slices are equal: the same length and all
func EqualIgnoreOrder2D[E constraints.Ordered](s1, s2 [][]E) bool {
	if len(s1) != len(s2) {
		return false
	}
	ss1 := make([]string, len(s1))
	ss2 := make([]string, len(s2))
	for i := range s1 {
		slices.Sort(s1[i])
		ss1[i] = jsonx.MarshalString(s1[i])
	}
	for i := range s2 {
		slices.Sort(s2[i])
		ss2[i] = jsonx.MarshalString(s2[i])
	}
	slices.Sort(ss1)
	slices.Sort(ss2)
	for i := range ss1 {
		if ss1[i] != ss2[i] {
			return false
		}
	}
	return true
}

// Subtract2D returns the elements in `s1` that aren't in `s2`.
func Subtract2D[E comparable](s1, s2 [][]E) [][]E {
	var res [][]E
	if len(s1) > 0 {
		m := make(map[string]struct{}, len(s2))
		for i := range s2 {
			key := jsonx.MarshalString(s2[i])
			m[key] = struct{}{}
		}
		for i := range s1 {
			key := jsonx.MarshalString(s1[i])
			if _, found := m[key]; !found {
				res = append(res, s1[i])
			}
		}
	}
	return res
}

// Combine2D returns the combination of s1 and s2 elements
func Combine2D[E comparable](arrays [][]E) [][]E {
	var total = 0
	var count = make([]int, len(arrays))
	if len(arrays) > 0 {
		total = 1
		for i := len(arrays) - 1; i >= 0; i-- {
			total *= len(arrays[i])
			count[i] = len(arrays[i])
		}
	}
	var result [][]E
	for i := 0; i < total; i++ {
		var p = make([]int, len(arrays))
		var val = i
		var item []E
		for row := len(arrays) - 1; row >= 0; row-- {
			col := val % count[row]
			p[row] = col
			val = val / count[row]
			item = append([]E{arrays[row][col]}, item...)
		}
		result = append(result, item)
	}
	return result
}

// Join2D concatenate the elements of an array with sep
func Join2D(s [][]string, sep string) []string {
	var result []string
	for _, v := range s {
		item := strings.Join(v, sep)
		result = append(result, item)
	}
	return result
}
