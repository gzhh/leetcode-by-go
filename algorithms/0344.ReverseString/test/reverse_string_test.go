package test

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T)  {
	var tests = []struct {
		str, want string
	}{
		{"hello", "olleh"},
		{"orange", "egnaro"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.str)
		t.Run(testname, func(t *testing.T) {
			ans := ReverseString(tt.str)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}