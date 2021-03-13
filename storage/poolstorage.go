package storage

import (
	"sync"

	"tam.io/homework/models"
)

type PoolStorage struct {
	Mutex sync.RWMutex
	Pools map[int64]models.PoolObject
}

var Pools PoolStorage = PoolStorage{Pools: make(map[int64]models.PoolObject)}

func (store *PoolStorage) Add(inputPool models.PoolObject) string {

	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	currentPool, found := store.Pools[*inputPool.PoolID]

	if found {
		currentPool.PoolValues = append(currentPool.PoolValues, inputPool.PoolValues...)
		store.Pools[*currentPool.PoolID] = currentPool

		return "appended"
	}

	store.Pools[*inputPool.PoolID] = inputPool
	return "inserted"

}

func (store *PoolStorage) Query(request models.PoolQueryRequest) (float64, error) {

	// result, err := util.CalculateQuantileNearestRank(r)

}
