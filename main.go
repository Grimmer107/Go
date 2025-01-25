package main // package declaration

// factored imports
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
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
}
