package main

import (
    "os"
    "path/filepath"
    "time"
)

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

func PopulateFileInfo(filename string) (FileInfo, error) {
    fi, err := os.Stat(filename)
    if err != nil {
        return FileInfo{}, err
    }

    expandedPath, err := filepath.EvalSymlinks(filename)
    if err != nil {
        return FileInfo{}, err
    }

    volume := filepath.VolumeName(filename)
    directory := filepath.Dir(filename)
    basename := filepath.Base(filename)
    extension := filepath.Ext(basename)

    return FileInfo{
        Name:         fi.Name(),
        Path:         filename,
        ExpandedPath: expandedPath,
        Size:         fi.Size(),
        Mode:         fi.Mode(),
        ModTime:      fi.ModTime(),
        IsDir:        fi.IsDir(),
        IsSymlink:    fi.Mode()&os.ModeSymlink != 0,
        Volume:       volume,
        SymlinkDest:  fi.Mode()&os.ModeSymlink != 0 ? expandedPath : "",
        Directory:    directory,
        Basename:     basename,
        Extension:    extension,
    }, nil
}
