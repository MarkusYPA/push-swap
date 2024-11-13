package main

import (
	"fmt"
)

// hiddenOrder leaves an already sorted sequence of numbers on stackA, moves the rest to
// stackB and then moves them back one-by-one into suitable gaps
func hiddenOrder(ins []string) []string {

	allOrders := [][]int{}
	for i := range stackA {
		getAllOrders(i, i, []int{}, &allOrders)
	}
	bestOs := bestOrders(&allOrders)
	//fmt.Println("Length of all orders:", len(allOrders), "Bests found:", len(bestOs), "Length of Best:", len(bestOs[0]))

	origStackA := make([]int, len(stackA))
	origStackB := make([]int, len(stackB))
	copy(origStackA, stackA)
	copy(origStackB, stackB)

	bestNewInsts := []string{}
	for i, best := range bestOs {

		stackA = make([]int, len(origStackA))
		stackB = make([]int, len(origStackB))
		copy(stackA, origStackA)
		copy(stackB, origStackB)

		bestValues := []int{}
		for _, ind := range best {
			bestValues = append(bestValues, stackA[ind])
		}
		newInsts := pushToB(bestValues)

		for len(stackB) > 0 {

			// Put elements from stackB into suitable gaps on stackA
			for {
				// Find nearest gap
				p1, p2, err := nearestGap()
				if err != nil {
					break
				}

				// Move p1 to the top of A, p2 to the top of B and push from B to A
				newInsts = append(newInsts, toTop(p1, stackA, "A")...)
				runComms(toTop(p1, stackA, "A"))
				newInsts = append(newInsts, toTop(p2, stackB, "B")...)
				runComms(toTop(p2, stackB, "B"))
				newInsts = append(newInsts, "pa")
				runComm("pa")
			}
		}

		newInsts = append(newInsts, toTop(smallestOnList(stackA), stackA, "A")...)
		runComms(toTop(smallestOnList(stackA), stackA, "A"))

		if i == 0 || len(newInsts) < len(bestNewInsts) {
			bestNewInsts = newInsts
		}
	}

	stackA = make([]int, len(origStackA))
	stackB = make([]int, len(origStackB))
	copy(stackA, origStackA)
	copy(stackB, origStackB)

	runComms(bestNewInsts)
	ins = append(ins, bestNewInsts...)

	return ins
}

// rotStack runs a command cmd a given number d of times and returns the commands
func rotStack(d int, cmd string) []string {
	comms := []string{}
	for i := 0; i < d; i++ {
		comms = append(comms, cmd)
	}
	runComms(comms)
	return comms
}

// toTop moves the number n to the top of stack s, specified by the string l
func toTop(n int, s []int, l string) []string {
	comms := []string{}
	posD, negD := distances(s, n)
	calls := []string{"", ""}
	if l == "A" {
		calls[0], calls[1] = "ra", "rra"
	} else {
		calls[0], calls[1] = "rb", "rrb"
	}

	if float64(posD) <= float64(negD) {
		comms = append(comms, rotStack(posD, calls[0])...)
	} else {
		comms = append(comms, rotStack(negD, calls[1])...)
	}
	return comms
}

// smallestOnList returns the smallest element on a slice of integers
func smallestOnList(s []int) int {
	small := s[0]
	for _, n := range s {
		if n < small {
			small = n
		}
	}
	return small
}

// goodGapOnA checks if the element elemB from stackB fits in the gap between index indA and index indA+1 on stackA
func midGapOnA(indA, elemB int) bool {
	for i := 1; i < len(stackA); i++ {
		if indA > 0 && stackA[indA-1] <= elemB && stackA[indA] >= elemB {
			return true
		}
	}

	if indA == 0 {
		aLast := stackA[len(stackA)-1]
		return aLast <= elemB && stackA[0] >= elemB
	}
	return false
}

