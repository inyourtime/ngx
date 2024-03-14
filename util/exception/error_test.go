package exception

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	expected := &Exception{
		Type:    TypeInternal,
		Message: "test error",
		Cause:   nil,
		Errors:  make(map[string][]string),
	}

	err := New(TypeInternal, "test error", nil)
	assert.NotNil(t, err, "New() returned nil error, expected an error")
	assert.Equal(t, expected, err, "New() returned incorrect error")
}

func TestValidation(t *testing.T) {
	expected := &Exception{
		Type:    TypeValidation,
		Message: "validation error",
		Cause:   nil,
		Errors:  make(map[string][]string),
	}

	err := Validation()
	assert.NotNil(t, err, "Validation() returned nil error, expected an error")
	assert.Equal(t, expected, err, "Validation() returned incorrect error")
}

func TestException_HasError(t *testing.T) {
	ex := &Exception{Errors: make(map[string][]string)}
	assert.False(t, ex.HasError(), "Expected HasError() to return false for empty Errors map")

	ex.AddError("key", "message")
	assert.True(t, ex.HasError(), "Expected HasError() to return true for non-empty Errors map")
}

func TestException_AddError(t *testing.T) {
	ex := &Exception{Errors: make(map[string][]string)}
	ex.AddError("key", "message")
	expected := map[string][]string{"key": {"message"}}
	assert.Equal(t, expected, ex.Errors, "AddError() didn't add error as expected")

	// Adding another error for the same key
	ex.AddError("key", "another message")
	expected = map[string][]string{"key": {"message", "another message"}}
	assert.Equal(t, expected, ex.Errors, "AddError() didn't add multiple errors for the same key")
}

func TestInto(t *testing.T) {
	cases := []struct {
		name     string
		err      error
		expected *Exception
	}{
		{
			name: "New error",
			err:  errors.New("test error"),
			expected: &Exception{
				Type:    TypeInternal,
				Message: "test error",
				Cause:   errors.New("test error"),
				Errors:  make(map[string][]string),
			},
		},
		{
			name:     "Nil error",
			err:      nil,
			expected: nil,
		},
		{
			name: "Exception error",
			err:  New(TypePermissionDenied, "Permission denied", nil),
			expected: &Exception{
				Type:    TypePermissionDenied,
				Message: "Permission denied",
				Cause:   nil,
				Errors:  make(map[string][]string),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ex := Into(tc.err)
			assert.Equal(t, tc.expected, ex, "Into() returned incorrect error")
		})
	}
}

func TestException_Error(t *testing.T) {
	t.Run("Has cause", func(t *testing.T) {
		causeErr := errors.New("cause error")
		ex := &Exception{
			Type:    TypeInternal,
			Message: "test error",
			Cause:   causeErr,
			Errors:  make(map[string][]string),
		}

		assert.Equal(t, causeErr.Error(), ex.Error(), "Error() method returned incorrect error message")
	})

	t.Run("no cause", func(t *testing.T) {
		ex := &Exception{
			Type:    TypeInternal,
			Message: "test error",
			Cause:   nil,
			Errors:  make(map[string][]string),
		}

		assert.Equal(t, "", ex.Error(), "Error() method returned incorrect error message")
	})
}
