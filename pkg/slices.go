package pkg

import "strings"

type Slices struct{}

func (w *Slices) append(slice []string, index int, context []string) []string {
	var before = make([]string, index)
	var after = make([]string, len(slice)-index)
	copy(before, slice[:index])
	copy(after, slice[index:])
	return append(append(before, context...), after...)
}

//SliceIndex 查询str 在切片中索引，如果不存在返回 -1
func (w *Slices) SliceIndex(slice []string, str string) int {
	for i, line := range slice {
		if strings.Contains(line, str) {
			return i + 1
		}
	}
	return -1
}
