package main

import (
	"fmt"

	"github.com/almamedia/go-mysql-datetime/hello"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello.HelloString("Susan"))
	fmt.Println(hello.Bonjour("Susan"))
}
