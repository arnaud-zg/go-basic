/*
** go-basic
**
** Made by zheng_b
*/

package main

import "fmt"

func fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		i, j = j, (i + j)
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}