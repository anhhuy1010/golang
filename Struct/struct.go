package main

import "fmt"

type Person struct { // định nghĩa một struct có tên là person gồm các thông tin sau Name, Age, city
	Name string
	Age  int
	city string
}

//1
var (
	v1 = Person{Name: "khongcoten", Age: 0, city: ""}
	v2 = Person{}
	p  = &Person{}
)

func main() {
	person1 := Person{ // khởi tạo đối tượng person1
		Name: "Anh Huy",
		Age:  21,
		city: "Ho Chi Minh",
	}
	person2 := Person{ // khởi tạo đối tượng person2
		Name: "Huy Anh",
		Age:  12,
		city: "Ha Noi",
	}

	//hiển thị thông tin của person1
	fmt.Println("Name: ", person1.Name)
	fmt.Println("Age: ", person1.Age)
	fmt.Println("city: ", person1.city)
	// hiển thị thông tin của person2
	fmt.Println("Name: ", person2.Name)
	fmt.Println("Age: ", person2.Age)
	fmt.Println("city: ", person2.city)

	person1.Name = "Trần Bảo Anh Huy" // thay đổi thông tin của của trường Nam
	fmt.Println("Name: ", person1.Name)
	fmt.Println("Age: ", person1.Age)
	fmt.Println("city: ", person1.city)
	// sử dụng con trỏ trong struct
	a := person2                // gán biến a là person2
	fmt.Println(a)              // hiển thị thông tin của biến a sau khi được gán
	p := &a                     // gán địa chỉ biến a vào con trỏ p
	fmt.Println(*p)             // hiển thị thông tin của con trỏ p sau khi được gán
	p.Name = "Trần Thị Tuyết A" //thay đổi thông tin của biến Name trong a bằng con trỏ p
	fmt.Println(a)              // hiển thị kết quả của a sau khi thay đổi thông tin
	// sử dụng hàm var trong struct kèm con trỏ
	//1
	fmt.Println(v1, p, v2)

}
