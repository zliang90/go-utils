package sort

func BubbleSort(nums []int, reversed bool) []int {
	lenNums := len(nums)

	if lenNums < 2 {
		return nums
	}

	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if reversed {
				// 倒序
				if nums[i] < nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			} else { // 正序
				if nums[i] > nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}

	return nums
}
