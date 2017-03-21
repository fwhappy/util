package util

// 创建一个切片，从0 ~ length
func GenRange(length int) []int {
	_range := make([]int, length, length)
	for i := 0; i < length; i++ {
		_range[i] = i
	}
	return _range
}