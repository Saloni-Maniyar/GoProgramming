package main

import "fmt"

type employee struct {
	eno   int
	esal  float32
	ename string
}

func main() {
	var e1 [100]employee
	var n, k int
	var max float32
	fmt.Println("Enter no of Employees you want:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter no:")
		fmt.Scan(&e1[i].eno)
		fmt.Println("Enter name:")
		fmt.Scan(&e1[i].ename)
		fmt.Println("Enter salary:")
		fmt.Scan(&e1[i].esal)

	}
	max = e1[0].esal
	for i := 0; i < n; i++ {
		if e1[i].esal > max {
			max = e1[i].esal
			k = i
		}
	}
	fmt.Println("Employee having maximum salary:")
	fmt.Println("Empno=", e1[k].eno)
	fmt.Println("Empname=", e1[k].ename)
	fmt.Println("Empsal=", e1[k].esal)
}
