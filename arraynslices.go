package main

import "fmt"

func arraynslice(){
	var arr = [3]string{"one","two","three"}
	var slice = []string{"one","two","three","four"}
	fmt.Println(arr)
	fmt.Println(slice)
}