package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// ToNums converts a string on numbers separated by spaces to a slice of ints
func ToNums(in string) ([]int, error) {
	nums := []int{}
	for _, numSt := range strings.Split(in, " ") {
		numIn, err := strconv.Atoi(numSt)
		if err != nil {
			return nil, fmt.Errorf("Error")
		}
		nums = append(nums, numIn)
	}
	return nums, nil
}

// BubSort is a bubble sort function that returns a slice of ints arranged according to the function f
func BubSort(s []int, f func(a, b int) bool) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if f(sorted[i], sorted[j]) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}

func IsGreater(a, b int) bool {
	return a > b
}

func IsSmaller(a, b int) bool {
	return a < b
}

// Distances returns how many instructions are needed to rotate a stack a in
// the positive and the negative directions so that element num is on top
func Distances(a []int, num int) (int, int) {

	// measure positive distance to num
	posDis := 0
	for i := 0; i < len(a); i++ {
		if a[i] == num {
			posDis = i
			break
		}
	}

	// deduce negative distance
	negDis := 0
	if posDis != 0 {
		negDis = len(a) - posDis
	}

	return posDis, negDis
}

// IsOnList checks if an int is on a slice of ints
func IsOnList(n int, l []int) bool {
	for _, num := range l {
		if n == num {
			return true
		}
	}
	return false
}
