package arrays

// Algorithm: Find The Lowest Value in an Array

// lowestValue returns the lowest integer value from a slice of int
// returns 0, false if [value] is empty
func lowestValue(value []int) (int, bool) {

	if len(value) == 0 {
		return 0, false
	}
	minVal := value[0]

	for _, v := range value {
		if v < minVal {
			minVal = v
		}
	}

	return minVal, true

}
