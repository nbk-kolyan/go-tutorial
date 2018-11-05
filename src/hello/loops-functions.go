package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	s := float64(1)

	for i:= 1; i<= 10; i++ {
		s = s - ((s*s)-x)/(2*s)
	}
	return s;
}

func Sqrt2(x float64) float64 {
	z, s := 1.0, float64(0)
	e := 1e-15  //pogreshnost
	i := 0
	for math.Abs(z-s) > e{
		s = z
		z = z - ((z*z)-x)/(2*z)
		i++
	}
	fmt.Printf("Sqrt2  -- for x:%v needed %v iterations \n", x, i)
	return z
}

func Sqrt(x float64) float64 {
	z := float64(1.)
	s := float64(0)
	i := 0
	for {
		z = z - (z*z - x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
		i++
	}
	fmt.Printf("Sqrt  -- for x:%v needed %v iterations \n", x, i)

	return s
}

func main() {
	var y float64
	for i:= 1; i<=5; i++ {
		y = float64(i)
		fmt.Println(Sqrt(y))
		fmt.Println(Sqrt2(y))
		fmt.Println(Sqrt1(y))
		fmt.Println(math.Sqrt(y))
	}
}

//https://go-tour-ru-ru.appspot.com/flowcontrol/8