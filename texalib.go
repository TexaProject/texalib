package texalib

import "strconv"

// Convert is responsible for converting and expoting the string array received from the /texa webpage
func Convert(xs []string) []uint64 {
	ArtiQSA := make([]uint64, 0, len(xs))

	for x := range xs {
		if s, err := strconv.ParseUint(xs[x], 10, 64); err == nil {
			// fmt.Printf("%T, %v\n", s, s)
			ArtiQSA = append(ArtiQSA, s)
		}
	}
	// fmt.Printf("%T %v", ArtiQSA, ArtiQSA)
	return ArtiQSA
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

// setHumanQSA creates and returns the Quantum Score Array for Human Intelligence
func setHumanQSA(xs []uint64) []uint64 {
	HumanQSA := make([]uint64, 0, len(xs))
	for _, x := range xs {
		if x == 1 {
			HumanQSA = append(HumanQSA, 0)
		}
		if x == 0 {
			HumanQSA = append(HumanQSA, 1)
		}
	}
	return HumanQSA
}
