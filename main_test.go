package main

import (
	"strconv"
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

type testCaseError struct {
	name     string
	input    string
	expected string
}

type testCaseGood struct {
	name     string
	input    string
	expected int
}

var testCasesError = []testCaseError{
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
	{
		name:     "100 random numbers A",
		input:    "158 8 140 126 148 86 74 106 93 40 92 11 182 57 49 37 91 97 118 112 172 3 43 28 176 38 17 186 83 23 109 96 21 85 198 46 134 18 103 80 102 51 137 192 35 77 120 115 199 179 95 151 183 39 15 123 152 132 159 2 61 4 153 156 133 7 122 128 150 100 19 200 72 59 189 14 27 98 84 174 52 20 164 195 5 69 116 131 187 147 33 9 136 79 163 82 162 196 42 16",
		expected: 699,
	},
	{
		name:     "100 random numbers B",
		input:    "24 147 176 79 128 170 18 172 50 129 109 26 49 88 124 180 140 186 96 165 104 113 164 21 28 197 23 198 43 175 37 137 89 10 171 70 47 183 41 67 72 7 141 193 52 155 71 73 34 63 22 122 166 25 107 123 185 81 62 163 142 190 74 106 177 38 150 132 162 15 126 95 99 92 4 30 169 125 12 77 69 119 103 139 188 58 174 11 135 93 173 40 39 14 16 75 157 160 108 33",
		expected: 699,
	},
}

func TestProduceInstructionsBad(t *testing.T) {
	for _, tc := range testCasesError {
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

	// 50 sequences of 50 numbers in max 250 turns each? Not an audit question. Not without hidden order optimization.
	for _, tc := range makeCases(50, 50, 250) {
		t.Run(tc.name, func(t *testing.T) {
			ins, _ := produceInstructions(tc.input)
			result := len(ins)
			if result > tc.expected {
				t.Errorf("\n%s Input was \"%s\"\nWant fewer than %v instructions\ngot: %v", tc.name, tc.input, tc.expected+1, result)
			}
		})
	}

	// 20 sequences of 100 numbers in max 699 turns each? ~1/500 requires optimization. 11.1.2025
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

// makeCases makes a slice of an 'amount' number of valid test cases each 'length' long that
// should be sorted in an 'expected' number of turns.
func makeCases(amount, length, expected int) []testCaseGood {

	rand.Seed(uint64(time.Now().UnixNano()))

	cases := make([]testCaseGood, amount)
	for i := range amount {

		name := strconv.Itoa(length) + "-random-" + strconv.Itoa(i+1)

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
