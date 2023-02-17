package kata

func ArrayDiff(a, b []int) []int {
	for _, bItem := range b {

		j := 0
		for i, aItem := range a {
			if aItem != bItem {
				a[j] = a[i]
				j++
			}
		}
		a = a[:j]
	}
	return a
}
