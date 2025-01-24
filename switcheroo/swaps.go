package switcheroo

import (
	"push-swap/stacks"
	"push-swap/utils"
	"reflect"
)

// OnlySwap never uses StackB, instead if performs a zig-zaggy bubble sort on StackA,
// swapping and rotating one way, then swapping and reverse rotating the other, back and forth.
func OnlySwap() []string {
	insts := []string{}

	start := 0
	end := len(stacks.StackA) - 1
	index := 0
	rotate := "ra"
	changes := 0 // changes made going in one direction

	if len(stacks.StackA) > 2 {
		for {
			// Swap when necessary
			if stacks.StackA[0] > stacks.StackA[1] {
				utils.RunComm("sa")
				insts = append(insts, "sa")
				changes++
			}

			if rotate == "ra" && index >= end-1 { // Going forward, turn back at end
				end-- // big one has been moved to the back, no need to go there again
				rotate = "rra"
				if changes == 0 {
					break
				}
				changes = 0

			} else if rotate == "rra" && index <= start { // Reversing, go forward at start
				start++ // small one has been moved to the front, no need to go there again
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

	// Take shortest route to rotate smallest to top
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
