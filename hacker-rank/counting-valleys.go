package main

import "fmt"

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path

 https://www.hackerrank.com/challenges/counting-valleys/
*/

func countingValleys(steps int32, path string) int32 {
	// Variables to track altitude and valleys
	var altitude, valleys int32

	for _, step := range path {
		if step == 'U' {
			altitude++
			if altitude == 0 {
				valleys++
			}
		} else if step == 'D' {
			altitude--
		}
	}

	return valleys
}

func main() {
	// Correct path string with 'U' and 'D'
	result := countingValleys(8, "UDDDUDUU")
	result2 := countingValleys(12, "DDUUDDUDUUUD")
	fmt.Println(result)  // Expected output: 1
	fmt.Println(result2) // Expected output: 1
}
