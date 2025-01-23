package main

import (
	"fmt"
	"push-swap/hiddenorder"
	"push-swap/sorttob"
	"push-swap/stacks"
	"push-swap/utils"
)

// validate returns false upon finding a duplicate on a slice of ints
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

// cleanInsts removes unnecessary instructions
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

// produceInstructions reads the input, runs two different sorting algorithms on it to get
// two separate sets of instructions as slices and chooses the shorter slice for output
func produceInstructions(argument string) ([]string, error) {
	var err error
	stacks.StackA, err = utils.ToNums(argument)
	if err != nil {
		return nil, fmt.Errorf("Error")
	}
	if !validate(stacks.StackA) {
		return nil, fmt.Errorf("Error")
	}
	stacks.ASorted = utils.BubSort(stacks.StackA, utils.IsGreater) // For checking stacks.StackA got sorted

	// save the stacks for the other algorithm
	origStackA := make([]int, len(stacks.StackA))
	origStackB := make([]int, len(stacks.StackB))
	copy(origStackA, stacks.StackA)
	copy(origStackB, stacks.StackB)

	instructions1 := []string{}
	instructions1 = sorttob.SortToBMethod(instructions1)
	instructions1 = cleanInsts(instructions1)

	// restore the original stacks
	stacks.StackA = make([]int, len(origStackA))
	stacks.StackB = make([]int, len(origStackB))
	copy(stacks.StackA, origStackA)
	copy(stacks.StackB, origStackB)

	instructions2 := []string{}
	instructions2 = hiddenorder.HiddenOrder(instructions2)
	instructions2 = cleanInsts(instructions2)

	// choose the shorter instructions
	var instructions []string
	if len(instructions1) < len(instructions2) {
		instructions = instructions1
	} else {
		instructions = instructions2
	}

	/* 	stacks.StackA = make([]int, len(origstacks.StackA)) // restore stacks and run commands for testing inside this program
	   	stacks.StackB = make([]int, len(origstacks.StackB))
	   	copy(stacks.StackA, origstacks.StackA)
	   	copy(stacks.StackB, origstacks.StackB)
	   	runComms(instructions) */

	return instructions, nil
}
