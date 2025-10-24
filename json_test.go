package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	tests := []struct {
		name           string
		code           int
		payload        interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid JSON response",
			code:           http.StatusOK,
			payload:        map[string]string{"status": "ok"},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok"}`,
		},
		{
			name:           "Created status with payload",
			code:           http.StatusCreated,
			payload:        map[string]string{"id": "123", "name": "test"},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":"123","name":"test"}`,
		},
		{
			name:           "Empty object",
			code:           http.StatusOK,
			payload:        map[string]string{},
			expectedStatus: http.StatusOK,
			expectedBody:   `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respondWithJSON(w, tt.code, tt.payload)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			contentType := w.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, w.Body.String())
			}
		})
	}
}

func TestRespondWithError(t *testing.T) {
	tests := []struct {
		name           string
		code           int
		msg            string
		logErr         error
		expectedStatus int
	}{
		{
			name:           "Client error (400)",
			code:           http.StatusBadRequest,
			msg:            "Invalid request",
			logErr:         errors.New("validation failed"),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Server error (500)",
			code:           http.StatusInternalServerError,
			msg:            "Internal server error",
			logErr:         errors.New("database connection failed"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Error without log error",
			code:           http.StatusNotFound,
			msg:            "Resource not found",
			logErr:         nil,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respondWithError(w, tt.code, tt.msg, tt.logErr)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			contentType := w.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
			}

			var response struct {
				Error string `json:"error"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Error != tt.msg {
				t.Errorf("Expected error message '%s', got '%s'", tt.msg, response.Error)
			}
		})
	}
}
