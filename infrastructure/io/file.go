package io

import (
	"io/ioutil"
	"os"
)

type FileWriter interface {
	Write(path, content string) error
}

type FileReader interface {
	Read(path string) (string, error)
}

type FileAccessor interface {
	FileWriter
	FileReader
	Exist(path string) bool
}

type LocalFileAccessor struct {
}

func (f *LocalFileAccessor) Exist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func NewFileAccessor() *LocalFileAccessor {
	return &LocalFileAccessor{}
}

func (f *LocalFileAccessor) Read(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (f *LocalFileAccessor) Write(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}
