package main

import (
	"fmt"
	"strconv"
)

func main() {
	umur := 25
	s := strconv.Itoa(umur)
	fmt.Println("umur saya : " + s)

	angka, err := strconv.Atoi("123154")
	if err == nil {
		fmt.Println("hasil", angka)
	}

	b, _ := strconv.ParseBool("true")
	fmt.Println("b", b)
	strconv.
}
