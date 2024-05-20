package main

import (
	"os"
	"path/filepath"
	"syscall"
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

// PopulateFileInfo populates a FileInfo struct with information about a file.
func PopulateFileInfo(filename string) (FileInfo, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return FileInfo{}, err
	}

	// Get expanded path
	expandedPath, err := filepath.Abs(filename)
	if err != nil {
		return FileInfo{}, err
	}

	// Get directory and basename
	directory, basename := filepath.Split(filename)

	// Get extension
	extension := filepath.Ext(basename)

	return FileInfo{
		Name:         fileInfo.Name(),
		Path:         filename,
		ExpandedPath: expandedPath,
		Size:         fileInfo.Size(),
		Mode:         fileInfo.Mode(),
		ModTime:      fileInfo.ModTime(),
		IsDir:        fileInfo.IsDir(),
		IsSymlink:    fileInfo.Mode()&os.ModeSymlink != 0,
		Volume:       filepath.VolumeName(filename),
		SymlinkDest:  fileInfo.Sys().(*syscall.Stat_t).X__st_ino, // Use Sys() for symlink destination
		Directory:    directory,
		Basename:     basename,
		Extension:    extension,
	}, nil
}
