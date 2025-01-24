package switcheroo

import (
	"push-swap/stacks"
	"push-swap/utils"
	"reflect"
)

func OnlySwap() []string {
	insts := []string{}

	start := 0
	end := len(stacks.StackA) - 1
	index := 0
	rotate := "ra"
	changes := 0

	if len(stacks.StackA) > 2 {
		for {
			if stacks.StackA[0] > stacks.StackA[1] {
				utils.RunComm("sa")
				insts = append(insts, "sa")
				changes++
			}

			if rotate == "ra" && index >= end-1 {
				end--
				rotate = "rra"
				if changes == 0 {
					break
				}
				changes = 0
			} else if rotate == "rra" && index <= start {
				start++
				rotate = "ra"
				if changes == 0 {
					break
				}
				changes = 0
			}

			utils.RunComm(rotate)
			insts = append(insts, rotate)

			if rotate == "ra" {
				index++
			} else {
				index--
			}
		}
	} else if len(stacks.StackA) == 2 {
		if stacks.StackA[0] > stacks.StackA[1] {
			utils.RunComm("sa")
			insts = append(insts, "sa")
		}
		return insts
	} else {
		return insts
	}

	// undo and remove last rotations
	for i := len(insts) - 1; i > 0; i-- {
		if insts[i] == "ra" {
			utils.RunComm("rra")
			insts = insts[:len(insts)-1]
		} else if insts[i] == "rra" {
			utils.RunComm("ra")
			insts = insts[:len(insts)-1]
		} else {
			break
		}
	}

	smallInd := 0
	for i, num := range stacks.StackA {
		if num == stacks.ASorted[0] {
			smallInd = i
			break
		}
	}

	if smallInd > len(stacks.StackA)/2 {
		rotate = "rra"
	} else {
		rotate = "ra"
	}

	for !reflect.DeepEqual(stacks.StackA, stacks.ASorted) {
		utils.RunComm(rotate)
		insts = append(insts, rotate)
	}

	return insts
}
