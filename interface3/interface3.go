package main

import "fmt"

func main() {
	var i interface{} = "hello" // khai báo một biến interface gán giá trị "hello"
	s := i.(string)             // khai báo biến s bằng biến i
	fmt.Println(s)              // in ra biến s

	s, ok := i.(string) // khai báo biến s và biến ok
	fmt.Println(s, ok)  // in ra biến s và ok

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
