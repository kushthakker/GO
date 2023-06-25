package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

type shape interface {
	getArea() (float64, error)
}

func main() {
	sa := square {
		sideLength: 10,
	}
	ta := triangle {
		height: 10,
		base: 6,
	}
	printArea(sa)
	printArea(ta)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() (float64,error) {
	fmt.Println(s, "p")
	return s.sideLength * s.sideLength, nil
}
func (t triangle) getArea() (float64,error) {
	fmt.Println(t, "t")
	return 0.5 * t.base * t.height, nil
}
