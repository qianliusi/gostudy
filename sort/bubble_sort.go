package sort

import "fmt"

func bubbleSort(d []int) {
	length := len(d)
	flag := true
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if d[j] > d[j+1] {
				flag = false
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
		if flag {
			fmt.Println(i)
			break
		}
	}
}
