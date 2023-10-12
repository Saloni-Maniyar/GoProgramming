package main

import "fmt"

type Student struct {
	sno        int
	m1, m2, m3 int
	sname      string
}

func main() {
	var s1 [10]Student
	var total, avg float64
	fmt.Println("Enter the Student Roll No:")
	fmt.Scan(&s1[0].sno)
	fmt.Println("Enter the Student Name:")
	fmt.Scan(&s1[0].sname)

	fmt.Println("Enter the student marks for 3 subjects:")
	fmt.Scan(&s1[0].m1)
	fmt.Scan(&s1[0].m2)
	fmt.Scan(&s1[0].m3)
	fmt.Println("Roll No=", s1[0].sno)

	fmt.Println("Name=", s1[0].sname)
	fmt.Println("Marks of 3 subjects", s1[0].m1, s1[0].m2, s1[0].m3)
	total = float64(s1[0].m1 + s1[0].m2 + s1[0].m3)
	avg = float64(total / 3)
	fmt.Println("Total Number:", total)
	fmt.Println("Average marks=", avg)
}
