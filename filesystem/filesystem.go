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
	secretKey string
}

func NewFileManager() *FileManager {
	return &FileManager{}
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
	data := []byte(content)
	if f.secretKey != "" {
		data = xorWithSecretKey(data, f.secretKey)
	}
	ioutil.WriteFile(path, data, 0666)
	return f.ReadFile(path)
}

func (f *FileManager) ReadFile(path string) (*FileInfo, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if f.secretKey != "" {
		data = xorWithSecretKey(data, f.secretKey)
	}
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		Filename:       fi.Name(),
		FilePath:       path,
		FileSize:       fi.Size(),
		FileUpdateTime: fi.ModTime().Format("2006-01-02 15:04:05"),
		FileContent:    string(data),
	}, nil
}

func (f *FileManager) UpdateFile(path string, content string) (*FileInfo, error) {
	data := []byte(content)
	if f.secretKey != "" {
		data = xorWithSecretKey(data, f.secretKey)
	}
	err := ioutil.WriteFile(path, data, 0666)
	if err != nil {
		return nil, err
	}
	return f.ReadFile(path)
}

func (f *FileManager) DeleteFile(path string) error {
	return os.Remove(path)
}

func (f *FileManager) SetSecretKey(secretKey string) {
	f.secretKey = secretKey
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

func xorWithSecretKey(data []byte, key string) []byte {
	dl, kl := len(data), len(key)
	for i := 0; i < dl; i++ {
		data[i] ^= key[i%kl]
	}
	return data
}
