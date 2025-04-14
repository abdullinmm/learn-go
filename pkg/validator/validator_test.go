package validator

import (
	"learn-go/pkg/models"
	"testing"
)

func TestValidator(t *testing.T) {
	test := []struct {
		// user models.User
		validator Validator
		expected  error
		name      string
	}{
		{
			validator: &models.User{Name: "Marsel", Email: "abdullinmm@gmail.com", Age: 20},
			expected:  nil,
			name:      "Valid user"},
		{
			validator: &models.User{Name: "Marsel", Email: "abdullinmm", Age: 18},
			expected:  models.ErrInvalidEmail,
			name:      "Invalid email (no @)"},
		{
			validator: &models.User{Name: "Marsel", Email: "abdullinmm@gmail.com", Age: 17},
			expected:  models.ErrUnderage,
			name:      "Underage user"},
		{
			validator: &models.User{Name: "Marsel", Email: "abdullinmm", Age: 17},
			expected:  models.ErrInvalidEmail,
			name:      "Multiple errors (invalid email and underage)",
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.validator.Validate(); err != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, err)
				t.Fail()
			}
		})
	}
}

// Тест для функции ValidatorUser
// func TestValidatorUser(t *testing.T) {
// 	type mockValidator struct{}

// 	var validatorCalled bool

// 	mockValidator.validator = func() error {
// 		validatorCalled = true
// 		return nil
// 	}

// 	err := validator.ValidatorUser(mockValidator{})
// 	assert.NoError(t, err)
// 	assert.True(t, validatorCalled)
// }
