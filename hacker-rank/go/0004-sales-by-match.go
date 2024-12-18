// https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true

package main

import (
	"sort"
)

func sockMerchant(n int32, ar []int32) int32 {
	var pairs int
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	i := 0
	for i < int(n)-1 {
		if ar[i] == ar[i+1] {
			pairs++
			i += 2 // Move to the next pair
		} else {
			i++ // Move to the next element
		}
	}

	return int32(pairs)
}

// func main() {
// 	n := int32(7)
// 	ar := []int32{1, 2, 1, 2, 1, 3, 2}
// 	result := sockMerchant(n, ar)
// 	fmt.Println(result)
// }
