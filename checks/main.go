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
	StackA  []int
	StackB  []int
	ASorted []int
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
	StackA, err = toNums(args[0])
	if err != nil {
		fmt.Println("Error")
		return
	}
	if !validate(StackA) {
		fmt.Println("Error")
		return
	}
	ASorted = bubSort(StackA, isGreater)

	scnr := bufio.NewScanner(os.Stdin)
	for scnr.Scan() {
		txt := scnr.Text()

		if txt == "exit" || txt == "quit" || txt == "" {
			break
		}

		switch txt {
		case "pa":
			StackB, StackA = push(StackB, StackA)
		case "pb":
			StackA, StackB = push(StackA, StackB)
		case "sa":
			StackA = swap(StackA)
		case "sb":
			StackB = swap(StackB)
		case "ss":
			StackA = swap(StackA)
			StackB = swap(StackB)
		case "ra":
			StackA = rotate(StackA)
		case "rb":
			StackB = rotate(StackB)
		case "rr":
			StackA = rotate(StackA)
			StackB = rotate(StackB)
		case "rra":
			StackA = revRotate(StackA)
		case "rrb":
			StackB = revRotate(StackB)
		case "rrr":
			StackA = revRotate(StackA)
			StackB = revRotate(StackB)
		default:
			continue
		}
	}

	if reflect.DeepEqual(StackA, ASorted) && len(StackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
