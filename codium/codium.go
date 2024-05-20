package main

import (
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name         string
	Path         string
	ExpandedPath string
	Size         int64
	Mode         os.FileMode
	ModTime      time.Time
	IsDir        bool
	IsSymlink    bool
	Volume       string
	SymlinkDest  string
	Directory    string
	Basename     string
	Extension    string
}

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
