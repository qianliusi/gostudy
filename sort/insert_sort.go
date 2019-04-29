package sort

func insertSort(d []int) {
	length := len(d)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if d[j+1] < d[j] {
				d[j+1], d[j] = d[j], d[j+1]
			}
		}
	}
}
