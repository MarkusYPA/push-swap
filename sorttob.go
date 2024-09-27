package main

func shortestToBLast(a, bs []int) []string {
	nextToB := len(bs) - 1 - len(stackB)
	Bnxt0 := bs[nextToB] // Next element to push to stackB

	/* 	add := 1
	   	for isOnList(Bnxt0, stackB) {
	   		fmt.Println("Loopin here", Bnxt0, stackB)
	   		nextToB = len(bs) - 1 - len(stackB) - add
	   		Bnxt0 = bs[nextToB] // Next element to push to stackB
	   		add++
	   	} */

	var Bnxt1 int
	var isB1 bool
	if nextToB > 0 {
		isB1 = true
		Bnxt1 = bs[nextToB-1]
	}
	/* 	var Bnxt2 int
	   	var isB2 bool
	   	if nextToB > 1 {
	   		isB2 = true
	   		Bnxt2 = bs[nextToB-2]
	   	} */

	/* 	// measure positive distance to next element
	   	posDis := 0
	   	for posDis < len(a) {
	   		if a[posDis] == Bnxt0 {
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

	   		if a[i] == Bnxt0 {
	   			break
	   		}
	   		negDis++
	   	} */

	// get positive and negative distances to next up on B
	posDis, negDis := distances(a, Bnxt0)

	//fmt.Println(posDis, negDis, BLst)
	comms := []string{}

	if posDis == 0 {
		comms = append(comms, "pb")
		runComms(comms)
		return comms
	}

	if posDis == 1 {
		/* 		if isB1 && a[0] == Bnxt1 {
			comms = append(comms, "pb", "sa", "pb", "sb")
			runComms(comms)
			return comms
		} */
		comms = append(comms, "sa", "pb")
		runComms(comms)
		return comms
	}

	if negDis == 1 {
		/* 		if isB1 && a[0] == Bnxt1 {
			comms = append(comms, "pb", "rra", "pb", "sb")
			runComms(comms)
			return comms
		} */
		comms = append(comms, "rra", "pb")
		runComms(comms)
		return comms
	}

	swapAtEnd := false
	//rotAtEnd := false
	if posDis <= negDis {
		for i := 0; a[i] != Bnxt0; i++ {
			if isB1 && !swapAtEnd && a[i] == Bnxt1 && len(stackA) > 6 { //&& isOnList(Bnxt1, bSorted) && isOnList(Bnxt0, bSorted) {
				comms = append(comms, "pb") // Push second-to-next to stackB and remember to swap
				swapAtEnd = true
				i++
			}

			/* 			if isB2 && !rotAtEnd && a[i] == Bnxt2 {
				comms = append(comms, "pb", "rrb") // Push third-to-next to stackB, rotate it away and remember to rotate back
				rotAtEnd = true
				fmt.Println("Now this!", Bnxt0, Bnxt2, "stackA len:", len(stackA), i)
				i++
				fmt.Println(stackB)
			} */
			//fmt.Print(a[i], " ")
			comms = append(comms, "ra")
		}
		comms = append(comms, "pb")
		if swapAtEnd {
			comms = append(comms, "sb")
			swapAtEnd = false
		}
		/* 		if rotAtEnd {
			comms = append(comms, "rb")
			rotAtEnd = false
		} */
	} else {
		for i := 0; a[i] != Bnxt0; i-- {

			if isB1 && !swapAtEnd && a[i] == Bnxt1 && len(stackA) > 8 && isOnList(Bnxt1, bSorted) {
				comms = append(comms, "pb") // Push second-to-next to stackB and remember to swap
				swapAtEnd = true

			}

			/* 			if isB2 && !rotAtEnd && a[i] == Bnxt2 {
				comms = append(comms, "pb", "rrb") // Push third-to-next to stackB, rotate it away and remember to rotate back
				rotAtEnd = true
				i++
			} */

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
		/* 		if rotAtEnd {
			comms = append(comms, "rb")
			rotAtEnd = false
		} */
	}

	runComms(comms)
	//fmt.Println(comms, stackB)
	return comms
}

// sort toBMethod first pushes n-2 elements from stackA to stackB
// so they are there in reverse order, and then gets them back
func sortToBMethod(ins []string) []string {
	bSorted = bubSort(stackA, isSmaller)
	if len(bSorted) > 2 {
		bSorted = bSorted[2:]
	}

	//fmt.Println("A sorted:", aSorted)

	for len(stackA) > 2 {
		ins = append(ins, shortestToBLast(stackA, bSorted)...)
	}

	//fmt.Println("stackA:", stackA)

	if len(stackA) == 2 && stackA[0] > stackA[1] {
		runComms([]string{"sa"})
		ins = append(ins, "sa")
	}

	//fmt.Println("B before returns:", stackB)

	for len(stackB) > 0 {
		runComms([]string{"pa"})
		ins = append(ins, "pa")
	}

	return ins
}
