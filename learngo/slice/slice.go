package slice

import "fmt"

//切片

func updateSlice(s []int) {

}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s:=arr[2:6]
	fmt.Println("arry[2:6]=", arr[2:6])
	fmt.Println("arry[:6]=", arr[:6])
	fmt.Println("arry[2:]=", arr[2:])
	fmt.Println("arry[:]=", arr[:])
}
