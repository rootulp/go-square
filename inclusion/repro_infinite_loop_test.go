package inclusion_test

import (
	"math"
	"testing"
	"time"

	"github.com/celestiaorg/go-square/v3/inclusion"
)

// TestRoundUpPowerOfTwoInfiniteLoop reproduces the infinite loop bug described
// in https://github.com/celestiaorg/go-square/issues/117. The loop-based
// RoundUpPowerOfTwo shifts `result` left until it reaches `input`. When input
// is math.MaxInt, the shift overflows the signed int to negative, then to zero.
// Since 0 < math.MaxInt is always true, the loop never terminates.
func TestRoundUpPowerOfTwoInfiniteLoop(t *testing.T) {
	done := make(chan struct{})
	go func() {
		inclusion.RoundUpPowerOfTwo(math.MaxInt)
		close(done)
	}()
	select {
	case <-done:
		t.Fatal("expected infinite loop but function returned")
	case <-time.After(3 * time.Second):
		t.Log("confirmed: RoundUpPowerOfTwo(math.MaxInt) hangs forever due to integer overflow in the left-shift loop")
	}
}
