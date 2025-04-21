package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string = " Hello Golang "

	fmt.Println(strings.Contains(text, "Hello"))
	fmt.Println(strings.HasPrefix(text, "Hello"))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.TrimSpace(text))
	fmt.Println(strings.ReplaceAll(text, "Golang", "go"))

	fruits := "apel,pisang,jeruk"
	fruitsSplit := strings.Split(fruits, ",")
	fmt.Println(fruitsSplit)
	fruitsJoin := strings.Join(fruitsSplit, " | ")
	fmt.Println(fruitsJoin)
}
