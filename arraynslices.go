package main

import "fmt"

func arraynslice(){
	var arr = [3]string{"one","two","three"}
	var slice = []string{"one","two","three","four"}

	for index,value:=range arr {
		fmt.Println(index,value)
	}

	for index,value:=range slice {
		fmt.Println(index,value)
	}
}