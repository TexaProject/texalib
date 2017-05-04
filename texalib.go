package texalib

import "strconv"

// Convert is responsible for converting and expoting the string array received from the /texa webpage
func Convert(xs []string) []uint64 {
	b := make([]uint64, 0, len(xs))

	for x := range xs {
		if s, err := strconv.ParseUint(xs[x], 10, 64); err == nil {
			// fmt.Printf("%T, %v\n", s, s)
			b = append(b, s)
		}
	}
	// fmt.Printf("%T %v", b, b)
	return b
}

// Total calculates the sum total of the given array
func Total(xs []uint64) uint64 {
	var total uint64
	total = 0
	for _, x := range xs {
		total += x
	}
	return total
}

// Average calculates the average of the given array
func Average(xs []uint64) uint64 {
	var total uint64
	total = 0
	for _, x := range xs {
		total += x
	}
	return total / uint64(len(xs))
}
