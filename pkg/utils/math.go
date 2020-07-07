package utils

func AverageUint64(xs []uint64) float64 {
	var total uint64 = 0
	for _, v := range xs {
		total += v
	}
	return float64(total) / float64(len(xs))
}
