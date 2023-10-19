package leetcoode

// sort
func hIndex(citations []int) int {

	quickSort(citations)

	var h, i = 0, len(citations) - 1
	for i >= 0 && citations[i] > h {
		h++
		i--
	}
	return h
}
