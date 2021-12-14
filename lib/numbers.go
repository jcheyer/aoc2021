package lib

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func HighUInt64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func HighInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LowInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func LowUInt64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func AbsDiffUInt64(a, b uint64) uint64 {
	if a > b {
		return a - b
	}

	return b - a
}
