package hello

import "fmt"

// Hello returns string "hello, world"
func Hello() string {
	return "hello, world"
}

// HelloString says hello to the parameter
func HelloString(hello string) string {
	return fmt.Sprintf("hello, %s", hello)
}

// Bonjour is French
func Bonjour(hello string) string {
	return fmt.Sprintf("Bonjour, %s.", hello)
}
