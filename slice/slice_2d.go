package slice

import (
	"github.com/minlib/go-util/json"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Distinct2D 移除重复的二维数组
func Distinct2D[E constraints.Ordered](s [][]E) [][]E {
	var res [][]E
	m := make(map[string]struct{}, len(s))
	for i := range s {
		key := json.ToJsonString(s[i])
		if _, found := m[key]; !found {
			m[key] = struct{}{}
			res = append(res, s[i])
		}
	}
	return res
}

// Equal2D
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

// EqualIgnoreOrder2D
func EqualIgnoreOrder2D[E constraints.Ordered](s1, s2 [][]E) bool {
	if len(s1) != len(s2) {
		return false
	}
	ss1 := make([]string, len(s1))
	ss2 := make([]string, len(s2))
	for i := range s1 {
		slices.Sort(s1[i])
		ss1[i] = json.ToJsonString(s1[i])
	}
	for i := range s2 {
		slices.Sort(s2[i])
		ss2[i] = json.ToJsonString(s2[i])
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
			key := json.ToJsonString(s2[i])
			m[key] = struct{}{}
		}
		for i := range s1 {
			key := json.ToJsonString(s1[i])
			if _, found := m[key]; !found {
				res = append(res, s1[i])
			}
		}
	}
	return res
}
