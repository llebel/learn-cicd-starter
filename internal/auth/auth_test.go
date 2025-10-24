package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		want      string
		wantError bool
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key-123"},
			},
			want:      "test-api-key-123",
			wantError: false,
		},
		{
			name:      "No Authorization Header",
			headers:   http.Header{},
			want:      "",
			wantError: true,
		},
		{
			name: "Malformed Authorization Header - Missing ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key-123"},
			},
			want:      "",
			wantError: true,
		},
		{
			name: "Malformed Authorization Header - Missing key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:      "",
			wantError: true,
		},
		{
			name: "Empty Authorization Header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			want:      "",
			wantError: true,
		},
		{
			name: "Valid API Key with extra spaces",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-key-with-spaces"},
			},
			want:      "my-key-with-spaces",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantError {
				t.Errorf("GetAPIKey() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrNoAuthHeaderIncluded(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
