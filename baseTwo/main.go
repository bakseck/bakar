package main

import "fmt"

func main() {

	var a int
	var r int
	var b []int
	fmt.Scan(&a)

	for a != 0 {
		r = a % 2
		b = append(b, r)
		a = a / 2 
	}
	
	for i:=len(b)-1; i>=0; i-- {
		fmt.Print(b[i])
	}
}