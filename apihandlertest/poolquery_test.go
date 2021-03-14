package apihandlertest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"tam.io/homework/models"
)

const POOL_QUERY_ENDPOINT = "/api/pools/query"

func TestPoolQueryEndpointOK(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	insertRequest := []byte(`{"poolId":2222, "poolValues":[1,2,3,4,5]}`)
	queryRequest := []byte(`{"poolId":2222, "percentile":99.5}`)

	_, err := http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(insertRequest))

	if err != nil {
		t.Error(err)
		return
	}

	_, err = http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(insertRequest))

	if err != nil {
		t.Error(err)
		return
	}

	res, err := http.Post(ts.URL+POOL_QUERY_ENDPOINT, "application/json", bytes.NewBuffer(queryRequest))

	if err != nil {
		t.Error(err)
		return
	}
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		t.Errorf("want status 200, got %d , response msg: %s", res.StatusCode, bodyBytes)
	}

	var resObj models.PoolQueryResponse

	err = json.Unmarshal(bodyBytes, &resObj)

	if err != nil {
		t.Error(err)
	}

	if *resObj.CalculatedQuantile != 5 {
		t.Errorf("Want CalculatedQuantile 5, got %d", *resObj.CalculatedQuantile)
	}

	if *resObj.TotalCount != 10 {
		t.Errorf("Want TotalCount 5, got %d", *resObj.TotalCount)
	}

}

func TestPoolQueryEndpointPoolNotFound(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	request := []byte(`{"poolId":11111, "percentile":99.5}`)
	res, err := http.Post(ts.URL+POOL_QUERY_ENDPOINT, "application/json", bytes.NewBuffer(request))

	if err != nil {
		t.Error(err)
		return
	}

	if res.StatusCode != http.StatusNotFound {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		t.Errorf("want status 404, got %d , response msg: %s", res.StatusCode, bodyBytes)
	}

}

func TestPoolQueryEndpointMissingRequiredFields(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	testcases := []struct {
		name               string
		request            []byte
		expectedStatusCode int
	}{
		{"missing ID", []byte(`{"percentile":10.0}`), http.StatusUnprocessableEntity},
		{"missing Percentile", []byte(`{"poolId":12345}`), http.StatusUnprocessableEntity},
		{"missing both", []byte(`{}`), http.StatusUnprocessableEntity},
		{"unidentify name", []byte(`{"somename":1234}`), http.StatusUnprocessableEntity},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			res, err := http.Post(ts.URL+POOL_QUERY_ENDPOINT, "application/json", bytes.NewBuffer(tc.request))

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

// RUN 1K concurrent query test
func TestPoolQueryEndpointConcurrent(t *testing.T) {

	ts := getTestServer()
	defer ts.Close()

	insertRequest := []byte(`{"poolId":3333, "poolValues":[1,2,3,4,5]}`)
	queryRequest := []byte(`{"poolId":3333, "percentile":99.5}`)

	_, err := http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(insertRequest))

	if err != nil {
		t.Error(err)
		return
	}

	_, err = http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(insertRequest))

	if err != nil {
		t.Error(err)
		return
	}

	t.Run("group", func(t *testing.T) {
		for i := 0; i < 1000; i++ {

			testName := "tc " + fmt.Sprintf("%d", i)

			t.Run(testName, func(t *testing.T) {

				t.Parallel() // enable concurrent test

				res, err := http.Post(ts.URL+POOL_QUERY_ENDPOINT, "application/json", bytes.NewBuffer(queryRequest))

				if err != nil {
					t.Error(err)
					return
				}
				bodyBytes, _ := ioutil.ReadAll(res.Body)
				if res.StatusCode != http.StatusOK {
					t.Errorf("want status 200, got %d , response msg: %s", res.StatusCode, bodyBytes)
				}

				var resObj models.PoolQueryResponse

				err = json.Unmarshal(bodyBytes, &resObj)

				if err != nil {
					t.Error(err)
				}

				if *resObj.CalculatedQuantile != 5 {
					t.Errorf("Want CalculatedQuantile 5, got %d", *resObj.CalculatedQuantile)
				}

				if *resObj.TotalCount != 10 {
					t.Errorf("Want TotalCount 5, got %d", *resObj.TotalCount)
				}

			})
		}
	})
}
