package string_to_integer

import (
	// "fmt"
	"testing"
)

func TestAtoi(t *testing.T) {
	vales := []struct {
		str string
		n   int
	}{
		{"123", 123},
		{"-123", -123},
		{"+", 0},
		{"+-2", 0},
		{"++1", 0},
		{"--1", 0},
		{"-+2", 0},
		{"234dfg89", 0},
		{"", 0},
	}

	for _, v := range vales {
		if v.n != MyAtoi(v.str) {
			t.Log(v.str, "failed.")
		}
	}
}
