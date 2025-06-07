package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

var python, js bool

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe              = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

type Coord struct {
	x int
	Y int
}

func main() {
	defer fmt.Println("============================exec")
	fmt.Println("random number: ", rand.Intn(100))
	fmt.Println(math.Pi)
	fmt.Println(add(7, 202020))
	a, b := swap("wow", "lol")
	fmt.Println(a, b)
	fmt.Println(split(12345))
	var c int
	fmt.Println(c, python, js)

	var i, j = 0, "huh"
	fmt.Println(i, j)
	d := 0
	fmt.Println(d)

	fmt.Printf("type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T value: %v\n", z, z)
	var ii int
	var ff float64
	var bb bool
	var ss string
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)
	var x, y int = 4, 6
	f := math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
	fmt.Println(Big * 0.1)

	sum := 0
	for i := range 10 {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum < 2000 {
		sum *= 2
	}
	fmt.Println(sum)
	if true {
		fmt.Println(true)
	}

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("mac")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("other os")
	}

	defer fmt.Println("---------------last")

	s := [6]int{
		2, 6, 3, 0, 8,
	}

	slice := s[0:3]
	slice_slice := slice[0:1]
	slice_slice[0] = 9
	fmt.Println(s)
}
