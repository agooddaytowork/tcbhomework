package util

import "testing"

func TestCalculateQuantileNearestRank(t *testing.T) {

	f1 := []int32{15, 20, 30, 40, 50}
	f2 := []int32{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
	f3 := []int32{20, 15, 40, 30, 50}

	testcases := []struct {
		name            string
		inputPool       []int32
		inputPercentile float64
		expected        int32
	}{
		{"tc1", f1, 0, 15},
		{"tc2", f1, 5, 15},
		{"tc3", f1, 30, 20},
		{"tc4", f1, 40, 20},
		{"tc1", f3, 0, 15},
		{"tc2", f3, 5, 15},
		{"tc3", f3, 30, 20},
		{"tc4", f3, 40, 20},
		{"tc5", f2, 25, 7},
		{"tc6", f2, 50, 8},
	}

	for _, tc := range testcases {

		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			t.Parallel()
			got, _ := CalculateQuantileNearestRank(tc.inputPool, tc.inputPercentile)

			if got != tc.expected {
				t.Errorf("got %d, want %d", got, tc.expected)
			}
		})
	}

}
