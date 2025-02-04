package main

import (
    "testing"
    "fmt"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "  hello world ",
            expected: []string{"hello", "world"},
        },
        {
            input: "  this_should    split into 5 words..",
            expected: []string{"this_should", "split", "into", "5", "words.."},
        },
        {
            input: "*****",
            expected: []string{"*****"},
        },
        {
            input: "",
            expected: []string{},
        },
    }

    for i, c := range cases {
        fmt.Printf("======= CASE %v ========\n", i)
        actual := cleanInput(c.input)
        // Check the length of the actual slice
        // if they don't match, use t.Errorf to print an error message
        // and fail the test
        if len(actual) != len(c.expected) {
            t.Errorf("Length not matching\n")
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            // Check each word in the slice
            // if they don't match, use t.Errorf to print an error message
            // and fail the test
            if word != expectedWord {
                t.Errorf("Failed match %v : %v expected\n", word, expectedWord)
            }
        }
    }
}
