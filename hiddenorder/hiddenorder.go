package hiddenorder

import (
	"fmt"
	"push-swap/stacks"
	"push-swap/utils"
	"sync"
)

var (
	ordMutex sync.Mutex
)

// hiddenOrder leaves an already sorted sequence of numbers on StackA, moves the rest to
// StackB, and then moves them back one-by-one into suitable gaps
func HiddenOrder(ins []string) []string {

	allOrders := [][]int{}

	wg := sync.WaitGroup{}
	for i := range stacks.StackA {
		wg.Add(1)
		getAllOrders(i, i, []int{}, &allOrders, &wg, true)
	}
	wg.Wait()

	bestOs := bestOrders(&allOrders)

	origStackA := make([]int, len(stacks.StackA))
	origStackB := make([]int, len(stacks.StackB))
	copy(origStackA, stacks.StackA)
	copy(origStackB, stacks.StackB)

	// for testing
	//bestOs := [][]int{{len(stacks.StackA) - 2, len(stacks.StackA) - 1}}

	bestNewInsts := []string{}
	for i, best := range bestOs {

		stacks.StackA = make([]int, len(origStackA))
		stacks.StackB = make([]int, len(origStackB))
		copy(stacks.StackA, origStackA)
		copy(stacks.StackB, origStackB)

		bestValues := []int{}
		for _, ind := range best {
			bestValues = append(bestValues, stacks.StackA[ind])
		}
		newInsts := pushToB(bestValues)

		for len(stacks.StackB) > 0 {

			// Put elements from stacks.StackB into suitable gaps on stacks.StackA
			for {
				// Find nearest gap
				p1, p2, err := nearestGap()
				if err != nil {
					break
				}

				// Move p1 to the top of A, p2 to the top of B and push from B to A
				newInsts = append(newInsts, toTop(p1, stacks.StackA, "A")...)
				utils.RunComms(toTop(p1, stacks.StackA, "A"))
				newInsts = append(newInsts, toTop(p2, stacks.StackB, "B")...)
				utils.RunComms(toTop(p2, stacks.StackB, "B"))
				newInsts = append(newInsts, "pa")
				utils.RunComm("pa")
			}
		}

		newInsts = append(newInsts, toTop(smallestOnList(stacks.StackA), stacks.StackA, "A")...)
		utils.RunComms(toTop(smallestOnList(stacks.StackA), stacks.StackA, "A"))

		if i == 0 || len(newInsts) < len(bestNewInsts) {
			bestNewInsts = newInsts
		}
	}

	stacks.StackA = make([]int, len(origStackA))
	stacks.StackB = make([]int, len(origStackB))
	copy(stacks.StackA, origStackA)
	copy(stacks.StackB, origStackB)

	utils.RunComms(bestNewInsts)
	ins = append(ins, bestNewInsts...)

	return ins
}

// rotStack runs a command cmd a given number d of times and returns the commands
func rotStack(d int, cmd string) []string {
	comms := []string{}
	for i := 0; i < d; i++ {
		comms = append(comms, cmd)
	}
	utils.RunComms(comms)
	return comms
}

// toTop moves the number n to the top of stack s, specified by the string l
func toTop(n int, s []int, l string) []string {
	comms := []string{}
	posD, negD := utils.Distances(s, n)
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

// midGapOnA checks if the element elemB from StackB fits in the gap between index indA and index indA+1 on StackA
func midGapOnA(indA, elemB int) bool {
	if indA > 0 && stacks.StackA[indA-1] <= elemB && stacks.StackA[indA] >= elemB {
		return true
	}
	if indA == 0 {
		aLast := stacks.StackA[len(stacks.StackA)-1]
		return aLast <= elemB && stacks.StackA[0] >= elemB
	}
	return false
}

// endGapOnA assumes stackA is sorted and randomly rotated.
// It returns true if the element at indA is the smallest one and elemB would be the biggest entry
// and therefore should be placed right in front of indA.
func endGapOnA(indA, elemB int) bool {
	if indA > 0 {
		aPrev := stacks.StackA[indA-1]
		aThis := stacks.StackA[indA]
		if aPrev > aThis && (elemB > aPrev || elemB < aThis) {
			return true
		}
	}
	if indA == 0 {
		aLast := stacks.StackA[len(stacks.StackA)-1]
		return stacks.StackA[0] < aLast && (elemB < stacks.StackA[0] || elemB > aLast)
	}
	return false
}

// min returns the smallest of the two given values
func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// nearestPair finds all the suitable gaps on stackA for each element on StackB, and chooses the move
// that can be executed with the fewest moves
func nearestGap() (int, int, error) {
	p1, p2 := 0, 0
	foundOne := false
	distOld := -1

	for i := 0; i < len(stacks.StackA); i++ {
		for j := 0; j < len(stacks.StackB); j++ {
			if midGapOnA(i, stacks.StackB[j]) || endGapOnA(i, stacks.StackB[j]) {
				if !foundOne {
					p1, p2 = stacks.StackA[i], stacks.StackB[j]
					distOld = min(utils.Distances(stacks.StackA, p1)) + min(utils.Distances(stacks.StackB, p2))
					foundOne = true
				} else {
					distNow := min(utils.Distances(stacks.StackA, stacks.StackA[i])) + min(utils.Distances(stacks.StackB, stacks.StackB[j]))

					if distNow < distOld { // New pair is closer than the old one
						p1, p2 = stacks.StackA[i], stacks.StackB[j]
						distOld = distNow
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

// pushToB moves values NOT on slice nums to stacks.StackB
func pushToB(nums []int) []string {
	ins := []string{}
	for _, n := range stacks.StackA {
		if !utils.IsOnList(n, nums) {
			ins = append(ins, "pb")
			utils.RunComm("pb")
		} else {
			ins = append(ins, "ra")
			utils.RunComm("ra")
		}
	}
	return ins
}

// bestOrder finds the longest sequences of ordered numbers
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

	return bests
}

// getAllOrders finds all sequences of increasing numbers through stacks.StackA starting from
// the element at index start
func getAllOrders(start int, index int, curSolution []int, orders *[][]int, wg *sync.WaitGroup, first bool) {

	// 2^20 should be enough to find one good enough to pass the audit
	if len(*orders) >= 1048576 {
		if first {
			wg.Done()
		}
		return
	}

	curSolution = append(curSolution, index)

	biggers := []int{} // Indices of all remaining values bigger than this, each used as the next step
	if index < start {
		if start == len(stacks.StackA)-1 {
			for i, n := range stacks.StackA[index:] {
				if n > stacks.StackA[index] {
					biggers = append(biggers, i+index)
				}
			}
		} else {
			for i, n := range stacks.StackA[index : start+1] {
				if n > stacks.StackA[index] {
					biggers = append(biggers, i+index)
				}
			}
		}
	} else {
		for i, n := range stacks.StackA[index:] {
			if n > stacks.StackA[index] {
				biggers = append(biggers, i+index)
			}
		}
		for i, n := range stacks.StackA[:start+1] {
			if n > stacks.StackA[index] {
				biggers = append(biggers, i)
			}
		}
	}

	// If no more biggers found, this solution is complete
	if len(biggers) == 0 {
		// copy values to a new slice to avoid pointer problems
		toSave := make([]int, len(curSolution))
		copy(toSave, curSolution)
		ordMutex.Lock()
		*orders = append(*orders, toSave)
		ordMutex.Unlock()

		if first {
			wg.Done()
		}
		return
	}

	//get all hidden orders for all the bigger numbers
	for _, n := range biggers {
		getAllOrders(start, n, curSolution, orders, wg, false)
	}

	if first {
		wg.Done()
	}

}
