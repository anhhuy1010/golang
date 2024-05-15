package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

// 1
func add(x int, y int) int { // hàm thêm hai kiểu số
	return x + y
}

// một cách gọi khác khi thêm hai đối số
// func add(x, y int) int{}
// 2
func swap(x, y, z string) (string, string, string) { // hàm hoán đổi
	return z, x, y // trả về theo thứ tự, kết quả hiển thị sẽ phụ thuộc vào kết quả trả về này
}

// 3
func split(sum int) (x, y int) {
	x = sum * 4 / 2
	y = sum - x
	return // kết quả trả về không có biến số gọi là trả về trống
}

// 4
var python, java bool // hàm var khai báo một danh sách các đối số, loại (var {đối số} {loại})
var i int

// ngoài ra khai báo biến var khác
// var i, j int = 1, 2 								-> thêm biến số trực tiếp vào
//var c, python, java = true, false, "no"			-> thêm biến số trực tiếp vào

// 5
var (
	bien1 bool       = false
	bien2 uint64     = 1<<64 - 1            // uint64 là kiểu dữ liệu chứa giá trị từ 0 đến 2^64 -1
	z     complex128 = cmplx.Sqrt(-5 + 12i) // complex128 là kiểu dữ liệu cho số phức với phần thực và phần ảo 		complex64 tương tự như complex64 nhưng độ chính xác thấp hơn nhưng nó chiếm ít dung lượng hơn
)

// ngoài ra khi khai báo biến var nhưng không đưa giá trị ban đầu vào kết quả vẫn sẽ hiển thị
// 7
const Pi = 3.14           //hằng số được khai báo như biến nhưng có thêm từ 'const'
const world = "blablabla" // hằng số không thể khai báo được bằng cách 	:=

// 10
func sqrt(x float64) string { //khởi tạo một căn bậc hai là số nguyên
	if x < 0 { // dùng hàm if
		return sqrt(-x) + "i" // trả về kết quả nếu đúng điều kiện
	}
	return fmt.Sprint(math.Sqrt(x)) // trả về kết quả nếu điều kiện if không đúng
}

// 10.1
func pow(x, n, lim float64) float64 { // khởi tạo một hàm với 3 giá trị loại số thực
	if v := math.Pow(x, n); v < lim { // dùng hàm if nếu như lũy thừa bé hơn giới hạn lim thì trả về kết quả v
		return v
	}
	return lim // nếu không thỏa mãn điều kiện if trả về giới hạn lim
}

// 11
func poww(x, n, lim float64) float64 { // khởi tạo một hàm với 3 giá trị loại số thực
	if v := math.Pow(x, n); v < lim { // dùng hàm if nếu lũy thừa bé hơn giới hạn thì trả về kết quả v
		return v
	} else { // hàm ngược lại
		fmt.Printf("%g >= %g\n", v, lim) // trả ra kết quả %g là hiển thị số dưới dạng ngắn gọn những vẫn được độ chính xác cao, %g\n trả về kết quả như %g nhưng xuống dòng, trả về kết quả của lim và v
	}
	return lim // nếu hai điều kiện trên sai thì trả về kết quả lim
}

