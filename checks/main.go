package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	stackA  []int
	stackB  []int
	aSorted []int
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

	scnr := bufio.NewScanner(os.Stdin)
	for scnr.Scan() {
		txt := scnr.Text()

		if txt == "exit" || txt == "quit" || txt == "" {
			break
		}

		switch txt {
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

	if reflect.DeepEqual(stackA, aSorted) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
