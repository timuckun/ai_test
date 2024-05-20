package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPopulateFileInfo(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_file")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Populate the FileInfo struct
	fileInfo, err := PopulateFileInfo(tempFile.Name())
	if err != nil {
		t.Fatalf("Error populating FileInfo: %v", err)
	}

	// Verify the populated fields
	if fileInfo.Name != filepath.Base(tempFile.Name()) {
		t.Errorf("Expected Name to be %s, got %s", filepath.Base(tempFile.Name()), fileInfo.Name)
	}
	if fileInfo.Path != tempFile.Name() {
		t.Errorf("Expected Path to be %s, got %s", tempFile.Name(), fileInfo.Path)
	}
	if fileInfo.ExpandedPath != filepath.Abs(tempFile.Name()) {
		t.Errorf("Expected ExpandedPath to be %s, got %s", filepath.Abs(tempFile.Name()), fileInfo.ExpandedPath)
	}
	if fileInfo.Size != 0 {
		t.Errorf("Expected Size to be 0, got %d", fileInfo.Size)
	}
	if fileInfo.Mode != 0 {
		t.Errorf("Expected Mode to be 0, got %d", fileInfo.Mode)
	}
	if fileInfo.ModTime.IsZero() {
		t.Errorf("Expected ModTime to be non-zero, got %v", fileInfo.ModTime)
	}
	if fileInfo.IsDir {
		t.Errorf("Expected IsDir to be false, got true")
	}
	if fileInfo.IsSymlink {
		t.Errorf("Expected IsSymlink to be false, got true")
	}
	if fileInfo.Volume != filepath.VolumeName(tempFile.Name()) {
		t.Errorf("Expected Volume to be %s, got %s", filepath.VolumeName(tempFile.Name()), fileInfo.Volume)
	}
	if fileInfo.SymlinkDest != "" {
		t.Errorf("Expected SymlinkDest to be empty, got %s", fileInfo.SymlinkDest)
	}
	if fileInfo.Directory != filepath.Dir(tempFile.Name()) {
		t.Errorf("Expected Directory to be %s, got %s", filepath.Dir(tempFile.Name()), fileInfo.Directory)
	}
	if fileInfo.Basename != filepath.Base(tempFile.Name()) {
		t.Errorf("Expected Basename to be %s, got %s", filepath.Base(tempFile.Name()), fileInfo.Basename)
	}
	if fileInfo.Extension != filepath.Ext(tempFile.Name()) {
		t.Errorf("Expected Extension to be %s, got %s", filepath.Ext(tempFile.Name()), fileInfo.Extension)
	}
}
