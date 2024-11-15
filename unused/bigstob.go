package main

func bigsToB(ins []string) []string {
	bigs := aSorted[len(aSorted)/2-1 : len(aSorted)-1]
	// Put bigger half to B
	ins = append(ins, moveToB(bigs)...)

	// Rotate biggest to top of A
	ins = append(ins, toTop(aSorted[len(aSorted)-1], stackA, "A")...)

	// move bigger half back to A in order
	for len(stackB) > 0 {
		target := biggestOnList(stackB)
		ins = append(ins, toTop(target, stackB, "B")...)
		if stackB[0] == target {
			ins = append(ins, "pa")
			runComm("pa")
		}
	}

	// move rest to B
	rest := aSorted[:len(aSorted)/2-1]
	ins = append(ins, moveToB(rest)...)

	// move smaller half back to A in order
	for len(stackB) > 0 {
		target := biggestOnList(stackB)
		ins = append(ins, toTop(target, stackB, "B")...)
		if stackB[0] == target {
			ins = append(ins, "pa")
			runComm("pa")
		}
	}

	return ins
}

func moveToB(bigs []int) []string {
	comms := []string{}
	for _, n := range stackA {
		if isOnList(n, bigs) {
			comms = append(comms, "pb")
			runComm("pb")
		} else {
			comms = append(comms, "ra")
			runComm("ra")
		}
		if len(stackB) == len(bigs) {
			break
		}
	}
	return comms
}

func biggestOnList(s []int) int {
	big := s[0]
	for _, n := range s {
		if n > big {
			big = n
		}
	}
	return big
}

func last(s []int) int {
	return s[len(s)-1]
}
