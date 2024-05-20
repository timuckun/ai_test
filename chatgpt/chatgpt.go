package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileInfo represents information about a file.
type FileInfo struct {
	Name         string      // Name of the file
	Path         string      // Full path of the file
	ExpandedPath string      // Expanded path of the file
	Size         int64       // Size of the file in bytes
	Mode         os.FileMode // Permission and mode bits
	ModTime      time.Time   // Modification time
	IsDir        bool        // Whether the file is a directory
	IsSymlink    bool        // Whether the file is a symbolic link
	Volume       string      // Volume of the file
	SymlinkDest  string      // Destination of symbolic link (if it is a symlink)
	Directory    string      // Directory of the file
	Basename     string      // Basename of the file
	Extension    string      // Extension of the file
}

// GetFileInfo populates a FileInfo struct for the given filename
func GetFileInfo(filename string) (*FileInfo, error) {
	// Get the absolute path
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	// Get file information
	fileInfo, err := os.Lstat(absPath)
	if err != nil {
		return nil, err
	}

	// Check if the file is a symlink
	isSymlink := fileInfo.Mode()&os.ModeSymlink != 0

	// Resolve symlink destination if it is a symlink
	symlinkDest := ""
	if isSymlink {
		symlinkDest, err = filepath.EvalSymlinks(absPath)
		if err != nil {
			return nil, err
		}
	}

	// Split the path into directory and base components
	directory := filepath.Dir(absPath)
	basename := filepath.Base(absPath)

	// Get the file extension
	extension := filepath.Ext(absPath)

	// Get the volume name (Windows specific, empty string on Unix)
	volume := ""
	if vol := filepath.VolumeName(absPath); vol != "" {
		volume = vol
	}

	return &FileInfo{
		Name:         fileInfo.Name(),
		Path:         absPath,
		ExpandedPath: absPath, // Assuming ExpandedPath is the same as absPath
		Size:         fileInfo.Size(),
		Mode:         fileInfo.Mode(),
		ModTime:      fileInfo.ModTime(),
		IsDir:        fileInfo.IsDir(),
		IsSymlink:    isSymlink,
		Volume:       volume,
		SymlinkDest:  symlinkDest,
		Directory:    directory,
		Basename:     basename,
		Extension:    extension,
	}, nil
}
