package main // package declaration

// factored imports
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// function syntax
func add(x int, y int) int {
	return x + y
}

// function with multiple return items
func swap(x, y string) (string, string) {
	return y, x
}

// function with naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// struct
type Vertex struct {
	X int
	Y int
}

// constant
const Pi = 3.14

// variable declaration
var c, java bool
var a, b int = 1, 2

// main() starting point of execution
func main() {
	var m, n int = 1, 2 // defining multiple variables
	k := 3              // declaration and initialization
	fmt.Println(c, java, a, b, m, n, k)

	fmt.Println("Hello, 世界")                        // utf-8 encoding scheme
	fmt.Println(math.Sqrt(64))                      // math package
	fmt.Println("Random number is:", rand.Intn(10)) // generate random integer
	fmt.Println("Sum:", add(7, 1))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(10))

	// factored variable declaration
	var (
		check  bool       = false                // bool type
		MaxInt uint64     = 1<<64 - 1            // integer type
		z      complex128 = cmplx.Sqrt(-5 + 12i) // complex type
		str    string     = "hello"              // string type
	)

	// string formatting %T for type, %v for value, %s for string, %q for quoted string
	fmt.Printf("Type %T Value %v\n", check, check)
	fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value %v\n", z, z)
	fmt.Printf("Type %T Value %q\n", str, str)

	// type conversion
	var i float64 = 298.0
	var j int = int(i)
	fmt.Printf("Type %T Value %v\n", j, j)

	// type inference
	f := 3.142
	fmt.Printf("Type %T Value %v\n", f, f)

	// constants
	const isChecked = true
	fmt.Printf("Type %T Value %v\n", isChecked, isChecked)
	fmt.Printf("Type %T Value %v\n", Pi, Pi)

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	// The init and post statements are optional and can be dropped
	i = 0
	for i < 10 {
		fmt.Print(i)
		i += 1
	}
	fmt.Println()

	// infinite loop
	// for {
	// 	fmt.Print("hello")
	// }

	// if - else
	e := 24
	if e%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// if with short statement
	p := 3.2
	if v := math.Pow(2, p); v > 10 {
		fmt.Println("Above limit")
	} else {
		fmt.Println("Within limit")
	}

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows")
	}

	// switch without condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer
	defer fmt.Println("World")
	fmt.Print("Hello ")

	// pointers
	e, h := 42, 2701

	g := &e         // point to e
	fmt.Println(*g) // read i through the pointer
	*g = 21         // set i through the pointer
	fmt.Println(e)  // see the new value of e

	g = &h         // point to f
	*g = *g / 37   // divide f through the pointer
	fmt.Println(h) // see the new value of f

	// struct
	v1 := Vertex{1, 2}
	fmt.Println(v1, v1.X, v1.Y)

	// pointer to struct
	ps := &v1
	ps.X = 1e9
	fmt.Println(v1)

	// struct literal
	v2 := Vertex{X: 7}
	fmt.Println(v2)

	// Arrays
	var s [2]string
	s[0] = "hello"
	s[1] = "world"
	fmt.Println(s[0], s[1], s)

	// array literal
	var q = [5]int{1, 2, 3, 4, 5}
	fmt.Println(q)

	// slices
	var slice []int = q[1:4]
	fmt.Println(slice)

	// slice literal
	sl := []int{1, 2, 3, 4}
	fmt.Println(sl)

	sl2 := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
	}
	fmt.Println(sl2)

	// slice defaults
	sd := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sd[:2])
	fmt.Println(sd[1:])

	// slice length and capacity
	sc := []int{2, 3, 5, 7, 11, 13}
	printSlice(sc)

	// Slice the slice to give it zero length.
	sc = sc[:0]
	printSlice(sc)

	// Extend its length.
	sc = sc[:4]
	printSlice(sc)

	// Drop its first two values.
	sc = sc[2:]
	printSlice(sc)

	// Nil slices
	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil!")
	}

	// creating slice with make
	csm := make([]int, 5)
	fmt.Println(csm, cap(csm), len(csm))

	csm = make([]int, 5, 8)
	fmt.Println(csm, cap(csm), len(csm))

	// slices with slices
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][0] = "x"
	board[1][1] = "x"
	board[2][2] = "x"

	for row := 0; row < 3; row++ {
		// for col := 0; col < 3; col++ {
		// 	fmt.Printf("%s ", board[row][col])
		// }
		// fmt.Println()
		fmt.Printf("%s\n", strings.Join(board[row], " "))
	}

	// appending to a slice
	var sa []int
	sa = append(sa, 1)
	fmt.Println(sa)

	sa = append(sa, 2, 3)
	fmt.Println(sa)

	// range
	sr := []int{0, 1, 4, 9, 16}
	for index, value := range sr {
		fmt.Printf("2^%d = %d\n", index, value)
	}

	// maps
	var mp map[string]Vertex
	// mp["me"] = Vertex{1, 2}  assigning keys to nil map
	fmt.Println(mp)

	mp = make(map[string]Vertex)
	mp["center"] = Vertex{2, 2}
	fmt.Println(mp)

	// map literals
	mp = map[string]Vertex{
		"Bell Labs": {
			40, -74,
		},
		"Google": {
			3, -1,
		},
	}
	fmt.Println(mp)

	// mutating maps
	mm := make(map[string]int)

	mm["Answer"] = 42
	fmt.Println("The value:", mm["Answer"])

	mm["Answer"] = 48
	fmt.Println("The value:", mm["Answer"])

	delete(mm, "Answer")
	fmt.Println("The value:", mm["Answer"])

	v, ok := mm["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
