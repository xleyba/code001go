/**
 * Given a list of numbers and a number k, return whether any two numbers
 * from the list add up to k.
 *
 * For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
 *
 * Bonus: Can you do this in one pass?
 */
package main

import (
	"fmt"
	"sort"
	//"math"
	//"sort"
)

func main() {
	arr := []int{1, 10, 22, 5, 8}
	sort.Ints(arr)

	i := 0
	l := len(arr) - 1

	for ok := true; ok; ok = i < l-1 {

		fmt.Printf("Comparando Array[%d]=%d con Array[%d]=%d", i, arr[i], l, arr[l])

		if arr[i]+arr[l] == 15 {
			fmt.Print("...Es OK\n")
			//break
		} else if arr[i]+arr[l] < 15 {
			i++
			fmt.Print("...No OK\n")
		} else { // A[i] + A[j] > sum
			l--
			fmt.Print("...No OK\n")
			break
		}
	}

	fmt.Printf("FIN")
}
