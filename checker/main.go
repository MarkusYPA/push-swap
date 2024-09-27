package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	if len(args) != 1 {
		panic("One argument please")
	}

	stackA := toNums(args[0])
	stackB := []int{}

	fmt.Println(stackA, stackB)

	scnr := bufio.NewScanner(os.Stdin)
	for scnr.Scan() {
		txt := scnr.Text()

		if txt == "exit" || txt == "quit" {
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

		fmt.Println(stackA, stackB)
	}
}
