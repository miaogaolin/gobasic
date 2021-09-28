package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFoundRequest = errors.New("404")
	ErrBadRequest      = errors.New("请求异常")
)

func GetError() error {
	// ...
	// 错误 1
	return ErrNotFoundRequest

	// ...
	// 错误 2
	return ErrBadRequest

	// ...
	// 错误 3
	path := "https://printlove.com"
	return fmt.Errorf("%s:%w", path, ErrNotFoundRequest)

	// ...
}

func HandleError() {
	err := GetError()
	if errors.Is(err, ErrNotFoundRequest) {
		// 错误 1,错误3
	} else if errors.Is(err, ErrBadRequest) {
		// 错误 2
	}
}
