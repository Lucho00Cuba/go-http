package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewServer tests the NewServer function
func TestNewServer(t *testing.T) {
	// Define the test cases
	tests := []struct {
		port     string
		expected string
	}{
		{"8080", "8080"},
		{"9090", "9090"},
		{"", ""}, // Caso para puerto vacío
	}

	for _, test := range tests {
		// Create a new server instance
		server := NewServer(test.port)

		// Check if the server's port matches the expected value
		if server.Port != test.expected {
			t.Errorf("NewServer(%q) = %q; want %q", test.port, server.Port, test.expected)
		}
	}
}

// Test the handler function for various status codes
func TestHandler(t *testing.T) {
	tests := []struct {
		path           string
		expectedStatus int
	}{
		{"/200/test", http.StatusOK},
		{"/404/test", http.StatusNotFound},
		{"/500/test", http.StatusInternalServerError},
		{"/invalid/test", http.StatusNotFound},
		{"/302/test", http.StatusFound},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.path, nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc((&Server{}).handler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
		}
	}
}

// Test logging of requests
func TestLogRequest(t *testing.T) {
	// Create a test server
	srv := httptest.NewServer((&Server{Port: "8080"}).logRequest((&Server{Port: "8080"}).handler))
	defer srv.Close()

	// Send a request to the test server
	resp, err := http.Get(srv.URL + "/200/test")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

// Test isValidStatusCode function
func TestIsValidStatusCode(t *testing.T) {
	validCodes := []int{100, 200, 300, 400, 500}
	invalidCodes := []int{99, 599, 600}

	for _, code := range validCodes {
		if !isValidStatusCode(code) {
			t.Errorf("Expected status code %d to be valid", code)
		}
	}

	for _, code := range invalidCodes {
		if isValidStatusCode(code) {
			t.Errorf("Expected status code %d to be invalid", code)
		}
	}
}
