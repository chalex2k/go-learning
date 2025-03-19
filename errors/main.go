package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	src := "go is awesome"
	_, err := indexOf(src, "js")
	if err, ok := err.(lookupError); ok {
		fmt.Println("err.src:", err.src)
		fmt.Println("err.substr:", err.substr)
	}
}

func dangerous() error {
	return errors.New("dangerous")

}

// error - интерфейс с одним методом Error() string

// Собственный тип ошибки
type lookupError struct {
	src    string
	substr string
}

func (e lookupError) Error() string {
	return fmt.Sprintf("'%s' not found in '%s'", e.substr, e.src)
}

func indexOf(src string, substr string) (int, error) {
	idx := strings.Index(src, substr)
	if idx == -1 {
		// Создаем и возвращаем ошибку типа `lookupError`.
		return -1, lookupError{src, substr}
	}
	return idx, nil
}
