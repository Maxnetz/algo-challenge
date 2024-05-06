package main

// https://www.hackerrank.com/challenges/drawing-book/problem?isFullScreen=true

func pageCount(n int32, p int32) int32 {
	// Calculate Page Turns from the Front:
	front := p / 2

	// Calculate Page Turns from the Back:
	back := (n - p) / 2

	// The adjustment if book pages are even
	if n%2 == 0 {
		back = (n + 1 - p) / 2
	}

	// Choosing the Minimum Turns:
	if front < back {
		return front
	} else {
		return back
	}
}

// func main() {
// 	n1 := int32(6)
// 	p1 := int32(2)
// 	result1 := pageCount(n1, p1)
// 	fmt.Println(result1)

// 	n2 := int32(5)
// 	p2 := int32(4)
// 	result2 := pageCount(n2, p2)
// 	fmt.Println(result2)
// }
