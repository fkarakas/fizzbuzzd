// +build blackbox

// Run blackbox test with a running fizzbuzzd

package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

var (
	host = "fizzbuzzd:8080"
)

func waitForFizzbuzzd(t *testing.T, host string) {

	for i := 0; i < 5; i++ {
		statusCode := ping(t, host)

		if statusCode == http.StatusOK {
			return
		}

		log.Printf("Waiting for fizzbuzzd '%s' ...", host)

		time.Sleep(3 * time.Second)
	}

	assert.Fail(t, fmt.Sprintf("fizzbuzzd '%s' not listening !", host))
}

func ping(t *testing.T, host string) int {
	endpoint := "http://" + host + "/"

	statusCode, _ := HttpRequest(t, http.MethodGet, endpoint)

	return statusCode
}

type GetArrayResponse struct {
	Result []string `json:"result"`
}

func getArray(t *testing.T, host string) int {
	endpoint := "http://" + host + "/api/v1/fizzbuzz/numbers/2/3/terms/bob/dylan?limit=234"

	statusCode, body := HttpRequest(t, http.MethodGet, endpoint)

	assert.Equal(t, http.StatusOK, statusCode)

	res := GetArrayResponse{}
	err := json.Unmarshal(body, &res)
	assert.Nil(t, err)

	assert.Equal(t, 234, len(res.Result))
	assert.Equal(t, "5", res.Result[4])
	assert.Equal(t, "bob", res.Result[3])
	assert.Equal(t, "dylan", res.Result[8])
	assert.Equal(t, "bobdylan", res.Result[11])

	return statusCode
}

func TestBlackBox(t *testing.T) {
	waitForFizzbuzzd(t, host)
	getArray(t, host)
}
