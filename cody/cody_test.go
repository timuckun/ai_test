package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPopulateFileInfo(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_file_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Get the expected FileInfo values
	expectedName := filepath.Base(tempFile.Name())
	expectedPath := tempFile.Name()
	expectedExpandedPath, _ := filepath.EvalSymlinks(tempFile.Name())
	expectedSize := int64(0)
	expectedMode := os.FileMode(0600)
	expectedModTime := tempFile.ModTime()
	expectedIsDir := false
	expectedIsSymlink := false
	expectedVolume := filepath.VolumeName(tempFile.Name())
	expectedSymlinkDest := ""
	expectedDirectory := filepath.Dir(tempFile.Name())
	expectedBasename := expectedName
	expectedExtension := filepath.Ext(expectedName)

	// Call the PopulateFileInfo function
	fileInfo, err := PopulateFileInfo(tempFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the FileInfo values match the expected values
	if fileInfo.Name != expectedName {
		t.Errorf("Unexpected Name: got %q, want %q", fileInfo.Name, expectedName)
	}
	if fileInfo.Path != expectedPath {
		t.Errorf("Unexpected Path: got %q, want %q", fileInfo.Path, expectedPath)
	}
	if fileInfo.ExpandedPath != expectedExpandedPath {
		t.Errorf("Unexpected ExpandedPath: got %q, want %q", fileInfo.ExpandedPath, expectedExpandedPath)
	}
	if fileInfo.Size != expectedSize {
		t.Errorf("Unexpected Size: got %d, want %d", fileInfo.Size, expectedSize)
	}
	if fileInfo.Mode != expectedMode {
		t.Errorf("Unexpected Mode: got %v, want %v", fileInfo.Mode, expectedMode)
	}
	if !fileInfo.ModTime.Equal(expectedModTime) {
		t.Errorf("Unexpected ModTime: got %v, want %v", fileInfo.ModTime, expectedModTime)
	}
	if fileInfo.IsDir != expectedIsDir {
		t.Errorf("Unexpected IsDir: got %v, want %v", fileInfo.IsDir, expectedIsDir)
	}
	if fileInfo.IsSymlink != expectedIsSymlink {
		t.Errorf("Unexpected IsSymlink: got %v, want %v", fileInfo.IsSymlink, expectedIsSymlink)
	}
	if fileInfo.Volume != expectedVolume {
		t.Errorf("Unexpected Volume: got %q, want %q", fileInfo.Volume, expectedVolume)
	}
	if fileInfo.SymlinkDest != expectedSymlinkDest {
		t.Errorf("Unexpected SymlinkDest: got %q, want %q", fileInfo.SymlinkDest, expectedSymlinkDest)
	}
	if fileInfo.Directory != expectedDirectory {
		t.Errorf("Unexpected Directory: got %q, want %q", fileInfo.Directory, expectedDirectory)
	}
	if fileInfo.Basename != expectedBasename {
		t.Errorf("Unexpected Basename: got %q, want %q", fileInfo.Basename, expectedBasename)
	}
	if fileInfo.Extension != expectedExtension {
		t.Errorf("Unexpected Extension: got %q, want %q", fileInfo.Extension, expectedExtension)
	}
}
