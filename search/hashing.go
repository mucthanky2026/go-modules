package search

func HashSearch(arr []int, target int) int {
	hashMap := make(map[int]int)
	for i, v := range arr {
		hashMap[v] = i
	}

	if index, found := hashMap[target]; found {
		return index
	}
	return -1
}