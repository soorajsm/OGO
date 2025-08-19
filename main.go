package main

import "fmt"

//globall variable

var myslice = []int{1, 2, 3}

func main() {
	// fmt.Println(hello("sooraj"))
	// newmain();
	// arraynslice();
	myslice=append(myslice,4)
	check()
}

func check (){
	fmt.Println(myslice)
}

// func hello(name string) string {
// 	return "hello"+" "+name+"!!"
// }
