package ogsclient_test

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type templateData struct {
	ServerUrl string
}

func fakeOgsServer(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	templateData := templateData{
		ServerUrl: server.URL,
	}
	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		serveJsonTemplate(w, filepath.Join("testdata", "user_resource.json"), templateData)
	})
	mux.HandleFunc("/me/games", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		page := r.URL.Query().Get("page")
		if page == "" {
			page = "1"
		}
		serveJsonTemplate(w, filepath.Join("testdata", fmt.Sprintf("simplified_games_resource_page_%s.json", page)), templateData)
	})
	return server
}

func serveJsonTemplate(w http.ResponseWriter, jsonResponseTemplateFilename string, data templateData) {
	template, err := template.ParseFiles(jsonResponseTemplateFilename)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	template.Execute(w, data)
}
