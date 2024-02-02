package utils

func Filter[T any](slice []*T, f func(*T) bool) []*T {
	var output []*T
	for _, e := range slice {
		if f(e) {
			output = append(output, e)
		}
	}
	return output
}
