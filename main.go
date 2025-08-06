package main

import "fmt"
func main() {
	fmt.Println(hello("sooraj"))
}

func hello(name string) string {
	return "hello"+" "+name+"!!"
}