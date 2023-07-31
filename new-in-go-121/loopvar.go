package main_test

import "testing"

// START OMIT
func TestAllEvenBuggy(t *testing.T) {
	testCases := []int{1, 2, 4, 6}
	for _, v := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			if v&1 != 0 {
				t.Fatal("odd v", v)
			}
		})
	}
}

// END OMIT
