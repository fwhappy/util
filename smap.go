package util

import (
	"sync"
)

// SmapLen 获取sync.Map的元素个数
func SmapLen(smap *sync.Map) int {
	length := 0
	smap.Range(func(k, v interface{}) bool {
		length++
		return false
	})
	return length
}
