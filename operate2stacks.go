package main

import "fmt"

func twoStacksMethod(ins []string) []string {

	// start by putting any two elements to stack B
	ins = append(ins, "pb", "pb")
	runComms([]string{"pb", "pb"})

	// sort the two elements
	if stackB[0] < stackB[1] {
		ins = append(ins, "sb")
		runComm("sb")
	}

	// Put some elements to B. B will be at least 3 long after this.
	for stackA[0] > stackB[0] || stackA[0] < last(stackB) {

		if stackA[0] > stackB[0] { // stackA[0] to first
			ins = append(ins, "pb")
			runComm("pb")
		} else if stackA[0] < last(stackB) { // stackA[0] to last
			ins = append(ins, "pb", "rb")
			runComms([]string{"pb", "rb"})
		} else if len(stackA) > 0 && len(stackB) > 1 && stackA[0] < stackB[0] && stackA[0] > stackB[1] { // stackA[0] to second
			ins = append(ins, "rb", "pb", "rrb")
			runComms([]string{"rb", "pb", "rrb"})
		}

		if len(stackA) < 1 || len(stackB) < 1 {
			break
		}
	}

	// find pairs and push one from A to B
	for len(stackA) > 0 {
		// Find nearest pair
		p1, p2, stacks, err := nearestPair()
		//fmt.Println(stacks, stackB)

		if err != nil {
			break
		}

		// Move the bigger of the two to the top of A , the other to the top of B and push from A to B
		if stacks == "AB" {
			runComms(toTop(p1, stackA, "A"))
			ins = append(ins, toTop(p1, stackA, "A")...)

			runComms(toTop(p2, stackB, "B"))
			ins = append(ins, toTop(p2, stackB, "B")...)

			runComm("pb")
			ins = append(ins, "pb")
		}

		// Move the smaller of the two to the top of A, the next smallest on B to
		// the top there and push from A to B. Then move the other to top of A and push.
		if stacks == "AA" {
			runComms(toTop(p2, stackA, "A"))
			ins = append(ins, toTop(p2, stackA, "A")...)

			/* 			runComms(toTop(nextSmallestOnList(stackB, p2), stackB, "B"))
			   			ins = append(ins, toTop(nextSmallestOnList(stackB, p2), stackB, "B")...) */

			runComm("pb")
			ins = append(ins, "pb")

			runComms(toTop(p1, stackA, "A"))
			ins = append(ins, toTop(p1, stackA, "A")...)

			runComm("pb")
			ins = append(ins, "pb")
		}

		// Move the bigger of the two to the top of B and push it to A, move the smaller to the top of B and push from A to B
		if stacks == "BB" {
			runComms(toTop(p1, stackB, "B"))
			ins = append(ins, toTop(p1, stackB, "B")...)

			runComm("pa")
			ins = append(ins, "pa")

			runComms(toTop(p2, stackB, "B"))
			ins = append(ins, toTop(p2, stackB, "B")...)

			runComm("pb")
			ins = append(ins, "pb")
		}

	}

	/* 	fmt.Println("A:", stackA)
	   	fmt.Println("B:", stackB) */

	// put the biggest one top on Stack B
	runComms(toTop(biggestOnList(stackB), stackB, "B"))
	ins = append(ins, toTop(biggestOnList(stackB), stackB, "B")...)

	// push all of B back to A
	for i := 0; len(stackB) > 0; i++ {
		var prevPush int
		first := true

		if len(stackB) > 1 {
			if nextOnList(aSorted, stackB[i]) == prevPush || first {
				runComm("pa")
				ins = append(ins, "pa")
				prevPush = stackB[i]
				first = false
				i--
			} else {
				runComms(toTop(biggestOnList(stackA), stackA, "A"))
				ins = append(ins, toTop(biggestOnList(stackA), stackA, "A")...)

			}
		} else {
			runComm("pa")
			ins = append(ins, "pa")
		}

	}

	runComms(toTop(smallestOnList(stackA), stackA, "A"))
	ins = append(ins, toTop(smallestOnList(stackA), stackA, "A")...)

	return ins
}

//previousOnList returns the value of the previous element on stack st
func previousOnList(st []int, val int) int {
	out := -1
	for i, n := range st {
		if n == val && i != 0 {
			out = st[i-1]
		}
	}
	return out
}

//nextSmallestOnList returns the value of the next element on stack st
func nextOnList(st []int, val int) int {
	out := -1
	for i, n := range st {
		if n == val && i != len(st)-1 {
			out = st[i+1]
		}
	}
	return out
}

// nearestPair finds the pair that can be joined with the fewest moves on stacks A and B
// Pairs are two values, one on A and the other on B, that are subsequent on a sorted
// list and when pushed with "pb", will be ordered correctly on B (big to small)
func nearestPair() (int, int, string, error) {
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

func isPair(a, b int) bool {
	for i := 1; i < len(aSorted); i++ {
		if aSorted[i] == a && aSorted[i-1] == b {
			return true
		}
	}
	return false
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func doSwap() (comms []string) {
	if len(stackA) > 1 && stackA[0] > stackA[1] {
		comms = append(comms, "sa")
		runComm("sa")
	}
	if len(stackB) > 1 && stackB[0] < stackB[1] {
		comms = append(comms, "sb")
		runComm("sb")
	}
	return
}
