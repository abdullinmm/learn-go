package main

import (
	"learn-go/pkg/strings"
	"testing"
)

func TestRevers(t *testing.T) {
	test := []struct {
		str, expected string
	}{
		{"ABC", "CBA"},
		{"ABC qwe", "ewq CBA"},
	}

	for _, tt := range test {
		result, _ := strings.Revers(tt.str)
		if result != tt.expected {
			t.Fail()
		}
	}
}
