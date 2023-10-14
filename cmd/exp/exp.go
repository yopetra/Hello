package main

import (
	"errors"
	"fmt"
)

func main() {
	err := B()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("this is ErrNotFound type")
	}
	fmt.Printf("unwrapped: %v\n", errors.Unwrap(err))
}

var ErrNotFound = errors.New("not found")

func A() error {
	return ErrNotFound
}

func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}
