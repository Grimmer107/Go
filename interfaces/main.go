package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interfaces are implemented implicitly
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare this.
func (t T) M() {
	fmt.Println(t.S)
}

func (f MyFloat) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v is of age %v", p.Name, p.Age)
}

func do(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("This is an integer: %d", t)
	case string:
		fmt.Printf("This is a string: %s", t)
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println()
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	// Interfaces are implemented implicitly
	var i I = T{"hello"}
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()

	// empty interface
	var ei interface{}
	describeEmpty(ei)

	ei = 42
	describeEmpty(ei)

	ei = "hello"
	describeEmpty(ei)

	// type assertions
	var ti interface{} = "hello"

	s := ti.(string)
	fmt.Println(s)

	s, ok := ti.(string)
	fmt.Println(s, ok)

	tf, ok := ti.(float64)
	fmt.Println(tf, ok)

	// f = ti.(float64) panic

	// type switches
	do(21)
	do("hello, world")
	do(true)

	// Stringer
	fmt.Println(Person{"Usama", 24})
}
