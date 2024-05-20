package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGetFileInfo(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	// Get file information using the function
	fileInfo, err := GetFileInfo(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}

	// Get expected values
	expectedName := filepath.Base(tmpFile.Name())
	expectedAbsPath, _ := filepath.Abs(tmpFile.Name())
	expectedDirectory := filepath.Dir(expectedAbsPath)
	expectedExtension := filepath.Ext(tmpFile.Name())
	expectedMode := os.FileMode(0666) // default mode for a created file, may differ based on umask

	// Check file information
	if fileInfo.Name != expectedName {
		t.Errorf("Expected Name: %s, got: %s", expectedName, fileInfo.Name)
	}
	if fileInfo.Path != expectedAbsPath {
		t.Errorf("Expected Path: %s, got: %s", expectedAbsPath, fileInfo.Path)
	}
	if fileInfo.Directory != expectedDirectory {
		t.Errorf("Expected Directory: %s, got: %s", expectedDirectory, fileInfo.Directory)
	}
	if fileInfo.Basename != expectedName {
		t.Errorf("Expected Basename: %s, got: %s", expectedName, fileInfo.Basename)
	}
	if fileInfo.Extension != expectedExtension {
		t.Errorf("Expected Extension: %s, got: %s", expectedExtension, fileInfo.Extension)
	}
	if fileInfo.Mode != expectedMode {
		t.Errorf("Expected Mode: %v, got: %v", expectedMode, fileInfo.Mode)
	}
	if fileInfo.IsDir {
		t.Errorf("Expected IsDir: false, got: %v", fileInfo.IsDir)
	}
	if fileInfo.IsSymlink {
		t.Errorf("Expected IsSymlink: false, got: %v", fileInfo.IsSymlink)
	}
	if fileInfo.Size != 0 {
		t.Errorf("Expected Size: 0, got: %d", fileInfo.Size)
	}
	if fileInfo.Volume != "" && fileInfo.Volume != filepath.VolumeName(expectedAbsPath) {
		t.Errorf("Expected Volume: %s, got: %s", filepath.VolumeName(expectedAbsPath), fileInfo.Volume)
	}

	// Check ModTime with a tolerance
	modTime := fileInfo.ModTime
	if time.Since(modTime) > time.Minute {
		t.Errorf("ModTime is too old: %v", modTime)
	}
}
