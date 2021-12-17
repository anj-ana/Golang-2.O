package main

import "fmt"

func updateName(x *string) {
	*x = "Sarah"
}

func main() {
	name := "Anju"
	m := &name
	fmt.Println(name)
	fmt.Println(m) //address of name
	fmt.Println(*m)
	updateName(m)
	fmt.Println(name)
	//fmt.Println(name)
	//fmt.Println(&name)
}
