package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "ajitama", "masukkan namamu")
	umur := flag.Int("age", 0, "masukkan umur")

	flag.Parse()
	fmt.Println("nama:", *name)
	fmt.Println("umur:", *umur)
}
