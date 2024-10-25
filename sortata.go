package main

import "fmt"

var bigHalf []int // Should we update this along the way?

func makeBigHalf(n int) {
	change := (len(aSorted) - n - 2) / 2
	bigHalf = aSorted[len(aSorted)/2-change : len(aSorted)-1]
}

func sortAtAMethod(ins []string) []string {
	ind := len(aSorted) - 2 // start from second-to-last, biggest one stays on A
	makeBigHalf(ind)

	for ind >= 0 {
		// push elements to stack B
		if isOnList(aSorted[ind], stackA) {
			posD, negD := distances(stackA, aSorted[ind])
			if float64(posD) < float64(negD)*1.5 { // favor positive
				ins = append(ins, toBPos(ind)...)
			} else {
				ins = append(ins, toBNeg(ind)...)
			}
		}

		// rotate stack B so target element is first
		if len(stackB) > 0 && stackB[0] != aSorted[ind] {
			posD, negD := distances(stackB, aSorted[ind])

			if float64(posD) <= float64(negD) {
				ins = append(ins, rotStack(posD, "rb")...)
			} else {
				ins = append(ins, rotStack(negD, "rrb")...)
			}
		}

		if len(stackB) > 0 && stackB[0] != aSorted[ind] {
			fmt.Println("\nTarget not first on stack B!")
			panic("Oh no!")
		}

		// move the element back to A, to correct index
		if stackA[0] != aSorted[ind+1] { // rotate stack A so it's right
			posD, negD := distances(stackA, aSorted[ind+1])
			if float64(posD) <= float64(negD) {
				ins = append(ins, rotStack(posD, "ra")...)
			} else {
				ins = append(ins, rotStack(negD, "rra")...)
			}
		}

		ins = append(ins, "pa")
		runComm("pa")

		ind--
		makeBigHalf(ind)
	}

	return ins
}

// toBPos pushes element at index ind to stack B, as well as any other encountered element on the bigger half while reversing
func toBPos(ind int) []string {
	comms := []string{}
	target := aSorted[ind]

	for i := 0; i < len(stackA); i++ {
		//fmt.Println(stackA[0])
		if stackA[0] == target { // stop rotating when element is found
			comms = append(comms, "pb")
			runComm("pb")
			break
		} else {
			if isOnList(stackA[0], bigHalf) && stackA[0] < target { // push bigger-half elements to stack B
				comms = append(comms, "pb")
				runComm("pb")
			} else {
				comms = append(comms, "ra") // pushing effectively rotated stack A to positive direction already
				runComm("ra")
			}
		}
		if i == len(stackA)-1 {
			fmt.Println("Element not found going positive!")
		}
	}
	return comms
}

// toBNeg pushes element at index ind (at beginning) to stack B, as well as any other encountered element on the bigger half while reversing
func toBNeg(ind int) []string {
	comms := []string{}
	target := aSorted[ind]

	for i := 0; i < len(stackA); i++ {

		if stackA[0] == target { // stop rotating when element is found
			comms = append(comms, "pb")
			runComm("pb")
			break
		} else {
			if isOnList(stackA[0], bigHalf) && stackA[0] < target { // push bigger-half elements to stack B
				comms = append(comms, "pb")
				runComm("pb")
			}
			comms = append(comms, "rra") // reverse rotation has to run whether element was pushed or not
			runComm("rra")
		}

		if i == len(stackA)-1 {
			fmt.Println("Element not found going negative!")
		}
	}
	return comms
}
