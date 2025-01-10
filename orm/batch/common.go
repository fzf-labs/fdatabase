package batch

import "math"

// getSQLQuantity 计算需要生成的 SQL 语句数量
func getSQLQuantity(length, batchSize int) int {
	return int(math.Ceil(float64(length) / float64(batchSize)))
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
