package main

import "fmt"

func secretOrder(ins []string) []string {

	allOrders := [][]int{}
	for i := range stackA {
		//fmt.Println("Now for", stackA[i], "at index", i)
		getAllOrders(i, i, []int{}, &allOrders)
	}

	fmt.Println("Length of all orders:", len(allOrders))
	/* 	for _, o := range allOrders {
		fmt.Println(o)
	} */

	bestO := bestOrder(&allOrders)
	fmt.Println("The best one:", bestO)

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

	//fmt.Println(index, biggers)

	// If no more biggers found, deduce this solution is complete
	if len(biggers) == 0 {
		*orders = append(*orders, curSolution)
		return
	}

	//search all hidden orders for all the bigger numbers
	for _, n := range biggers {
		getAllOrders(start, n, curSolution, orders)
	}
}
