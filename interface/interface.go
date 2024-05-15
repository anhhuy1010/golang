package main

import (
	"fmt"
)

type hcn interface {
	chuvi() float64
	dientich() float64
}
type dulieu struct {
	width float64
	heigh float64
}

func (r dulieu) chuvi() float64 {
	return (r.width + r.heigh) * 2
}

func (r dulieu) dientich() float64 {
	return r.width * r.heigh
}

//1

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	s := dulieu{2, 3}
	fmt.Println("chu vi của hình chữ nhật: ", s.chuvi())
	fmt.Println("diện tích hình chữ nhật: ", s.dientich())

	//1
	var i I = T{"hello"}
	i.M()
}
