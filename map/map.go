package main

import (
	"fmt"
)

type a struct {
	lat, long float64
}

var m map[string]a

func main() {
	m = make(map[string]a)
	m["bela lab"] = a{
		40.345654, -76.45533,
	}

	var m = map[string]a{
		"bell labs": a{
			40.34345, -122.33445,
		},
		"google": a{
			36.34223, -122.080922,
		},
	}
	fmt.Println(m)

	// một cách code khác
	//type Vertex struct {
	//	Lat, Long float64
	//}

	//var m = map[string]Vertex{
	//	"Bell Labs": {40.68433, -74.39967},
	//	"Google":    {37.42202, -122.08408},
	//}

	//func main() {
	//	fmt.Println(m)
	//}
	/////////////////////////////

	n := make(map[string]int)
	n["answer"] = 42
	fmt.Println("the value: ", n["answer"])
	n["answer"] = 48
	fmt.Println("the value: ", n["answer"])
	delete(m, "Answer")
	fmt.Println("the value: ", m["answer"])

	v, ok := n["answer"]
	fmt.Println("the value: ", v, "present?", ok)
}
