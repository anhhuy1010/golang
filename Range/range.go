package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128} // tạo một mảng cắt có các giá trị như

func main() {
	for i, v := range pow { // khởi tạo vòng lắp for  có 2 biến số và range  bao gồm i sẽ duyệt qua có bao nhiêu phần tử, v sẽ duyệt qua các phần tử có trong mảng
		fmt.Printf("2**%d = %d\n", i, v) // hiển thị kết quả
	}

	poww := make([]int, 10)
	for i := range pow {
		poww[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
