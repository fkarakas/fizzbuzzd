// Utils for server testing
package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// Wrap http call, returns http code and response body
func HttpRequest(t *testing.T, httpMethod string, url string) (int, []byte) {

	req, err := http.NewRequest(httpMethod, url, nil)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{
		Timeout: 20 * time.Second,
	}

	response, err := httpClient.Do(req)
	assert.Nil(t, err)

	var statusCode int
	var bytes []byte

	if response != nil {
		defer response.Body.Close()

		statusCode = response.StatusCode

		bytes, err = ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
	}

	return statusCode, bytes
}
