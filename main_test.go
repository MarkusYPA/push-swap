package main

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

type testCaseBad struct {
	name     string
	input    string
	expected string
}

type testCaseGood struct {
	name     string
	input    string
	expected int
}

var testCasesBad = []testCaseBad{
	{
		name:     "Number as word",
		input:    "0 one 2 3",
		expected: "Error",
	},
	{
		name:     "Duplicate number",
		input:    "1 2 2 3",
		expected: "Error",
	},
	{
		name:     "6 given numbers",
		input:    "2 1 3 6 5 8",
		expected: "",
	},
	{
		name:     "6 sorted numbers",
		input:    "0 1 2 3 4 5",
		expected: "",
	},
	{
		name:     "5 random numbers A",
		input:    "96 63 17 88 77",
		expected: "",
	},
	{
		name:     "5 random numbers B",
		input:    "61 36 29 15 13",
		expected: "",
	},
}

var testCasesGood = []testCaseGood{
	{
		name:     "empty",
		input:    "",
		expected: 0,
	},
	{
		name:     "6 given numbers",
		input:    "2 1 3 6 5 8",
		expected: 8,
	},
	{
		name:     "6 sorted numbers",
		input:    "0 1 2 3 4 5",
		expected: 0,
	},
	{
		name:     "5 random numbers",
		input:    "96 63 17 88 77",
		expected: 11,
	},
	{
		name:     "5 random numbers",
		input:    "61 36 29 15 13",
		expected: 11,
	},
	// These test cases take about 1 min each
	/* 	{
		name:     "100 random numbers A",
		input:    fileToString("testcases/100a.txt"),
		expected: 699,
	}, */
	/* 	{
		name:     "100 random numbers B",
		input:    fileToString("testcases/100b.txt"),
		expected: 699,
	}, */
}

func fileToString(s string) string {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(file)
}

func TestProduceInstructionsBad(t *testing.T) {
	for _, tc := range testCasesBad {
		t.Run(tc.name, func(t *testing.T) {
			_, err := produceInstructions(tc.input)
			var result string
			if err != nil {
				result = err.Error()
			}
			if tc.expected != result {
				t.Errorf("\n%s Input was \"%s\"\nWant the error message: \"%v\"\ngot: \"%v\"", tc.name, tc.input, tc.expected, result)
			}
		})
	}
}

func TestProduceInstructionsGood(t *testing.T) {
	for _, tc := range testCasesGood {
		t.Run(tc.name, func(t *testing.T) {
			ins, _ := produceInstructions(tc.input)
			result := len(ins)
			if result > tc.expected {
				t.Errorf("\n%s Input was \"%s\"\nWant fewer than %v instructions\ngot: %v", tc.name, tc.input, tc.expected+1, result)
			}
		})
	}
}

func TestProduceInstructionsRandom(t *testing.T) {

	// Can 1000 random sequences of five numbers be arranged in max 11 turns each?
	// About 1/30 requires hidden order optimization 11.1.2025
	for _, tc := range makeCases(1000, 5, 11) {
		t.Run(tc.name, func(t *testing.T) {
			ins, _ := produceInstructions(tc.input)
			result := len(ins)
			if result > tc.expected {
				t.Errorf("\n%s Input was \"%s\"\nWant fewer than %v instructions\ngot: %v", tc.name, tc.input, tc.expected+1, result)
			}
		})
	}

	// 50 numbers in max 250 turns? Not an audit question. Not without hidden order optimization.
	for _, tc := range makeCases(50, 50, 250) {
		t.Run(tc.name, func(t *testing.T) {
			ins, _ := produceInstructions(tc.input)
			result := len(ins)
			if result > tc.expected {
				t.Errorf("\n%s Input was \"%s\"\nWant fewer than %v instructions\ngot: %v", tc.name, tc.input, tc.expected+1, result)
			}
		})
	}

	// 100 numbers in max 699 turns? ~1/500 requires optimization. 11.1.2025
	for _, tc := range makeCases(20, 100, 699) {
		t.Run(tc.name, func(t *testing.T) {
			ins, _ := produceInstructions(tc.input)
			result := len(ins)
			if result > tc.expected {
				t.Errorf("\n%s Input was \"%s\"\nWant fewer than %v instructions\ngot: %v", tc.name, tc.input, tc.expected+1, result)
			} else {
				//fmt.Println(result)
			}
		})
	}
}

func makeCases(amount, length, expected int) []testCaseGood {

	rand.Seed(uint64(time.Now().UnixNano()))

	cases := make([]testCaseGood, amount)
	for i := range amount {

		name := "random" + strconv.Itoa(i)

		numbers := make([]int, length)
		for i := 0; i < length; i++ {
			numbers[i] = i
		}
		rand.Shuffle(len(numbers), func(i, j int) { // Randomize order
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})

		input := ""
		for _, num := range numbers {
			input += strconv.Itoa(num) + " "
		}

		cases[i].name, cases[i].expected, cases[i].input = name, expected, input[:len(input)-1]
	}

	return cases
}
