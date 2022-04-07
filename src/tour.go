package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// pic[i][x] = uint8(x^i)
			pic[i][x] = uint8(x*i)
			// pic[i][x] = uint8((x+i)/2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
