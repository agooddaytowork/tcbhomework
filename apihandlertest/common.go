package apihandlertest

import (
	"log"
	"net/http/httptest"

	"tam.io/homework/restapi"
)

func getTestServer() *httptest.Server {
	handler, err := restapi.GetAPIHandler()
	if err != nil {
		log.Fatal(err)
	}
	return httptest.NewServer(handler)
}
