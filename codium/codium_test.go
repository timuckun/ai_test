package main

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"
)

func TestPopulateFileInfoWithValidPath(t *testing.T) {
	// Setup: Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	// Test
	info, err := PopulateFileInfo(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify FileInfo is populated correctly
	if info.Name != filepath.Base(tmpFile.Name()) {
		t.Errorf("Expected name %s, got %s", filepath.Base(tmpFile.Name()), info.Name)
	}
	if info.Size != 0 {
		t.Errorf("Expected size 0, got %d", info.Size)
	}
	if info.IsDir {
		t.Errorf("Expected IsDir false, got true")
	}
}

func TestPopulateFileInfoWithInvalidPath(t *testing.T) {
	// Test with a non-existent file path
	_, err := PopulateFileInfo("/path/to/nonexistent/file")
	if err == nil {
		t.Errorf("Expected an error for non-existent file path, got nil")
	}
}

func TestPopulateFileInfoWithRestrictedPermissions(t *testing.T) {
	// Setup: Create a file with restricted permissions
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	// Restrict permissions
	if err := os.Chmod(tmpFile.Name(), 0000); err != nil {
		t.Fatalf("Failed to set file permissions: %s", err)
	}

	// Test
	_, err = PopulateFileInfo(tmpFile.Name())
	if err == nil {
		t.Errorf("Expected permission error, got nil")
	}

	// Cleanup: Restore permissions to delete the file
	os.Chmod(tmpFile.Name(), 0666)
}
