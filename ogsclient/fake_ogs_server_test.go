package ogsclient_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeOgsServer(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		serveJson(w, filepath.Join("testdata", "user_resource.json"))
	})
	mux.HandleFunc("/api/v1/megames", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		serveJson(w, filepath.Join("testdata", "simplified_games_resource_page_1.json"))
	})
	return httptest.NewServer(mux)
}

func serveJson(w http.ResponseWriter, jsonResponeFilePath string) {
	response, err := ioutil.ReadFile(jsonResponeFilePath)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
