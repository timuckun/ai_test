package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPopulateFileInfo(t *testing.T) {
	filename, err := filepath.Abs("./testdata/file.txt")
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	fileInfo, err := PopulateFileInfo(filename)
	if err != nil {
		t.Fatalf("PopulateFileInfo failed: %v", err)
	}

	if fileInfo.Name != "file.txt" {
		t.Errorf("Incorrect Name: got %q, want %q", fileInfo.Name, "file.txt")
	}

	if fileInfo.Path != filename {
		t.Errorf("Incorrect Path: got %q, want %q", fileInfo.Path, filename)
	}

	if fileInfo.ExpandedPath != filename {
		t.Errorf("Incorrect ExpandedPath: got %q, want %q", fileInfo.ExpandedPath, filename)
	}

	info, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}

	if fileInfo.Size != info.Size() {
		t.Errorf("Incorrect Size: got %d, want %d", fileInfo.Size, info.Size())
	}

	if fileInfo.Mode != info.Mode() {
		t.Errorf("Incorrect Mode: got %v, want %v", fileInfo.Mode, info.Mode())
	}

	if !fileInfo.ModTime.Equal(info.ModTime()) {
		t.Errorf("Incorrect ModTime: got %v, want %v", fileInfo.ModTime, info.ModTime())
	}

	if fileInfo.IsDir != info.IsDir() {
		t.Errorf("Incorrect IsDir: got %t, want %t", fileInfo.IsDir, info.IsDir())
	}

	if fileInfo.IsSymlink != (info.Mode()&os.ModeSymlink != 0) {
		t.Errorf("Incorrect IsSymlink: got %t, want %t", fileInfo.IsSymlink, (info.Mode()&os.ModeSymlink != 0))
	}

	// Add more checks for other fields if needed
}
