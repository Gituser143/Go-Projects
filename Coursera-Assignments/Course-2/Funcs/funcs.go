package main

import "fmt"

func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * t * t) + (v_o * t) + s_o
	}

	return fn
}

func main() {

	var a, v, s, t float64

	fmt.Println("Enter Acceleration value")
	fmt.Scanf("%f", &a)

	fmt.Println("Enter Velocity value")
	fmt.Scanf("%f", &v)

	fmt.Println("Enter Initial Displacement value")
	fmt.Scanf("%f", &s)

	fn := GenDisplaceFn(a, v, s)

	fmt.Println("Enter Time value")
	fmt.Scanf("%f", &t)

	fmt.Println("Displacement:", fn(t))
}
