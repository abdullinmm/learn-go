package strings

import "testing"

func TestRevers(t *testing.T) {
	test := []struct {
		name, str, expected string
	}{
		{"String with out space", "ABC", "CBA"},
		{"String with space", "ABC qwe", "ewq CBA"},
		{"Empty string", "", ""},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Revers(tt.str)
			if result != tt.expected {
				t.Fail()
			}
		})
	}
}
