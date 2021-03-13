package apihandlertest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

const POOL_ADD_ENDPOINT = "/api/pools/add"

func TestPoolAddEndpointOK(t *testing.T) {

	ts := getTestServer()
	defer ts.Close()

	insertRequest := []byte(`{
		"poolId":12345,
		"poolValues":[]
	}`)

	res, err := http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(insertRequest))

	if err != nil {

		t.Error(err)
		return
	}

	if res.StatusCode != http.StatusOK {

		bodyBytes, _ := ioutil.ReadAll(res.Body)
		t.Errorf("want status 200, got %d , resp msg: %s", res.StatusCode, bodyBytes)

	}
}

func TestPoolAddEndPointMissingRequiredFields(t *testing.T) {

	ts := getTestServer()
	defer ts.Close()

	testcases := []struct {
		name               string
		request            []byte
		expectedStatusCode int
	}{
		{"missing ID", []byte(`{"poolValues":[]}`), http.StatusUnprocessableEntity},
		{"missing PoolValue", []byte(`{"poolId":12345}`), http.StatusUnprocessableEntity},
		{"missing both", []byte(`{}`), http.StatusUnprocessableEntity},
		{"unidentify name", []byte(`{"somename":1234}`), http.StatusUnprocessableEntity},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			res, err := http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(tc.request))

			if err != nil {
				t.Error(err)
				return
			}
			if res.StatusCode != tc.expectedStatusCode {
				bodyBytes, _ := ioutil.ReadAll(res.Body)
				t.Errorf("want status 422, got %d , response msg: %s", res.StatusCode, bodyBytes)
			}

		})
	}

}
