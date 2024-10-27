package main

// shortestToBLast finds the shortest way to place the largest value from stackA to stackB and does it
func shortestToBLast(a, bs []int) []string {
	nextToB := len(bs) - 1 - len(stackB)
	Bnxt0 := bs[nextToB] // Next element to push to stackB

	var Bnxt1 int
	var isB1 bool
	if nextToB > 0 { // Get the next in line if there is one
		isB1 = true
		Bnxt1 = bs[nextToB-1]
	}

	// get positive and negative distances to next up on B
	posDis, negDis := distances(a, Bnxt0)

	comms := []string{}

	if posDis == 0 {
		comms = append(comms, "pb")
		runComms(comms)
		return comms
	}

	if posDis == 1 {
		comms = append(comms, "sa", "pb")
		runComms(comms)
		return comms
	}

	if negDis == 1 {
		comms = append(comms, "rra", "pb")
		runComms(comms)
		return comms
	}

	swapAtEnd := false
	if posDis <= negDis {
		for i := 0; a[i] != Bnxt0; i++ {
			if isB1 && !swapAtEnd && a[i] == Bnxt1 && len(stackA) > 6 {
				comms = append(comms, "pb") // Push second-to-next to stackB and remember to swap
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
			if isB1 && !swapAtEnd && a[i] == Bnxt1 && len(stackA) > 8 && isOnList(Bnxt1, bSorted) {
				comms = append(comms, "pb") // Push second-to-next to stackB and remember to swap
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

	runComms(comms)
	return comms
}

// sort toBMethod first pushes n-2 elements from stackA to stackB
// so they are there in reverse order, and then pushes them back
func sortToBMethod(ins []string) []string {
	bSorted = bubSort(stackA, isSmaller)
	if len(bSorted) > 2 {
		bSorted = bSorted[2:]
	}

	for len(stackA) > 2 {
		ins = append(ins, shortestToBLast(stackA, bSorted)...)
	}

	if len(stackA) == 2 && stackA[0] > stackA[1] {
		runComms([]string{"sa"})
		ins = append(ins, "sa")
	}

	for len(stackB) > 0 {
		runComms([]string{"pa"})
		ins = append(ins, "pa")
	}

	return ins
}
