package main

import "fmt"

func main() {
	var arr = [7]int{5, 1, 33, 23, 1, -5, 1}

	var i int

	//var last = arr[len(arr)-1]
	//fmt.Println(last)
	var k = 0
	var temp = 0
	var j = 0
	for i = 0; i < len(arr); i++ {
		k = i + 1

		for j = k; j < len(arr)-1; j++ {
			if k < len(arr)-1 && arr[k] == 1 {

				k = k + 1
			}
		}

		if arr[i] == 1 && i != len(arr)-1 {
			//fmt.Println(arr[i])

			temp = arr[k]

			arr[k] = arr[i] // 2=2,5=5,45=12
			arr[i] = temp
			fmt.Println(arr)

		}

	}

	//fmt.Println(arr)
}
