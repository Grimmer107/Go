package main

import "fmt"

func indexOf[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// generic type
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

	// type parameters
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(indexOf(si, 3))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(indexOf(ss, "hello"))

	// genric type
	list := List[int]{nil, 1}
	fmt.Println(list)
}
