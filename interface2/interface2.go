package main

import "fmt"

func main() {
	var i interface{} // khai báo biến i kiểu interface
	describe(i)       // trả về kết quá của biến i

	i = 42 // gán cho i một giá trị bằng 42
	describe(i)

	i = "hello" // gán cho i một giá trị là ""hello"
	describe(i)
}

func describe(i interface{}) { // khởi tạo hàm desctibe
	fmt.Printf("(%v, %T)\n", i, i) // hiểu thị kết quả
}
