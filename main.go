package main

import "fmt"
func main() {
	fmt.Println(hello("sooraj"))
	newmain();
}

func hello(name string) string {
	return "hello"+" "+name+"!!"
}