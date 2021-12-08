package main

import "fmt"

var x, y int

func main() {
	fmt.Printf("Enter the value:\n", x)
	fmt.Scanf("%d\n", &x)
	fmt.Printf("Enter the value:", y)
	fmt.Scanf("%d\n", &y)
	c := x + y
	fmt.Printf("The sum is %d", c)
	fmt.Scanf("%d\n", &c)

	if c%2 == 0 {
		fmt.Printf("The value is even")
	} else {
		fmt.Printf("The value is odd")
	}

}
