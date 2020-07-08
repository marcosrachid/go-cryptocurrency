package utils

func AverageInt64(xs []int64) float64 {
	var total int64 = 0
	for _, v := range xs {
		total += v
	}
	return float64(total) / float64(len(xs))
}
