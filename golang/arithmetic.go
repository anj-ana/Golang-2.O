package main

import "fmt"

var a, b int32

func main() {
	fmt.Println("Enter the value:")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Enter the value:")
	fmt.Scanf("%d\n", &b)
	c := a + b
	fmt.Printf("%d\n", c)
	d := a * b
	fmt.Printf("%d\n", d)
	e := a - b
	fmt.Printf("%d\n", &e)

}
