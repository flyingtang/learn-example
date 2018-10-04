package sort

// 冒泡排序算法改进
func BubbleSort(s []int) []int{
	length := len(s)

	for i := 0; i < length - 1; i++ {
		status := false

		for j := i + 1; j < length; j++ {
			if s[i] > s[j] {
				status = true
				s[i], s[j] = s[j], s[i]
			}

		}
		if status == false {
			break
		}

	}
	return s
}
