package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 123456"},
			},
			wantAPIKey: "123456",
			wantErr:    false,
		},
		{
			name: "malformed auth header",
			headers: http.Header{
				"Authorization": []string{"123456"},
			},
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() gotApiKey = %v, want %v,", gotAPIKey, tt.wantAPIKey)
			}
		})
	}
}
