package main

import (
	"fmt"
	"math"
)

type Vertex struct { // tạo một kiểu dữ liệu mới với hai biến số là x,y dạng số thực
	X, Y float64
}

func (v *Vertex) Scale(f float64) { // con trỏ v được gắn với kiểu vertex đồng thời hàm scale là một phương thức được gắn với kiểu dữ liệu vertex. -> nó cũng có thể được gọi trebe các thể hiện kiểu vertex thông qua con trỏ
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) { // khởi tạo hàm có tên ScaleFunc với tham số đầu tiên là con trọ đến đối tượng kiểu vertex, tham số thứ hai là f với kiểu float
	v.X = v.X * f
	v.Y = v.Y * f
}

func (q Vertex) Abs() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y)
}

func AbsFunc(w Vertex) float64 {
	return math.Sqrt(w.X*w.X + w.Y*w.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{3, 4}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	q := Vertex{3, 4}
	fmt.Println(q.Abs())
	fmt.Println(AbsFunc(q))

	w := &Vertex{3, 4}
	fmt.Println(q.Abs())
	fmt.Println(AbsFunc(*w))
}
