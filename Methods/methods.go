package main

import (
	"fmt"
	"math"
)

type Vertexx struct { // khởi tạo struct với 2 tham số ban đầu là X, Y
	X, Y float64
}

// 1
type Myfloat float64

func (f Myfloat) Abs() float64 { // tạo ta một kiểu dữ liệu mới dựa trên float64 và định nghĩa phương thức Abs()
	if f < 0 { // nếu f < 0 thì trả về giá trị tuyệt đối cảu f
		return float64(-f)
	}
	return float64(f) // trả về f
}

// 2
func (v *Vertexx) Scale(f float64) { // đưa con trỏ vào phương thức vertexx và khởi tạo hàm scale f kiểu float64
	v.X = v.X * f // lấy tọa độ của x nhân lên cho f
	v.Y = v.Y * f // lấy tọa độ của y nhân lên cho f
}

// 3
func Scale(v *Vertexx, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertexx) Abs() float64 { // khởi tạo phương thức vertexx với tham số v, hàm Abs() dùng để tính toán và trả về độ lớn của vector 2D
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// cách sử dụng khác
func Abs(v Vertexx) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertexx{3, 4} // khởi tạo v phương thức Vertexx truyền hai tham số là 3, 4
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	//1
	f := Myfloat(-math.Sqrt2) // gán f bằng kiểu dữ liệu Myfloat với -math.sqrt2 là giá trị âm của căn bậc 2
	fmt.Println(f.Abs())
	//2
	v.Scale(10) // truyền tham số vào hàm scale đã được định nghĩa ở trên
	fmt.Println(v.Abs())
	//3
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
