package main

import "fmt"

func main() {
	var os = []string{"Windows", "Linux", "Apple"}
	fmt.Println(len(os))
	os = append(os, "Mac")
	fmt.Println(os)

	//range of slice
	Range := os[1:3]
	fmt.Println(Range)
	Range = os[1:]
	fmt.Println(Range)
	Range = os[:4]
	fmt.Println(Range)

	//for loop
	for i := 0; i < len(os); i++ {
		fmt.Println(os[i])
	}
}
