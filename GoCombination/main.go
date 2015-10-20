package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	reds := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "22", "23", "24", "25", "26", "28", "29", "30", "31", "32", "33"}
	result := Combinaion(reds, 6)

	//print
	fmt.Println("length:\r", len(result))
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i] + "\r")
	}
}

func Combinaion(arr []string, m int) []string {

	if m == 1 {
		return arr
	}

	result := make([]string, 0)

	if len(arr) == m {
		var str string
		for i := 0; i < len(arr); i++ {
			str = str + arr[i]
			if i != len(arr)-1 {
				str = str + ","
			}
		}
		result = append(result, str)
		return result
	}

	firstItem := arr[0]

	tempArr1 := Combinaion(append(arr[1:]), m-1)
	for i := 0; i < len(tempArr1); i++ {
		result = append(result, firstItem+","+tempArr1[i])
	}

	tempArr2 := Combinaion(append(arr[1:]), m)

	result = append(result, tempArr2...)

	return result
}
