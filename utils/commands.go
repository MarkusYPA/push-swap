package utils

import "push-swap/stacks"

// RunComm is a shortcut to running one command
func RunComm(s string) {
	RunComms([]string{s})
}

// RunComms runs the commands given as a slice of strings
func RunComms(strs []string) {
	for _, s := range strs {

		switch s {
		case "pa":
			stacks.StackB, stacks.StackA = push(stacks.StackB, stacks.StackA)
		case "pb":
			stacks.StackA, stacks.StackB = push(stacks.StackA, stacks.StackB)
		case "sa":
			stacks.StackA = swap(stacks.StackA)
		case "sb":
			stacks.StackB = swap(stacks.StackB)
		case "ss":
			stacks.StackA = swap(stacks.StackA)
			stacks.StackB = swap(stacks.StackB)
		case "ra":
			stacks.StackA = rotate(stacks.StackA)
		case "rb":
			stacks.StackB = rotate(stacks.StackB)
		case "rr":
			stacks.StackA = rotate(stacks.StackA)
			stacks.StackB = rotate(stacks.StackB)
		case "rra":
			stacks.StackA = revRotate(stacks.StackA)
		case "rrb":
			stacks.StackB = revRotate(stacks.StackB)
		case "rrr":
			stacks.StackA = revRotate(stacks.StackA)
			stacks.StackB = revRotate(stacks.StackB)
		default:
			continue
		}
	}
}

// push moves the first element of the slice s1 to the top of s2
func push(s1, s2 []int) ([]int, []int) {

	if len(s1) == 0 {
		return s1, s2
	}

	out := []int{s1[0]}
	out = append(out, s2...)

	return s1[1:], out
}

// swap changes the first two elements around
func swap(s []int) []int {
	if len(s) > 1 {
		s[0], s[1] = s[1], s[0]
	}
	return s
}

// rotate puts the first element last
func rotate(s []int) []int {
	if len(s) > 1 {
		return append(s[1:], s[0])
	} else {
		return s
	}
}

// revRotate puts the last element first
func revRotate(s []int) []int {
	if len(s) > 1 {
		return append(s[len(s)-1:], s[0:len(s)-1]...)
	} else {
		return s
	}
}
