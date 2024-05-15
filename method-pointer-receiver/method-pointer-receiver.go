package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v *vertex) Scale(f float64) { // gán con trỏ v có loại là vertex vào scale()
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *vertex) Abs() float64 { // gán con trỏ v vào vertex và Abs()
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &vertex{3, 4}                                         // truyền tham số vào
	fmt.Printf("Trước khi scale: %+v, Abs: %v\n ", v, v.Abs()) //%+v trong printf dùng để in ra các giá trị của biến đồng thời in ra tên các trường trong biến struct
	v.Scale(5)                                                 // truyền tham số vào
	fmt.Printf("Sau khi scale: %+v, Abs: %v\n", v, v.Abs())    //%vtrong printf dùng để in ra giá trị của một biến theo cách phù hipwj với kiểu dữ liệu của biến đó
}
