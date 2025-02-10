package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method defined for type Vertex with value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method defined for type Vertex with pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2.0)
	fmt.Println(v.Abs())

	// Methods and pointer indirection
	// method with pointer receiiver taking either value or pointer
	v.Scale(2)
	fmt.Println(v.Abs())

	pr := &v
	pr.Scale(2)
	fmt.Println(pr.Abs())

	// method with value receiiver taking either value or pointer
	fmt.Println(v.Abs())
	fmt.Println(pr.Abs())
}
