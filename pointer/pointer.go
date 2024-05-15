package main

import "fmt"

func main() {
	i, j := 42, 2701 // khai báo biến i có giá trị bằng 42, j có giá trị bằng 2701

	p := &i         // gán địa chỉ biến i vào con trỏ p
	fmt.Println(*p) //in ra giá trị của con trỏ sau khi được gán		-> 42
	*p = 21         // thay đổi giá trị của con trỏ p
	fmt.Println(i)  // in ra giá trị của i			-> 21
	fmt.Println(&p) // in ra địa chỉ của con trỏ
	p = &j
	*p = *p / 37
	fmt.Println(j)
	// &p dùng để lấy địa chỉ của con trỏ
	//*p dùng để đọc hoặc thao tác với giá trị mà con trỏ đang trỏ đến.
}
