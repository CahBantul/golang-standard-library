package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello \n World!")

	br := bufio.NewReader(r)

	line, err := br.ReadString('!')
	if err != nil {
		fmt.Println("errot", err)
	}

	fmt.Println(line)

	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w, os.Stdin)
	fmt.Fprint(w, "World")

	w.Flush()
}
