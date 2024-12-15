package generics

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func getMaxOfType[T constraints.Ordered](s []T) T {
	var max T
	if len(s) == 0 {
		return max
	}
	max = s[0]
	for i := 1; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	return max
}

func filter[T any](s []T, predicate func (T) bool) []T{
	result := []T{}
	for _, d := range s {
		if predicate(d) {
			result = append(result, d)
		}
	}
	return result
}

func contains[T comparable](s []T, v T) bool {
	for _, d := range s {
		if v == d {
			return true
		}
 	}
	return false
}
func sorting[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}