func endGapOnA(indA, elemB int) bool {
	for i := 1; i < len(stackA); i++ {
		if indA > 0 {
			aPrev := stackA[indA-1]
			aThis := stackA[indA]
			if aPrev > aThis && (elemB > aPrev || elemB < aThis) {
				return true
			}
		}
	}
	if indA == 0 {
		aLast := stackA[len(stackA)-1]
		return stackA[0] < aLast && (elemB < stackA[0] || elemB > aLast)
	}
	return false
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func nearestEndGap() (int, int, error) {
	p1, p2 := 0, 0
	foundOne := false
	distOld := -1

	for i := 0; i < len(stackA); i++ {
		for j := 0; j < len(stackB); j++ {
			if endGapOnA(i, stackB[j]) {
				if !foundOne {
					p1, p2 = stackA[i], stackB[j]
					distOld = min(distances(stackA, p1)) + min(distances(stackB, p2))
					foundOne = true
				} else {
					distNow := min(distances(stackA, stackA[i])) + min(distances(stackB, stackB[j]))

					if distNow < distOld { // New pair is closer than the old one
						p1, p2 = stackA[i], stackB[j]
					}
				}
			}
		}
	}

	if foundOne {
		return p1, p2, nil
	} else {
		return p1, p2, fmt.Errorf("no gap found")
	}
}

// nearestPair finds the gap on stackA that fits an element from stackB that can be filled with
// the fewest moves
func nearestGap() (int, int, error) {
	p1, p2 := 0, 0
	foundOne := false
	distOld := -1

	for i := 0; i < len(stackA); i++ {
		for j := 0; j < len(stackB); j++ {
			if midGapOnA(i, stackB[j]) || endGapOnA(i, stackB[j]) {
				if !foundOne {
					p1, p2 = stackA[i], stackB[j]
					distOld = min(distances(stackA, p1)) + min(distances(stackB, p2))
					foundOne = true
				} else {
					distNow := min(distances(stackA, stackA[i])) + min(distances(stackB, stackB[j]))

					if distNow < distOld { // New pair is closer than the old one
						p1, p2 = stackA[i], stackB[j]
					}
				}
			}
		}
	}

	if foundOne {
		return p1, p2, nil
	} else {
		return p1, p2, fmt.Errorf("no gap found")
	}
}

// pushToB moves values NOT on slice nums to stackB
func pushToB(nums []int) []string {
	ins := []string{}
	for _, n := range stackA {
		if !isOnList(n, nums) {
			ins = append(ins, "pb")
			runComm("pb")
		} else {
			ins = append(ins, "ra")
			runComm("ra")
		}
	}
	return ins
}

// bestOrder finds the longest sequence of ordered numbers with the lowest starting index
func bestOrders(allOrders *[][]int) [][]int {
	longest := len((*allOrders)[0])
	for _, o := range *allOrders {
		if len(o) > longest {
			longest = len(o)
		}
	}
	bests := [][]int{}
	for _, o := range *allOrders {
		if len(o) == longest {
			bests = append(bests, o)
		}
	}
	for _, o := range *allOrders {
		if len(o) > 1 && len(o) == longest-1 && len(bests) <= 5000 {
			bests = append(bests, o)
		}
		/* if len(bests) == 5000 {
			fmt.Println("at -1")
			break
		} */
	}
	for _, o := range *allOrders {
		if len(o) > 1 && len(o) == longest-2 && len(bests) <= 5000 {
			bests = append(bests, o)
		}
		/* if len(bests) == 5000 {
			fmt.Println("at -2")
			break
		} */
	}
	return bests
}

// Store examined values on a map with their indices at that solution (later is better)
//var founds map[int]int = map[int]int{}

// getAllOrders finds all sequences of increasing numbers through stackA starting from
// the element at index start
func getAllOrders(start int, index int, curSolution []int, orders *[][]int) {
	curSolution = append(curSolution, index)

	// Find the indices of all remaining values bigger than this
	biggers := []int{}
	if index < start {
		if start == len(stackA)-1 {
			for i, n := range stackA[index:] {
				if n > stackA[index] {
					biggers = append(biggers, i+index)
				}
			}
		} else {
			for i, n := range stackA[index : start+1] {
				if n > stackA[index] {
					biggers = append(biggers, i+index)
				}
			}
		}
	} else {
		for i, n := range stackA[index:] {
			if n > stackA[index] {
				biggers = append(biggers, i+index)
			}
		}
		for i, n := range stackA[:start+1] {
			if n > stackA[index] {
				biggers = append(biggers, i)
			}
		}
	}

	// If no more biggers found, this solution is complete
	if len(biggers) == 0 {
		// copy values to a new slice to avoid pointer problems
		toSave := make([]int, len(curSolution))
		copy(toSave, curSolution)
		*orders = append(*orders, toSave)

		/* 		for i, n := range toSave {
			if _, ok := founds[n]; !ok || founds[n] < len(curSolution) {
				founds[n] = i
			}
		} */

		return
	}

	//get all hidden orders for all the bigger numbers
	for _, n := range biggers {

		// if not yet found or the found value is worse than what we have now
		//if _, ok := founds[n]; !ok || founds[n] <= len(curSolution) {
		getAllOrders(start, n, curSolution, orders)
		//}

	}
}
