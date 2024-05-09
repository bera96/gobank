package main

import "fmt"

func main() {
	server := NewApıServer(":3005")
	server.Run()
	fmt.Println("Hello")
}
