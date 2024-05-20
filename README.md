These files were created using the following prompt.

```
Given the following struct
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

write a function that will populate this struct given a file name and write a test for it.
```
