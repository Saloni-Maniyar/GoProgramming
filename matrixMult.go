package main

import "fmt"

func main() {
	var m, n, p, q, total int
	var a [5][5]int
	var b [5][5]int
	var c [5][5]int
	fmt.Println("Enter the order of first Matrix:")
	fmt.Scanln(&m, &n)
	fmt.Println("Enter the order of second matrix:")
	fmt.Scanln(&p, &q)
	if n != p {
		fmt.Println("The matrix can not be multiplied!!")
	} else {
		fmt.Println("Enter the First matrix Elements:")
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Scan(&a[i][j])
			}
		}
		fmt.Println("Enter the Second matrix Elements:")
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Scan(&b[i][j])
			}
		}
		for i := 0; i < m; i++ {
			for j := 0; j < p; j++ {
				for k := 0; k < p; k++ {
					total = total + a[i][k]*b[k][j]
				}
				c[i][j] = total
				total = 0
			}
		}
		fmt.Println("Multiplication Matrix:")
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(" ", c[i][j])
			}
			fmt.Print("\n")
		}

	}
}
