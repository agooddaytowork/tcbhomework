package util

import "testing"

// Test cases got from https://en.wikipedia.org/wiki/Percentile#The_nearest-rank_method
func TestCalculateQuantileNearestRankSortAndUnsorted(t *testing.T) {

	sorted := []int32{15, 20, 35, 40, 50}
	unsorted := []int32{50, 40, 20, 35, 15}

	testcases := []struct {
		name            string
		inputPool       []int32
		inputPercentile float64
		expected        int32
	}{
		{"tc1", sorted, 5, 15},
		{"tc2", sorted, 30, 20},
		{"tc3", sorted, 40, 20},
		{"tc4", sorted, 50, 35},
		{"tc5", sorted, 100, 50},
		{"tc6", unsorted, 30, 20},
		{"tc7", unsorted, 50, 35},
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

func TestCalculateQuantileNearestRankEmptyPoolInput(t *testing.T) {

	_, err := CalculateQuantileNearestRank([]int32{}, 20)

	if err == nil {
		t.Errorf("Expected empty input error")
	}

	if err.Error() != "Input Pool must not be empty" {
		t.Errorf("Expected empty input error")
	}

}

func TestCalculateQuantileNearestRankOutBoundPercentile(t *testing.T) {
	sorted := []int32{15, 20, 35, 40, 50}

	want := "Percentile must be between 0 - 100"
	testcases := []struct {
		name            string
		inputPool       []int32
		inputPercentile float64
		expected        string
	}{
		{"tc1", sorted, -5, want},
		{"tc2", sorted, 101.11, want},
	}

	for _, tc := range testcases {

		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			t.Parallel()
			_, err := CalculateQuantileNearestRank(tc.inputPool, tc.inputPercentile)

			if err == nil || err.Error() != want {
				t.Errorf("Expected outbound error")
			}
		})
	}
}
