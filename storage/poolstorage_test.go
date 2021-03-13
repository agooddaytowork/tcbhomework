package storage

import (
	"testing"

	"github.com/openlyinc/pointy"
	"tam.io/homework/models"
)

func TestPoolStorageAdd(t *testing.T) {

	poolID := int64(12345)
	pool := models.PoolObject{
		PoolID:     &poolID,
		PoolValues: []int32{1, 2, 3, 4},
	}

	want := "inserted"
	got := Pools.Add(pool)

	if want != got {
		t.Fatalf("got %s, want %s", got, want)
	}

	want = "appended"
	got = Pools.Add(pool)

	if want != got {
		t.Fatalf("got %s, want %s", got, want)
	}

	queryRequest := models.PoolQueryRequest{
		Percentile: pointy.Float64(10.0),
		PoolID:     pointy.Int64(12345),
	}

	got1, err := Pools.Query(queryRequest)
	want1 := int32(1)
	if err != nil {
		t.Error(err)
	}

	if got1 != want1 {
		t.Fatalf("got %d, want %d", got1, want1)
	}

}
