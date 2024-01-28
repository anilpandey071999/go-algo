package algorithms

import ("sort")
func BinarySearch(slices []int, target int) bool{
	sort.Ints(slices)
	lo := 0
	hi := len(slices) -1
// 1,2,3,4,5

	for lo < hi {
		mid := ( lo + hi ) / 2
		if slices[mid] == target{
			return true
		}else if slices[mid] < target{
			lo = mid + 1
		}else{
			hi = mid
		}
	}
	return false
}