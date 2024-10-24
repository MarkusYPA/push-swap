package main

import (
	"fmt"
)

func secretOrder(ins []string) []string {

	allOrders := [][]int{}
	for i := range stackA {
		getAllOrders(i, i, []int{}, &allOrders)
	}

	fmt.Println("Length of all orders:", len(allOrders))
	bestO := bestOrder(&allOrders)
	fmt.Print("The best one is ", len(bestO), " long: ")
	for _, ind := range bestO {
		fmt.Printf("%v ", stackA[ind])
	}
	fmt.Println(bestO)
	fmt.Println()
	fmt.Println()

	bestValues := []int{}
	for _, ind := range bestO {
		bestValues = append(bestValues, stackA[ind])
	}
	ins = append(ins, pushToB(bestValues)...)

	/*




	 */

	return ins
}

// nearestPair finds the pair that can be joined with the fewest moves on stacks A and B
// Pairs are two values, one on A and the other on B, that are subsequent on a sorted
// list and when pushed with "pb", will be ordered correctly on B (big to small)
func nearestPairB() (int, int, string, error) {
	p1, p2 := 0, 0
	foundOne := false
	distOld := -1
	stacks := ""

	for i := 0; i < len(stackA); i++ {
		for j := 0; j < len(stackB); j++ {
			if isPair(stackA[i], stackB[j]) {
				if !foundOne {
					p1, p2 = stackA[i], stackB[j]
					distOld = min(distances(stackA, p1)) + min(distances(stackB, p2))
					stacks = "AB"
					foundOne = true
				} else {
					distNow := min(distances(stackA, stackA[i])) + min(distances(stackB, stackB[j]))

					if distNow < distOld { // New pair is closer than the old one
						p1, p2 = stackA[i], stackB[j]
						stacks = "AB"
					}
				}
			}
		}
	}

	for i := 0; i < len(stackA)-1; i++ {
		for j := i + 1; j < len(stackA); j++ {
			if isPair(stackA[i], stackA[j]) {
				if !foundOne {
					p1, p2 = stackA[i], stackA[j]
					distOld = min(distances(stackA, p1)) + min(distances(stackA, p2)) + 1
					stacks = "AA"
					foundOne = true
				} else {
					distNow := min(distances(stackA, stackA[i])) + min(distances(stackA, stackA[j])) + 1

					if distNow < distOld { // New pair is closer than the old one
						p1, p2 = stackA[i], stackA[j]
						stacks = "AA"
					}
				}
			}
		}
	}

	for i := 0; i < len(stackB)-1; i++ {
		for j := i + 1; j < len(stackB); j++ {
			if isPair(stackB[i], stackB[j]) && j-i != 1 {
				if !foundOne {
					p1, p2 = stackB[i], stackB[j]
					distOld = min(distances(stackB, p1)) + min(distances(stackB, p2)) + 1
					stacks = "BB"
					foundOne = true
				} else {
					distNow := min(distances(stackB, stackB[i])) + min(distances(stackB, stackB[j])) + 1

					if distNow < distOld { // New pair is closer than the old one
						p1, p2 = stackB[i], stackB[j]
						stacks = "BB"
					}
				}
			}
		}
	}

	if foundOne {
		return p1, p2, stacks, nil
	} else {
		return p1, p2, stacks, fmt.Errorf("no pairs found")
	}

}

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
func bestOrder(allOrders *[][]int) []int {
	best := (*allOrders)[0]
	for _, o := range *allOrders {
		if len(o) > len(best) {
			best = o
		}
		if len(o) == len(best) && o[0] < best[0] {
			best = o
		}
	}
	return best
}

// getAllOrders finds all sequences on increasing numbers through stackA starting from the element at index start
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
		return
	}

	//search all hidden orders for all the bigger numbers
	for _, n := range biggers {
		getAllOrders(start, n, curSolution, orders)
	}
}
