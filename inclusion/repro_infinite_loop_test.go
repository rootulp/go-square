package inclusion_test

import (
	"math"
	"testing"

	"github.com/celestiaorg/go-square/v3/inclusion"
)

// TestRoundUpPowerOfTwoInfiniteLoop reproduces the infinite loop bug described
// in https://github.com/celestiaorg/go-square/issues/117. The loop-based
// RoundUpPowerOfTwo shifts `result` left until it reaches `input`. When input
// is math.MaxInt, the shift overflows the signed int to negative, then to zero.
// Since 0 < math.MaxInt is always true, the loop never terminates.
func TestRoundUpPowerOfTwoInfiniteLoop(t *testing.T) {
	inclusion.RoundUpPowerOfTwo(math.MaxInt)
}
