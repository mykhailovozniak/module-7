package main

import (
	"io/ioutil"
	"log"
	"module-7/pkg/models/mock"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type testServer struct {
	*httptest.Server
}

func newTestApplication(t *testing.T) *application {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &application{materials: &mock.MaterialModel{}, infoLog: infoLog, errorLog: errorLog}
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}
