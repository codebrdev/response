package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	JSON(rec, http.StatusCreated, map[string]string{"a": "b"})

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Errorf("content-type = %q", ct)
	}
	if rec.Header().Get("encoding") != "" {
		t.Error("não deve enviar o header inválido 'encoding'")
	}
	if rec.Body.String() != "{\"a\":\"b\"}\n" {
		t.Errorf("body = %q", rec.Body.String())
	}
}
