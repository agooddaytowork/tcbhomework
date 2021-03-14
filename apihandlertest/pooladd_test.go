package apihandlertest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const POOL_ADD_ENDPOINT = "/api/pools/add"

func TestPoolAddEndpointOK(t *testing.T) {

	ts := getTestServer()
	defer ts.Close()

	insertRequest := []byte(`{"poolId":12345,"poolValues":[1,2,3,4,5]}`)

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

// RUN 1K concurent add requests
func TestPoolAddEndpointConcurent(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	insertRequest := []byte(`{"poolId":123456789,"poolValues":[1,2,3,4,5]}`)

	t.Run("Group", func(t *testing.T) {

		for i := 0; i < 1000; i++ {

			testName := "tc " + fmt.Sprintf("%d", i)

			t.Run(testName, func(t *testing.T) {

				t.Parallel() // enable concurrent

				res, err := http.Post(ts.URL+POOL_ADD_ENDPOINT, "application/json", bytes.NewBuffer(insertRequest))

				if err != nil {
					t.Error(err)
					return
				}
				if res.StatusCode != http.StatusOK {

					bodyBytes, _ := ioutil.ReadAll(res.Body)
					t.Errorf("want status 200, got %d , resp msg: %s", res.StatusCode, bodyBytes)
				}
			})

		}

	})

}
