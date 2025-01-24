package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args[0]) == 0 {
		return
	}
	if len(args) > 1 {
		log.Fatalln("Use with one argument, e.g.: ./push-swap \"7 12 0 31 3\" ")
	}

	instructions, err := produceInstructions(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, ins := range instructions {
		fmt.Println(ins)
	}

	//fmt.Println(len(instructions), "instructions to sort")

	// test if the stack got sorted
	/* fmt.Println(len(instructions), "instructions")
	if reflect.DeepEqual(stacks.StackA, stacks.ASorted) {
		fmt.Println("Stack A is sorted")
	} else {
		fmt.Println("Stack A is NOT sorted")
	} */

}
