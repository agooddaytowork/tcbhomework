package storage

import (
	"testing"

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
}
