// version_test.go
package version

import (
	"testing"
)

// TestVersion checks if the VERSION variable has the expected value
func TestVersion(t *testing.T) {
	// Set the expected version value
	expectedVersion := "v1.0.0"

	// Temporarily set the VERSION variable to the expected value for testing
	// Note: This is just for the example; normally, you set VERSION using -ldflags
	VERSION = expectedVersion

	// Check if the VERSION variable matches the expected value
	if VERSION != expectedVersion {
		t.Errorf("VERSION = %q; want %q", VERSION, expectedVersion)
	}
}
