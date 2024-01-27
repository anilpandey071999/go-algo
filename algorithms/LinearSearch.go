package algorithms

func LinearSearch(slice []int, target int) bool {
	for _, ele := range slice {
		if ele == target {
			return true
		}
	}
	return false
}
