package main

import ("fmt"
		"math/cmplx")

var c, python, java bool
var i, j int = 1, 2
var(
	a bool = false
	b uint64 = 1<<64 -1
	d complex128 = cmplx.Sqrt( - 5 +12i)
	)
const Pi = 3.14
func main() {
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, "\n",python, java)
	k := 3
	fmt.Println(k)
	fmt.Printf("Type- %T  Value- %v \n", a,a)
	fmt.Printf("Type- %T  Value- %v \n", b,b)
	fmt.Printf("Type- %T  Value- %v \n", d,d)
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Printf("Type- %T  Value- %v \n", u,u)
	fmt.Println(Pi)
}
/*bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128*/
