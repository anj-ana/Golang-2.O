package main

import "fmt"

func main() {
	x := 6
	y := 2
	c := x > y
	fmt.Printf("%t\n", c)
	d := x < y
	fmt.Printf("%t\n", d)
	e := x >= y
	fmt.Printf("%t\n", e)
	f := x <= y
	fmt.Printf("%t\n", f)
	g := x == y
	fmt.Printf("%t\n", g)
	h := x != y
	fmt.Printf("%t\n", h)
	i := x > y && x == y
	fmt.Printf("%t\n", i)
	j := !(x <= y)
	fmt.Printf("%t\n", j)
	k := x == y || x < y
	fmt.Printf("%t\n", k)

}
