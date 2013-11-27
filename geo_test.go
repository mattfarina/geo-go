package geo

// import (
// 	"testing"
// )

const MAX_DIFF float64 = 0.0000001

func diff(a, b float64) bool {
	if a > b {
		return (a - b) > MAX_DIFF
	}
	return (b - a) > MAX_DIFF
}
