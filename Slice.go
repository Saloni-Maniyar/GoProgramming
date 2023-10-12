package main

import "fmt"

func main() {
	str := []string{
		"Mango", "Apple", "Banana",
	}
	fmt.Println(str)
	a := make([]string, len(str))
	s1 := append(str, "pineapple")
	fmt.Println(s1)
	fmt.Println(copy(a, str))
	fmt.Printf("%V", a)
	fmt.Println(a)
}
