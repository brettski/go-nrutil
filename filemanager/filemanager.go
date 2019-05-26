package filemanager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Filemanager struct maintains path and current file
type Filemanager struct {
	BasePath    string
	currentfile *os.FileInfo
}

// NewFilemanager creates and returns an instance of Filemanager
// basepath is the working directory for our files
func NewFilemanager(basepath string, isCreate bool) (*Filemanager, error) {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(basepath)
	if err != nil {
		if os.IsNotExist(err) && isCreate {
			err = os.MkdirAll(basepath, os.ModePerm)
			if err != nil {
				return nil, fmt.Errorf("error in os.mkdirall %s", err)
			}
			fileInfo, err = os.Stat(basepath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	log.Printf("basepath, fileinfo: %v", fileInfo)

	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("The basepath, %s, is not a directory", basepath)
	}
	log.Println("sending struct")
	return &Filemanager{
		BasePath: basepath,
	}, nil
}

// Exists checks for the existence of a file based on basepath of Filemanager
// filename is the name of file to check if exists
func (fm *Filemanager) Exists(filename string) bool {
	if len(filename) < 1 && fm.currentfile == nil {
		log.Printf("No filename provided or available to check if exists")
	}

	// set new filename
	fileinfo, err := os.Stat(fmt.Sprintf("%s/%s", fm.BasePath, filename))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal("Unable to access file. ", err)
		return false
	}
	fm.currentfile = &fileinfo
	return true

}

// WriteFile write a new or overwrites a current file indicated by filename.
// file path is dermined by Filemanager.BasePath.
func (fm *Filemanager) WriteFile(filename string, data []byte) error {
	if len(filename) < 1 {
		return fmt.Errorf("WriteFile, filename cannot be an empty string")
	}

	fullfile := filepath.Join(fm.BasePath, filename)
	log.Printf("fullfile: %s", fullfile)

	wfile, err := os.OpenFile(
		fullfile,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer wfile.Close()

	writtenb, err := wfile.Write(data)
	if err != nil {
		return err
	}

	log.Printf("Wrote %d bytes to %s", writtenb, wfile.Name())
	return nil
}
