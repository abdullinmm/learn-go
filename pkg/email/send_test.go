package email

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	test := []struct {
		user   User
		name   string
		answer string
	}{
		{
			user:   User{Name: "Marsel", Email: "abdullinmm@gmail.com"},
			name:   "Send message",
			answer: "Sending email for Marsel for email addres abdullinmm@gmail.com"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := fmt.Sprintf("%v", SendEmail(&tt.user))
			if result != tt.answer {
				t.Errorf("Expected %q, got %q", tt.answer, result)
				t.Fail()
			}
		})
	}
}
