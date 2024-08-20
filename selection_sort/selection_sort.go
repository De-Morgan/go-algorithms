package selectionsort

/*
Selection Sort
The Selection Sort algorithm finds the lowest value in an array and moves it to the front of the array.
*/

func selectionSort(nums []int) {

	n := len(nums)

	for i := range n {
		minIndex := i
		for j := i; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j

			}
		}
		if i == minIndex {
			continue
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]

	}

}
