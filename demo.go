package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scanln(&m, &n)

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&res[i][j])
		}
	}
	fmt.Println(res)
}