func main() {
	fmt.Println("số yêu thích của tôi là: ", rand.Intn(10))
	// rand là hàm random
	fmt.Println("căn bậc hai của một số là:", math.Sqrt(7))
	// math.sqrt là hàm căn bậc hai
	fmt.Println("số pi là:", math.Pi)
	// math.pi là hàm số pi
	//1
	fmt.Println("x + y =", add(42, 13))
	// hàm tính toán sau khi thêm hai đối số
	//2
	a, b, c := swap("hahaha", "hehehe", "hihihihi")
	fmt.Println(a, b, c)
	// hiển thị sau khi hoán đổi
	//3
	fmt.Println(split(17)) // kết quả hiện thị sau khi trả về hàm return trống
	//4
	fmt.Println(i, python, java) // kết quả hiển thị của hàm var
	//5
	fmt.Printf("Type: %T Value: %v\n", bien1, bien1) // %T là định dạng sử dụng trong hàm Printf. được sử dụng để in ra loại dữ liệu của biến (int,bool,comlex128,unit64...)
	fmt.Printf("Type: %T Value: %v\n", bien2, bien2) // %v\n là định dạng sử dụng trong hàm Printf. được sử dụng để in ra giá trị của biến
	fmt.Printf("Type: %T Value: %v\n", z, z)
	//6
	var i float32                        // gán biến i là loại số phức
	ki := i                              // cho biến ki bằng biến i
	fmt.Printf("ki la là loại %T\n", ki) //-> loại của ki giống loại của biến i
	//7
	fmt.Println("người ta nói", world) // kết quả hiển thị sau khi khai báo một hằng số
	fmt.Println("số Pi bằng:", Pi)
	//8
	sum := 2                  // khai báo giá trị cho biến đầu tiên
	for i := 0; i < 10; i++ { // cho chạy vòng lặp for, cho i bằng không và i bé hơn 10 thì i tăng
		sum += i // giá trị sum cộng liên tục với giá trị chạy trong hàm for
	}
	fmt.Println(sum) //hiển thị kết quả
	//8.1
	summ := 1         // gán giá trị ban đầu cho biến
	for summ < 1000 { // dùng vòng lặp for cho chạy từ 1 đến 999
		summ += summ // giá trị được cộng dồn đến hết
	}
	fmt.Println(summ) // hiển thị kết quả
	//9
	summm := 1
	for summm < 1000 { // dùng vòng lặp while cho chạy từ 1 đến 999		điểm khác nhạ=au giữa hai vòng lặp là for có điều kiện ban đầu, while thì không, trong code thì khi dùng for sẽ có dấu ; để có thể ngăn cách giữa các điều kiện, còn while vẫn dùng hàm for nhưng mà không có dấu ; để ngăn cách các điều kiện.
		summm += summm // giá trị được cộng dồn đến hết
	}
	fmt.Println(summm)
	//10
	fmt.Println(sqrt(2), sqrt(-4)) // kết quả
	//10.1
	fmt.Println(
		pow(3, 2, 10), // trường hợp thỏa điều kiện trả về kết quả v
		pow(3, 3, 20)) //trường hợp không thỏa điều kiện trả về kết quả lim

	//11
	fmt.Println(
		poww(3, 2, 10),
		poww(3, 3, 20))
	//12
	fmt.Print("go runs on ")
	switch os := runtime.GOOS; os { // hàm switch là câu lệnh điều khiển sử dụng để kiểm tra các giá trị của biến hoặc biểu thức. bằng hệ điều hành mà máy đang sử dụng ( runtime.GOOS là hàm kiểm tra máy tính đang sử dụng hệ điều hành gì)
	case "darwin": // nếu runtime.GOOS trả về darwin thì sẽ hiển thị kết quả OS X
		fmt.Println("OS X")
	case "linux": // nếu runtime.GOOS trả về linux thì sẽ hiển thị kết quả là linux
		fmt.Print("Linux")
	default: // trường hợp mặc định sẽ hiển thị khi kết quả trả về không có như hai trường hợp trên
		fmt.Printf("%s.\n", os)
	}
	//12.1
	fmt.Println("when saturday?")
	today := time.Now().Weekday() // lấy thời gian hiện tại trong tuần
	switch time.Saturday {        // gọi hàm switch để kiểm tra các giá trị của biến hoặc biểu thức
	case today + 0: // trường hợp ngày hôm nay thì hiển thị kết quả
		fmt.Println("today")
	case today + 1: // trường hợp là ngày hôm nay cộng thêm 1 thì hiển thị kết quả
		fmt.Println("Tomorrow")
	case today + 2: //trường hợp là ngày hôm nay cộng thêm 2 thì hiển thị ra kết quả
		fmt.Println("In two days")
	default: // trường hợp nếu là ngày mà ngày hôm nay cộng thêm một số lớn hơn 2 thì hiển thị kết quả
		fmt.Println("To far away")
	}
	//12.2
	t := time.Now() // gọi hàm thời gian thực tế hiện tại
	switch {        // gọi hàm switch để kiểm tra các mốc thời gian
	case t.Hour() < 12: //nếu thời gian hiện tại bé hơn 12 thì hiển thị kết quả
		fmt.Println("chào buổi sáng")
	case t.Hour() < 17: // nếu thời gian hiện tại lớn hơn 12 và nhỏ hơn 17 thì hiển thị kết quả
		fmt.Println("chào buổi chiều")
	default: //nếu thời gian hiện tại không thỏa mãn hai trường hợp trên thì hiển thị kết quả
		fmt.Println("chào buổi tối")
	}
	//13
	defer fmt.Println("world") // hàm defer là hàm trì hoãn, hàm này sẽ được thực thi khi những hàm khác đã trả ra kết quả
	fmt.Println("Helloooo nhaaaaa")
	//nếu có nhiều hàm defer trong một gói thì sau khi những hàm khác đã trả ra kết quả những hàm defer sẽ lần lượt trả ra kết quả

	//14

}
