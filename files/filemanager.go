package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileManager contains methods to interact with files and the file system
//go:generate mockgen -destination=mock_files/filemanager.go . FileManager
type FileManager interface {
	// FilepathExists returns true if the file or directory exists, false otherwise
	FilepathExists(filePath string) (exists bool, err error)
	// FileExists returns true if the path is an existing file, false otherwise
	FileExists(filePath string) (exists bool, err error)
	// DirectoryExists returns true if the path is an existing directory, false otherwise
	DirectoryExists(filePath string) (exists bool, err error)
	// GetOwnership obtains the user ID and group ID for a file or directory
	GetOwnership(filePath string) (userID, groupID int, err error)
	// SetOwnership sets the user ID and group ID for a file or directory
	SetOwnership(filePath string, userID, groupID int) error
	// GetUserPermissions obtains the permissions for a file or directory
	GetUserPermissions(filePath string) (read, write, execute bool, err error)
	// SetUserPermissions sets the permissions for a file or directory
	SetUserPermissions(filepath string, mod os.FileMode) error
	// ReadFile reads an entire file and returns its data
	ReadFile(filePath string) (data []byte, err error)
	// WriteLinesToFile writes some lines to a file
	WriteLinesToFile(filePath string, lines []string, setters ...WriteOptionSetter) error
	// Touch creates a file at the specified file path
	Touch(filePath string, setters ...WriteOptionSetter) error
	// WriteToFile writes some data to a file
	WriteToFile(filePath string, data []byte, setters ...WriteOptionSetter) error
	// CreateDir creates a directory at the file path specified
	CreateDir(filePath string, setters ...WriteOptionSetter) error
	// Remove removes a file or directory at the file path specified
	Remove(filePath string) (err error)
}

type fileManager struct {
	fileStat    func(name string) (os.FileInfo, error)
	isNotExist  func(err error) bool
	readFile    func(filename string) ([]byte, error)
	filepathDir func(path string) string
	mkdirAll    func(path string, perm os.FileMode) error
	writeFile   func(filename string, data []byte, perm os.FileMode) error
	chown       func(name string, uid int, gid int) error
	chmod       func(name string, mod os.FileMode) error
	rm          func(path string) error
}

func NewFileManager() FileManager {
	return &fileManager{
		fileStat:    os.Stat,
		isNotExist:  os.IsNotExist,
		readFile:    ioutil.ReadFile,
		filepathDir: filepath.Dir,
		mkdirAll:    os.MkdirAll,
		writeFile:   ioutil.WriteFile,
		chown:       os.Chown,
		chmod:       os.Chmod,
		rm:          os.RemoveAll,
	}
}
