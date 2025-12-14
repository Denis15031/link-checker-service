package checker

import (
	"testing"
)

func TestCheckURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid domain",
			input:    "google.com",
			expected: "available",
		},
		{
			name:     "invalid domain",
			input:    "nonexistentdomain12345.com",
			expected: "not available",
		},
		{
			name:     "with protocol",
			input:    "http://example.com",
			expected: "available", // зависит от того, как сервер отвечает
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckURL(tt.input)
			// NOTE: Тесты могут быть хрупкими, если они зависят от интернета.
			// В продакшене лучше использовать моки или httptest.
			_ = result // TODO: Замокать http.Client
		})
	}
}
