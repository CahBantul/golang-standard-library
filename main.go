package main

import "fmt"

type User struct {
	Name      string
	Age       int
	IsMarried bool
}

func main() {
	user := User{Name: "fardan", Age: 30, IsMarried: true}

	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%s\n", user.Name)
	fmt.Printf("%d\n", user.Age)
	fmt.Printf("%t\n", user.IsMarried)
	fmt.Print("halo")
	fmt.Println("halo")
	fmt.Print("halo\n")
	pesan := fmt.Sprintf("1+1= %d", 1+1)
	fmt.Println(pesan)

	var name string

	fmt.Print("masukkan nama kamu: ")
	fmt.Scanln(&name)
	fmt.Printf("halo, %s \n", name)
}
