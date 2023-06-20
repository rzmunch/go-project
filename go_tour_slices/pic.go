package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	//Slice of lenght dy
	//Made by slices of lengt dx
	//sdy := make([]uint8, dy)
	//sdx := make([]uint8, dx)
	picss := make([][]uint8, dy)
	for y := range picss {
		//picss[i] = (make([]uint8, dx))
		//Add elements
		for x := 0; x < dx; x++ {
			picss[y] = append(picss[y], uint8(x*y)) //Append adds to the slice
		}
		fmt.Printf("%v\n", picss[y])

	}
	//s = append(s, 0)
	//fmt.Printf("%d\n", sdy)
	//fmt.Printf("%d\n", sdx)
	return picss
}

func main() {
	pic.Show(Pic)
}
