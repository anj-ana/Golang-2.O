package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var y float64
	var z float64
	var a float64
	fmt.Printf("Enter the value of x:\n")
	fmt.Scanf("%f\n", &x)
	fmt.Printf("Enter the value of y:\n")
	fmt.Scanf("%f\n", &y)
	fmt.Printf("Enter the value of z:\n")
	fmt.Scanf("%f\n", &z)
	fmt.Printf("Enter the value of a:\n")
	fmt.Scanf("%f\n", &a)

	c := ((math.Sqrt(x) + math.Sqrt(y)) * z) / a
	fmt.Printf("((Sqrt(%f) +Sqrt(%f)) * %f)/ %f = %f ", x, y, z, a, c)

}
