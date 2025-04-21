package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}

	return a / b, nil
}
func main() {
	hasil, err := divide(10, 2)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("hasil", hasil)
	}
}
