package storage

import (
	"errors"
	"fmt"
	"sync"

	"tam.io/homework/models"
	"tam.io/homework/util"
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

// Return Calculated Quantile, Total Count Length and Error
func (store *PoolStorage) Query(request models.PoolQueryRequest) (int32, int64, error) {

	store.Mutex.Lock()
	pool, found := store.Pools[*request.PoolID]
	store.Mutex.Unlock()

	if !found {
		return 0, 0, errors.New("Pool with id " + fmt.Sprintf("%d", *request.PoolID) + " not found")
	}
	result, err := util.CalculateQuantileNearestRank(pool.PoolValues, *request.Percentile)

	return result, int64(len(pool.PoolValues)), err
}
