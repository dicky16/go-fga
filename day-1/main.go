package main

import (
	"fmt"

	"github.com/dicky16/go-fga/user"
)

var class string = "ok"

func main() {
	class = "okok"
	print()
	//call function user
	user.User("ok")
	fmt.Printf("nama : %v", 9)
	fmt.Println()
	read()
}

func read() {
	fmt.Println(class)
}
