package main

import (
	"os"
	"path/filepath"
	"strings"
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

	fileInfo, err = filepath.EvalSymlinks(filename)
	if err != nil {
		return FileInfo{}, err
	}

	expandedPath, err := filepath.EvalSymlinks(filename)
	if err != nil {
		return FileInfo{}, err
	}

	dir, file := filepath.Split(filename)
	ext := filepath.Ext(file)
	base := strings.TrimSuffix(file, ext)

	return FileInfo{
		Name:         fileInfo.Name(),
		Path:         filename,
		ExpandedPath: expandedPath.String(),
		Size:         fileInfo.Size(),
		Mode:         fileInfo.Mode(),
		ModTime:      fileInfo.ModTime(),
		IsDir:        fileInfo.IsDir(),
		IsSymlink:    fileInfo.Mode()&os.ModeSymlink != 0,
		Volume:       fileInfo.Sys().(*syscall.Win32FileAttributeData).VolumeSerialNumber,
		SymlinkDest:  fileInfo.SymlinkDest(),
		Directory:    dir,
		Basename:     base,
		Extension:    ext,
	}, nil
}
