package main

import (
	"fmt"
	"io"
)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}


func main() {
	
}