package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestItemHandler_GetItem verifies that the GetItem handler returns the expected JSON response.
func TestItemHandler_GetItem(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/item/42", nil)
	// Set a fake path variable manually (requires Go 1.22+)
	req.SetPathValue("id", "42")

	rec := httptest.NewRecorder()

	handler := ItemHandler{}
	handler.GetItem(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

	body, _ := io.ReadAll(res.Body)

	var data map[string]string
	err := json.Unmarshal(body, &data)
	assert.NoError(t, err)
	assert.Equal(t, "Item #42 retrieved successfully!", data["msg"])
}

// TestWithCORS_GET ensures the CORS headers are correctly set for normal GET requests.
func TestWithCORS_GET(t *testing.T) {
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := withCORS(inner)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "*", res.Header.Get("Access-Control-Allow-Origin"))
	assert.Contains(t, res.Header.Get("Access-Control-Allow-Methods"), "GET")
	assert.Contains(t, res.Header.Get("Access-Control-Allow-Headers"), "Content-Type")
}

// TestWithCORS_OPTIONS ensures that preflight OPTIONS requests are handled correctly.
func TestWithCORS_OPTIONS(t *testing.T) {
	// This handler should never be called for OPTIONS
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called for OPTIONS requests")
	})

	handler := withCORS(inner)

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNoContent, res.StatusCode)
	assert.Equal(t, "*", res.Header.Get("Access-Control-Allow-Origin"))
	assert.Contains(t, res.Header.Get("Access-Control-Allow-Methods"), "OPTIONS")
}
