package insertionsort

// Insertion Sort
// The Insertion Sort algorithm uses one part of the array to hold the sorted values,
// and the other part of the array to hold values that are not sorted yet.
func insertionSort(nums []int) {
	for i := range len(nums) {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}

	}
}
