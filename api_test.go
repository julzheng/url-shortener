package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	. "url-shortener/internal/types"
)

func TestShortenGenerationRoute(t *testing.T) {
	router := SetupRouter()
	ja := jsonassert.New(t)

	urlValues := url.Values{}
	urlValues.Set("url", "http://google.com")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/shorten", strings.NewReader(urlValues.Encode()))
	router.ServeHTTP(w, req)

	ja.Assertf(w.Body.String(), `{"shortened_url": "<<PRESENCE>>"}`)
}

func TestRedirectRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()

	urlValues := url.Values{}
	urlValues.Set("url", "http://google.com")
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/shorten", strings.NewReader(urlValues.Encode()))

	router.ServeHTTP(w, req)

	var shortenResp ShortenResp
	err := json.NewDecoder(w.Body).Decode(&shortenResp)
	if err != nil {
		fmt.Println("err: ", err)
	}

	redirectReq, _ := http.NewRequest(http.MethodGet, shortenResp.ShortenedURL, nil)

	router.ServeHTTP(w, redirectReq)

	assert.Equal(t, http.StatusOK, w.Code)
}
