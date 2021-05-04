package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPreviousPrimeNumberController(t *testing.T) {
	router := routerSetup()
	w := httptest.NewRecorder()

	// set JSON body
	reqParam := `{"limit":55}`
	resParam := `{"highestPrime":53}`

	// Mock HTTP Request and it's return
	req, err := http.NewRequest("POST", "/get-previous-prime", strings.NewReader(reqParam))

	// make sure request was executed
	assert.NoError(t, err) // Serve Request and recorded data
	router.ServeHTTP(w, req)

	// Test results
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resParam, w.Body.String())
}