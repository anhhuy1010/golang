package main

import "fmt"

func main() {
	var a [2]string         // khai báo một mảng loại string có độ dài là 2 bắt đầu từ 0 kết thúc 1
	a[0] = "Hello"          // gán cho mảng đầu tiên giá trị
	a[1] = "World"          // gán cho mảng thứ hai giá trị
	fmt.Println(a[0], a[1]) // hiển thị kết quả của hai mảng sau khi được gán
	fmt.Println(a)          // hiển thị kết quả mảng a

	primes := [6]int{2, 3, 5, 7, 11, 13} // khai báo một mảng loại int có độ dài là 6 bắt đầu từ 0 kết thúc là 5
	fmt.Println(primes)                  //hiển thị kết quả của mảng vừa khai báo

	//có thể tùy chọn hiển thị những vị trí mà mình mong muốn trong mảng
	var s []int = primes[1:5] // tạo ra một slice tên là s cắt mảng primes từ vị trí số 1 đến vị trí số 5
	fmt.Println(s)

	names := [4]string{ // khai báo một mảng
		"hi",
		"hii",
		"hiii",
		"hiiii",
	}
	fmt.Println(names) // hiển thị mảng vừa khai báo
	d := names[0:3]    // slice d có độ dài 0:3 của mảng trên
	b := names[2:3]    //slice b có độ dài 2:3 của mảng trên
	fmt.Println(d)     // hiển thị kết quả
	fmt.Println(b)     // hiển thị kết quả
	d[0] = "iiiii"     // thay đổi giá trị tại vị trí d[0]
	fmt.Println(d, b)  // hiển thị kết quả
	fmt.Println(names) // hiển thị kết quả
	// khi cắt không nhất thiết phải chọn khoảng cách từ đâu đến đâu có thể chọn bất kỳ vị trí nào mà mình mong muốn
	q := []int{2, 3, 5, 7, 11, 13} // tạo mảng cắt kiểu int
	fmt.Println(q)
	q = q[1:4]     // mảng cắt từ khoảng 1:4
	fmt.Println(q) // hiển thị dữ liệu sau khi cắt

	q = q[:2] // hiển thị dữ liệu bắt đầu từ 0 đến dưới vị trí số 2
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)

	var w []int                    // tạo một mảng cắt không có dữ liệu
	fmt.Println(s, len(s), cap(s)) // in độ dài của mảng cắt và dung lượng của mảng cắt
	if w == nil {                  // nếu mảng cắt không có kết quả trả về
		fmt.Println("nil!!!") // hiển thị thông báo
	}
	x := make([]int, 5)
	printSlice("x", x)
	y := make([]int, 0, 5)
	printSlice("y", y)
	z := y[:2]
	printSlice("z", z)
	c := z[2:5]
	printSlice("c", c)

	// có thể thêm mảng cắt bằng hàm append
	// var s []int
	// s = append(s,1,2,3)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
