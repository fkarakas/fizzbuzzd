package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	r *gin.Engine
)

func setup() {
	r = NewRouter(gin.TestMode, "9.9.9")
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func serveHTTP(t *testing.T, r *gin.Engine, method string, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	assert.Nil(t, err)

	r.ServeHTTP(w, req)

	return w
}

type GetArrayResponse struct {
	Result []string `json:"result"`
}

type PingResponse struct {
	Version string `json:"version"`
}

// Check ping responding with the version number
func TestPing(t *testing.T) {
	w := serveHTTP(t, r, http.MethodGet, "/")
	assert.Equal(t, http.StatusOK, w.Code)

	res := PingResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err)

	assert.Equal(t, "9.9.9", res.Version)
}

func TestGetArrayOK(t *testing.T) {
	w := serveHTTP(t, r, http.MethodGet, "/api/v1/fizzbuzz/numbers/2/3/terms/bob/dylan")
	assert.Equal(t, http.StatusOK, w.Code)

	res := GetArrayResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err)

	assert.Equal(t, 100, len(res.Result))
	assert.Equal(t, "5", res.Result[4])
	assert.Equal(t, "bob", res.Result[3])
	assert.Equal(t, "dylan", res.Result[8])
	assert.Equal(t, "bobdylan", res.Result[11])
}

// Test the limit parameter
func TestGetArrayLimit(t *testing.T) {
	w := serveHTTP(t, r, http.MethodGet, "/api/v1/fizzbuzz/numbers/2/3/terms/bob/dylan?limit=123")
	assert.Equal(t, http.StatusOK, w.Code)

	res := GetArrayResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err)

	assert.Equal(t, 123, len(res.Result))
}

// Some bad request tests
func TestGetArrayBadRequest(t *testing.T) {
	w := serveHTTP(t, r, http.MethodGet, "/api/v1/fizzbuzz/numbers/2/0/terms/bob/dylan?limit=123")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetArrayNotFound(t *testing.T) {
	w := serveHTTP(t, r, http.MethodGet, "/api/v1/fizzbuzz/numbers/2/terms/bob/")
	assert.Equal(t, http.StatusNotFound, w.Code)
}
