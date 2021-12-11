package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("***Map***")
	places := make(map[string]int) //map is a keyword with a key and a vaue
	food := make(map[string]int)

	places["place1"] = rand.Intn(50)
	places["place2"] = rand.Intn(50)

	food["food1"] = rand.Intn(100)
	food["food2"] = rand.Intn(100)

	fmt.Println(places["place1"])
	fmt.Println(places["place2"])
	fmt.Println(food["food1"])
	fmt.Println(food["food2"])

	//example 2
	Names := map[string]string{"name": "Anju", "branch": "cse", "skills": "Java & Golang"}
	fmt.Println(Names["branch"])
}
