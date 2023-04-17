package main

import "fmt"

var lst = [3]int{10, 20, 30}

func demo(arr *[3]int) {
	//(*arr)[0], (*arr)[1], (*arr)[2] = 1, 2, 3
	arr[0], arr[1], arr[2] = 1, 2, 3
	fmt.Println(*arr)
}

func main() {
	demo(&lst)       // [1 2 3]
	fmt.Println(lst) // [1 2 3]
}
