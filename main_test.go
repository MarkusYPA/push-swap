package main

import (
	"log"
	"os"
	"testing"
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
			if tc.expected < result {
				t.Errorf("\n%s Input was \"%s\"\nWant fewer than %v instructions\ngot: %v", tc.name, tc.input, tc.expected+1, result)
			}
		})
	}
}
