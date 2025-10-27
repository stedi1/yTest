// list.go
package list

func Min(list []int) int {
	var min int
	if len(list) < 1 {
		return 0
	}
	min = list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}
