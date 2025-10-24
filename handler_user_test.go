package main

import (
	"testing"
)

func TestGenerateRandomSHA256Hash(t *testing.T) {
	// Test that the function generates a hash
	hash1, err := generateRandomSHA256Hash()
	if err != nil {
		t.Fatalf("generateRandomSHA256Hash() failed: %v", err)
	}

	// Check that the hash is not empty
	if hash1 == "" {
		t.Error("generateRandomSHA256Hash() returned an empty string")
	}

	// Check that the hash is the correct length (SHA256 is 64 hex characters)
	if len(hash1) != 64 {
		t.Errorf("Expected hash length of 64, got %d", len(hash1))
	}

	// Test that multiple calls generate different hashes
	hash2, err := generateRandomSHA256Hash()
	if err != nil {
		t.Fatalf("generateRandomSHA256Hash() failed on second call: %v", err)
	}

	if hash1 == hash2 {
		t.Error("generateRandomSHA256Hash() generated the same hash twice")
	}

	// Test that the hash contains only valid hex characters
	for _, char := range hash1 {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f')) {
			t.Errorf("Hash contains invalid character: %c", char)
		}
	}
}
