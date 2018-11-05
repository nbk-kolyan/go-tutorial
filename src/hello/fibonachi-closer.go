package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	y := 1
	// this values dont die after function call - its like static vars
	return func () int {
		fmt.Println(x , y)
		x, y = y, x + y   //x := y then y := y + old value of x (before this row)
		fmt.Println(x , y)
		return x
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//https://go-tour-ru-ru.appspot.com/moretypes/26