package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

const Si = 3.14

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))

	newF, newS := swap("one", "two")
	fmt.Println(newF, newS)

	one, two, tree := split(22)
	fmt.Println(one, two, tree)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	var q, w int = 1, 2
	k := 3
	ci, pythone, javac := true, false, "no!"

	fmt.Println(q, w, k, ci, pythone, javac)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var o int
	var f float64
	var b bool
	var s string
	var n uint64
	var m uintptr
	fmt.Printf("%v %v %v %q %v %v\n", o, f, b, s, n, m)

	var x, y int = 3, 4
	var l float64 = math.Sqrt(float64(x*x + y*y))
	var v uint = uint(l)
	fmt.Println(x, y, v)

	fmt.Printf("v is of type %T\n", v)

	fmt.Println(Si)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	//fmt.Printf("v is of type %T\n", Big) // constant 1267650600228229401496703205376 overflows int
	fmt.Printf("v is of type %v\n", Small)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y, s int) {
	x = sum * 4 / 9
	y = sum - x
	s = x + 7
	return
}
