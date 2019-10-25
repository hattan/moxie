package accumulate

func Accumulate(input []string, fn func(string) string) []string {
	size := len(input)
	result := make([]string, len(input))
	for i := 0; i < size; i++ {
		result[i] = fn(input[i])
	}
	return result
}
