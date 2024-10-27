package main

import (
	"fmt"
	"log"
	"os"
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
func toNums(in string) ([]int, error) {
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

	// remove all pairs of pb-pa
	for found {
		found = false
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

	// turn pairs of rra-rrb or rrb-rra to rrr
	found = true
	for found {
		found = false
		for i := 0; i < len(ins)-1; i++ {
			if (ins[i] == "rra" && ins[i+1] == "rrb") || (ins[i] == "rrb" && ins[i+1] == "rra") {
				ins[i] = "rrr"
				if i == len(ins)-2 {
					ins = ins[:i+1]
				} else {
					ins = append(ins[:i+1], ins[i+2:]...)
				}
				found = true
			}
		}
	}

	// turn pairs of ra-rb or rb-ra to rr
	found = true
	for found {
		found = false
		for i := 0; i < len(ins)-1; i++ {
			if (ins[i] == "ra" && ins[i+1] == "rb") || (ins[i] == "rb" && ins[i+1] == "ra") {
				ins[i] = "rr"
				if i == len(ins)-2 {
					ins = ins[:i+1]
				} else {
					ins = append(ins[:i+1], ins[i+2:]...)
				}
				found = true
			}
		}
	}

	// turn pairs of sa-sb or sb-sa to ss
	found = true
	for found {
		found = false
		for i := 0; i < len(ins)-1; i++ {
			if (ins[i] == "sa" && ins[i+1] == "sb") || (ins[i] == "sb" && ins[i+1] == "sa") {
				ins[i] = "ss"
				if i == len(ins)-2 {
					ins = ins[:i+1]
				} else {
					ins = append(ins[:i+1], ins[i+2:]...)
				}
				found = true
			}
		}
	}

	return ins
}

// distances returns how many instructions are needed to rotate stackA in
// the positive and the negative directions so that num is on top
func distances(a []int, num int) (int, int) {

	// measure positive distance to num
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

func last(s []int) int {
	return s[len(s)-1]
}

func validate(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		for j := i + 1; j < len(stack); j++ {
			if stack[i] == stack[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args[0]) == 0 {
		return
	}
	if len(args) > 1 {
		log.Fatalln("Use with one argument, e.g.: ./push-swap \"7 12 0 31 3\" ")
	}

	var err error
	stackA, err = toNums(args[0])
	if err != nil {
		fmt.Println("Error")
		return
	}
	if !validate(stackA) {
		fmt.Println("Error")
		return
	}
	aSorted = bubSort(stackA, isGreater)

	origStackA := make([]int, len(stackA))
	origStackB := make([]int, len(stackB))
	copy(origStackA, stackA)
	copy(origStackB, stackB)

	//instructions = bigsToB(instructions)

	/* 	if len(stackA) < 35 {
	   		// Method where the elements (-2 smallest) are first transferred
	   		// to stackB to reverse order
	   		instructions = sortToBMethod(instructions)
	   	} else {
	   		// Method where each next element (and any incidental bigger-half element)
	   		// is first moved to B then back to A at the correct position
	   		instructions = sortAtAMethod(instructions)

	   		// Method where the biggre half is moved to B at the beginning
	   		//instructions = bigsToB(instructions)
	   	} */

	instructions1 := []string{}
	instructions1 = sortToBMethod(instructions1)
	instructions1 = cleanInsts(instructions1)

	stackA = make([]int, len(origStackA))
	stackB = make([]int, len(origStackB))
	copy(stackA, origStackA)
	copy(stackB, origStackB)

	instructions2 := []string{}
	instructions2 = hiddenOrder(instructions2)
	instructions2 = cleanInsts(instructions2)

	stackA = make([]int, len(origStackA))
	stackB = make([]int, len(origStackB))
	copy(stackA, origStackA)
	copy(stackB, origStackB)

	var instructions []string
	if len(instructions1) < len(instructions2) {
		instructions = instructions1
	} else {
		instructions = instructions2
	}

	runComms(instructions)

	for _, ins := range instructions {
		fmt.Println(ins)
	}

	/*
		 	fmt.Println(len(instructions), "instructions")
			if reflect.DeepEqual(stackA, aSorted) {
				fmt.Println("Stack A is sorted")
			} else {
				fmt.Println("Stack A is NOT sorted")
			}
	*/
}
