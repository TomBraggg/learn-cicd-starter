package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
	}{
		{
			name:        "Valid API Key",
			headers:     http.Header{"Authorization": []string{"ApiKey AuthKey123"}},
			expectedKey: "AuthKey123",
			expectError: false,
		},
		{
			name:        "Missing Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Malformed Authorization Header - No ApiKey Prefix",
			headers:     http.Header{"Authorization": []string{"Bearer AuthKey123"}},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Malformed Authorization Header - No Key",
			headers:     http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got error: %v", tt.expectError, err)
			}
			if !reflect.DeepEqual(gotKey, tt.expectedKey) {
				t.Errorf("expected key: %s, got key: %s", tt.expectedKey, gotKey)
			}
		})
	}
}
