package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	stackA  []int
	stackB  []int
	aSorted []int
	bSorted []int
)

// toNums converts a string on numbers separated by spaces to a slice of ints
func toNums(in string) []int {
	nums := []int{}
	for _, numSt := range strings.Split(in, " ") {
		numIn, err := strconv.Atoi(numSt)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		nums = append(nums, numIn)
	}
	return nums
}

// bubSort is a bubble sort function that returns a slice of ints arranged according to the function f
func bubSort(s []int, f func(a, b int) bool) []int {
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

func isGreater(a, b int) bool {
	return a > b
}

func isSmaller(a, b int) bool {
	return a < b
}

// runComm is a shortcut to running one command
func runComm(s string) {
	runComms([]string{s})
}

func runComms(strs []string) {
	for _, s := range strs {

		switch s {
		case "pa":
			stackB, stackA = push(stackB, stackA)
		case "pb":
			stackA, stackB = push(stackA, stackB)
		case "sa":
			stackA = swap(stackA)
		case "sb":
			stackB = swap(stackB)
		case "ss":
			stackA = swap(stackA)
			stackB = swap(stackB)
		case "ra":
			stackA = rotate(stackA)
		case "rb":
			stackB = rotate(stackB)
		case "rr":
			stackA = rotate(stackA)
			stackB = rotate(stackB)
		case "rra":
			stackA = revRotate(stackA)
		case "rrb":
			stackB = revRotate(stackB)
		case "rrr":
			stackA = revRotate(stackA)
			stackB = revRotate(stackB)
		default:
			continue
		}
	}
}

func push(s1, s2 []int) ([]int, []int) {
	out := []int{}
	if len(s1) > 0 {
		out = append(out, s1[0])
		for i := 0; i < len(s2); i++ {
			out = append(out, s2[i])
		}
		s1 = s1[1:]
	} else {
		return s1, s2
	}
	return s1, out
}

func swap(s []int) []int {
	if len(s) > 1 {
		s[0], s[1] = s[1], s[0]
	}
	return s
}

func rotate(s []int) []int {
	if len(s) > 1 {
		return append(s[1:], s[0])
	} else {
		return s
	}
}

func revRotate(s []int) []int {
	if len(s) > 1 {
		return append(s[len(s)-1:], s[0:len(s)-1]...)
	} else {
		return s
	}
}

func cleanInsts(ins []string) []string {
	found := true
	for found {
		found = false
		// remove all pairs of pb-pa
		for i := 0; i < len(ins)-1; i++ {
			if ins[i] == "pb" && ins[i+1] == "pa" {
				if i == len(ins)-2 {
					ins = ins[:i]
				} else {
					ins = append(ins[:i], ins[i+2:]...)
				}
				found = true
			}
		}

	}
	return ins
}

func distances(a []int, num int) (int, int) {
	// measure positive distance to next element
	posDis := 0
	for posDis < len(a) {
		if a[posDis] == num {
			break
		}
		posDis++
	}

	// measure negative distance
	negDis := 0
	for i := 0; negDis < len(a); i-- {
		if i < 0 {
			i = len(a) - 1
		}

		if a[i] == num {
			break
		}
		negDis++
	}

	//fmt.Println("Stack and num:", a, num, "Distances:", posDis, negDis)
	return posDis, negDis
}

// isOnList checks if an int is on a slice of ints
func isOnList(n int, l []int) bool {
	for _, num := range l {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		panic("One argument please")
	}

	stackA = toNums(args[0])
	aSorted = bubSort(stackA, isGreater)
	instructions := []string{}

	//instructions = bigsToB(instructions)

	if len(stackA) < 35 {
		// Method where the elements (-2 smallest) are first transferred
		// to stackB to reverse order
		instructions = sortToBMethod(instructions)
	} else {
		// Method where each next element (and any incidental bigger-half element)
		// is first moved to B then back to A at the correct position
		instructions = sortAtAMethod(instructions)

		// Method where the biggre half is moved to B at the beginning
		//instructions = bigsToB(instructions)
	}

	instructions = cleanInsts(instructions)

	fmt.Println(instructions)
	fmt.Println("A:", stackA)
	fmt.Println(len(instructions), "instructions")

	if reflect.DeepEqual(stackA, aSorted) {
		fmt.Println("Stack A is sorted")
	} else {
		fmt.Println("Stack A is NOT sorted")
	}
}
