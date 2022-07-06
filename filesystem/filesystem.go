package filesystem

import (
	"io/ioutil"
	"os"
	"regexp"
)

type FileInfo struct {
	Filename       string
	FilePath       string
	FileSize       int64
	FileUpdateTime string
	FileContent    string
}

type FileManager struct {
}

func (f *FileManager) CreateFile(path string, content string) (*FileInfo, error) {
	re := regexp.MustCompile(`^(.*/)[^/]+$`)
	folder := re.ReplaceAllString(path, "$1")
	if path != folder && !isExist(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}
	_, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile(path, []byte(content), 0666)
	return f.ReadFile(path)
}

func (f *FileManager) ReadFile(path string) (*FileInfo, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	stream, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		Filename:       fi.Name(),
		FilePath:       path,
		FileSize:       fi.Size(),
		FileUpdateTime: fi.ModTime().Format("2006-01-02 15:04:05"),
		FileContent:    string(stream),
	}, nil
}

func (f *FileManager) UpdateFile(path string, content string) (*FileInfo, error) {
	err := ioutil.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return nil, err
	}
	return f.ReadFile(path)
}

func (f *FileManager) DeleteFile(path string) error {
	return os.Remove(path)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
