package bubblesort

/*
Bubble Sort
Bubble Sort is an algorithm that sorts an array from the lowest value to the highest value.
*/

func bubbleSort(nums []int) {
	n := len(nums)
	var swaped bool
	for i := range n {
		swaped = false
		for j := range n - i - 1 {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swaped = true
			}
		}
		if !swaped {
			break
		}
	}
}
