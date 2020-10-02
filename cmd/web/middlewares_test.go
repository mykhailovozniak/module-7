package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorrelationId(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/hello", nil)

	if err != nil {
		t.Fatal(err)
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	correlationId(next).ServeHTTP(rr, r)

	rs := rr.Result()

	if rs.StatusCode != http.StatusOK {
		t.Errorf("status code should be 200")
	}
}
