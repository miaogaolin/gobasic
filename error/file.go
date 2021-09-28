package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

var ErrEmpty = errors.New("内容为空")

type FileEmptyError struct {
	Filename string
	Err      error
}

func (e *FileEmptyError) Error() string {
	return fmt.Sprintf("%s %v", e.Filename, e.Err)
}

func (e *FileEmptyError) Unwrap() error {
	return e.Err
}

func LoadConfig() (string, error) {
	filename := "./config.json"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("读取文件出错:%w", err)
	}

	content := string(b)
	if len(content) == 0 {
		return "", &FileEmptyError{
			Filename: filename,
			Err:      errors.New("内容为空"),
		}
	}
	return content, nil
}
