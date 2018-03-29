package simplify

import "strconv"

func parseFloats(items []string) []float64 {
	result := make([]float64, len(items))
	for i, item := range items {
		f, err := strconv.ParseFloat(item, 64)
		if err != nil {
			panic(err)
		}
		result[i] = f
	}
	return result
}
