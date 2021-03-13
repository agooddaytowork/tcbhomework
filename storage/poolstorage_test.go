package storage

import (
	"fmt"
	"testing"

	"github.com/openlyinc/pointy"
	"tam.io/homework/models"
)

func TestPoolStorageInsertThenAppend(t *testing.T) {

	pool := models.PoolObject{
		PoolID:     pointy.Int64(12345),
		PoolValues: []int32{1, 2, 3, 4},
	}
	// inserted once
	want := "inserted"
	got := Pools.Add(pool)

	if want != got {
		t.Fatalf("got %s, want %s", got, want)
	}

	// appended once
	want = "appended"
	got = Pools.Add(pool)

	if want != got {
		t.Fatalf("got %s, want %s", got, want)
	}

	// appended twice
	got = Pools.Add(pool)

	if want != got {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPoolStorageInsertConcurent(t *testing.T) {

	for i := 0; i < 1000; i++ {
		testName := fmt.Sprintf("tc %d", i)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			pool := models.PoolObject{
				PoolID:     pointy.Int64(12345),
				PoolValues: []int32{1, 2, 3, 4},
			}
			// appended once
			want := "appended"
			got := Pools.Add(pool)

			if want != got {
				t.Fatalf("got %s, want %s", got, want)
			}
			// appended twice
			got = Pools.Add(pool)

			if want != got {
				t.Fatalf("got %s, want %s", got, want)
			}
		})
	}
}

func TestPoolStorageQueryPoolExist(t *testing.T) {

	pool := models.PoolObject{
		PoolID:     pointy.Int64(1111),
		PoolValues: []int32{15, 20, 35, 40, 50},
	}

	Pools.Add(pool)

	queryRequest := models.PoolQueryRequest{
		Percentile: pointy.Float64(5.0),
		PoolID:     pointy.Int64(1111),
	}
	wantResult := int32(15)
	wantTotalCount := int64(5)

	gotResult, gotTotalCount, err := Pools.Query(queryRequest)

	if err != nil {
		t.Error(err)
	}

	if gotResult != wantResult {

		t.Errorf("Result error: want %d , got %d", wantResult, gotResult)
	}

	if gotTotalCount != wantTotalCount {
		t.Errorf("Total Count error: want %d , got %d", wantTotalCount, gotTotalCount)

	}

}

func TestPoolStorageQueryPoolExistConcurent(t *testing.T) {

	queryRequest := models.PoolQueryRequest{
		Percentile: pointy.Float64(5.0),
		PoolID:     pointy.Int64(1111),
	}
	wantResult := int32(15)
	wantTotalCount := int64(5)

	for i := 0; i < 1000; i++ {
		testName := fmt.Sprintf("tc %d", i)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			gotResult, gotTotalCount, err := Pools.Query(queryRequest)

			if err != nil {
				t.Error(err)
			}

			if gotResult != wantResult {

				t.Errorf("Result error: want %d , got %d", wantResult, gotResult)
			}

			if gotTotalCount != wantTotalCount {
				t.Errorf("Total Count error: want %d , got %d", wantTotalCount, gotTotalCount)

			}
		})
	}
}

func TestPoolStorageQueryPoolNotFound(t *testing.T) {
	queryRequest := models.PoolQueryRequest{
		Percentile: pointy.Float64(5.0),
		PoolID:     pointy.Int64(2222),
	}

	_, _, err := Pools.Query(queryRequest)

	if err == nil {
		t.Errorf("Expected pool not found error")
	}
}
