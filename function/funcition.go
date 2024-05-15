package main

import (
	"fmt"
	"math"
)

// 1
func adder() func(int) int { // hàm adder trả về một hàm khác có biến sum cục bộ
	sum := 0                 // cho sum cục bộ bằng 0
	return func(x int) int { // trả về hàm func ban đầu và sum sẽ thành sum += x
		sum += x
		return sum // trả về kết quả sum hiện tại
	}
}

func compute(fn func(float64, float64) float64) float64 { // khởi tạo hàm  với hai biến số thuộc kiểu fl64, kết quả trả về là kiểu fl64
	return fn(3, 4) //trả về kết quả
}
func main() {
	hypot := func(x, y float64) float64 { // khởi tạo hàm cho hai biến số x, t là loại fl64, trả về kết quả fl64
		return math.Sqrt(x*x + y*y) // kết quả trả về là căn bậc hai của (x^2 +y^2)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	//1
	pos, neg := adder(), adder() // khởi tạo hai hàm adder()
	for i := 0; i < 10; i++ {    // cho chạy vòng lặp for chạy từ 0-9
		fmt.Println( // hiển thị kết quả pos và neg
			pos(i),
			neg(-2*i),
		)
	}
}
