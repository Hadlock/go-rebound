package status_api

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHomeRedirects(t *testing.T) {	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/index.html", nil)

	handler, pattern := Handlers().Handler(req)
	handler.ServeHTTP(w, req)

	// Root returns a redirect for directory indexes
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Home page didn't return %v, returned %v", http.StatusMovedPermanently, w.Code)
	}

	// Handler pattern should be "/"
	if pattern != "/" {
		t.Errorf("Pattern handler should be \"\\\", was %s", pattern)
	}
}
