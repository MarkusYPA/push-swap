package sorttob

import (
	"push-swap/stacks"
	"push-swap/utils"
)

// shortestToBLast finds the shortest way to place the largest value from StackA to StackB and does it
func shortestToBLast(a []int) []string {
	nextToB := len(stacks.BSorted) - 1 - len(stacks.StackB)
	Bnxt0 := stacks.BSorted[nextToB] // Next element to push to stacks.StackB

	var Bnxt1 int
	var isB1 bool
	if nextToB > 0 { // Get the next in line if there is one
		isB1 = true
		Bnxt1 = stacks.BSorted[nextToB-1]
	}

	// get positive and negative distances to next up on B
	posDis, negDis := utils.Distances(a, Bnxt0)

	comms := []string{}

	if posDis == 0 {
		comms = append(comms, "pb")
		utils.RunComms(comms)
		return comms
	}

	if posDis == 1 {
		comms = append(comms, "sa", "pb")
		utils.RunComms(comms)
		return comms
	}

	if negDis == 1 {
		comms = append(comms, "rra", "pb")
		utils.RunComms(comms)
		return comms
	}

	swapAtEnd := false
	if posDis <= negDis {
		for i := 0; a[i] != Bnxt0; i++ {
			if isB1 && !swapAtEnd && a[i] == Bnxt1 && len(stacks.StackA) > 6 {
				comms = append(comms, "pb") // Push second-to-next to StackB and remember to swap
				swapAtEnd = true
				continue
			}
			comms = append(comms, "ra")
		}
		comms = append(comms, "pb")
		if swapAtEnd {
			comms = append(comms, "sb")
			swapAtEnd = false
		}
	} else {
		for i := 0; a[i] != Bnxt0; i-- {
			if isB1 && !swapAtEnd && a[i] == Bnxt1 && len(stacks.StackA) > 8 && utils.IsOnList(Bnxt1, stacks.BSorted) {
				comms = append(comms, "pb") // Push second-to-next to StackB and remember to swap
				swapAtEnd = true
			}
			comms = append(comms, "rra")
			if i == 0 {
				i = len(a)
			}
		}
		comms = append(comms, "pb")
		if swapAtEnd {
			comms = append(comms, "sb")
			swapAtEnd = false
		}
	}

	utils.RunComms(comms)
	return comms
}

// sort toBMethod first pushes n-2 elements from StackA to StackB
// so they are there in reverse order, and then pushes them back
func SortToBMethod() []string {
	ins := []string{}

	stacks.BSorted = utils.BubSort(stacks.StackA, utils.IsSmaller)
	if len(stacks.BSorted) > 2 {
		stacks.BSorted = stacks.BSorted[2:]
	}

	for len(stacks.StackA) > 2 {
		ins = append(ins, shortestToBLast(stacks.StackA)...)
	}

	if len(stacks.StackA) == 2 && stacks.StackA[0] > stacks.StackA[1] {
		utils.RunComms([]string{"sa"})
		ins = append(ins, "sa")
	}

	for len(stacks.StackB) > 0 {
		utils.RunComms([]string{"pa"})
		ins = append(ins, "pa")
	}

	return ins
}
