/*
** go-basic
**
** Made by zheng_b
*/

package main

import "fmt"
import "math"

const delta = 1e-8

func Sqrt(x float64) float64 {
	z := x
	prev := 0.0	
	for {
		z = (z - ((math.Pow(z, 2) - x) / (2 * z)))
		if math.Abs(prev - z) < delta {
			break
		}
		prev = z
	}
	return z
}

func display_line() {
	fmt.Println("==================================================")
}

func main() {
	display_line()
	for i := 1; i <= 10; i++ {
		value := float64(i)
		real_result := math.Sqrt(value)
		newton_result := Sqrt(value)		
        fmt.Println("Calculating Sqrt for\t", value)
        fmt.Println("Real result\t\t", real_result)
        fmt.Println("Newton result\t\t", newton_result)
		display_line()
	}
}