package restapi

import (
	"net/http"

	"github.com/go-openapi/loads"
	"tam.io/homework/restapi/operations"
)

func getAPI() (*operations.PoolserviceAPI, error) {
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}
	api := operations.NewPoolserviceAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}
	h := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}
