package handlers

import (
	"net/http/httptest"
	"testing"
)

/*
func TestHandlers(t *testing.T) {
	mux := http.NewServeMux()
	SetupHandlers(mux)

	testSrv := httptest.NewServer(mux)
	defer testSrv.Close()

	resp, err := http.Get(testSrv.URL + "/healthcheck")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respBody := string(data)
	t.Log(respBody)
}
*/

func TestHealthCheckHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/healthcheck", nil)
	defer req.Body.Close()
	healthCheckHandler(w, req)

	byteBuf := w.Body
	expectedResponse := "ok"
	if byteBuf.String() != "ok" {
		t.Fatalf("Expected: %s, Got: %s\n", expectedResponse, byteBuf.String())
	}
}

func TestIndexHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	defer req.Body.Close()
	indexHandler(w, req)

	byteBuf := w.Body
	expectedResponse := indexHtml
	if byteBuf.String() != expectedResponse {
		t.Fatalf("Expected: %s, Got: %s\n", expectedResponse, byteBuf.String())
	}
}
