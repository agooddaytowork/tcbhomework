package apihandler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"tam.io/homework/models"
	"tam.io/homework/restapi/operations/pool"
	"tam.io/homework/storage"
)

// func HandlePoolAddRequest()

type insertPoolHandler struct{}

func (handler *insertPoolHandler) Handle(params pool.InsertPoolParams) middleware.Responder {

	result := storage.Pools.Add(*params.Body)
	responseMsg := models.PoolObjectAddResponse{Status: &result}
	return pool.NewInsertPoolOK().WithPayload(&responseMsg)

}

func NewInsertPoolHandler() pool.InsertPoolHandler {
	return &insertPoolHandler{}
}

type queryPoolHander struct{}

func (handler *queryPoolHander) Handle(params pool.QuerryPoolParams) middleware.Responder {

	calculatedQuantile, totalCount, err := storage.Pools.Query(*params.Body)

	if err != nil {

		return pool.NewQuerryPoolDefault(http.StatusNotFound).WithPayload(models.ErrorResponse(err.Error()))
	}

	successRespMsg := models.PoolQueryResponse{CalculatedQuantile: &calculatedQuantile, TotalCount: &totalCount}

	return pool.NewQuerryPoolOK().WithPayload(&successRespMsg)

}

func NewQueryPoolHandler() pool.QuerryPoolHandler {
	return &queryPoolHander{}
}
