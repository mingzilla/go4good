package main

import (
	"fmt"
	"math/rand"
)

// By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand

func add(x int, y int) int {
	return x + y
}

func swap(x int, y int) (int, int) {
	return y, x
}

func sumNums() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func os(v string) string {
	switch v {
	case "darwin":
		return "Mac"
	case "linux":
		return "Linux"
	default:
		return "Windows"
	}
}

func pt() {
	var i int = 42  // type: int
	var p *int = &i // type: *int
	var x int = *p  // type: int - This is known as "dereferencing" or "indirecting"
	fmt.Println(i, p, x)
}

type Vertex struct {
	X int
	Y int
}

func main() {
	x := rand.Int()
	y := rand.Int()

	v1 := add(x, y)
	v2, v3 := swap(x, y)
	v4 := sumNums()
	v5 := os("win")
	v6 := Vertex{1, 2}

	m := make(map[string]Vertex)
	m["one"] = Vertex{}

	pt()
	fmt.Println(v1, v2, v3, v4, v5)
	fmt.Println(v6, v6.X)
	fmt.Println(m["one"])
}
