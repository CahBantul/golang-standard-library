package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("masukkan sesuatu")

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("gagal membuat file", err)
		return
	}

	defer file.Close()

	_, err = io.Copy(file, os.Stdin)
	if err != nil {
		fmt.Println("gagal copy file", err)
		return
	}

	fmt.Println("Isi berhasil disimpan ke output.txt")
}
