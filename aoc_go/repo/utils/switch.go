package utils

import "sort"

// Reverse
//  @param s
func Reverse(s interface{}) {
	sort.SliceStable(s, func(i, j int) bool {
		return true
	})
}
