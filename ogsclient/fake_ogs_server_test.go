package ogsclient_test

import (
	"fmt"
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
		page := r.URL.Query().Get("page")
		if page == "" {
			page = "1"
		}
		serveJson(w, filepath.Join("testdata", fmt.Sprintf("simplified_games_resource_page_%s.json", page)))
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
