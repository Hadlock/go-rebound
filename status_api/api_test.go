package status_api

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func beforeTest(method string, url string) (*httptest.ResponseRecorder, *http.Request, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)

	return w, req, err
}


// I can't figure out how to do content testing;
// although the server works, when I test it I only get
// 301s and 404s. Have to try to test this indirectly.
// Should ask somebody.
func TestHomeUsesCorrectPattern(t *testing.T) {
	_, req, _ := beforeTest("GET", "/index.html")

	_, pattern := Handlers().Handler(req)

	// Handler pattern should be "/"
	if pattern != "/" {
		t.Errorf("Pattern handler should be \"\\\", was %s", pattern)
	}
}

func TestHomeRedirects(t *testing.T) {	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/index.html", nil)

	handler, _ := Handlers().Handler(req)
	handler.ServeHTTP(w, req)

	// Root returns a redirect for directory indexes
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Home page didn't return %v, returned %v", http.StatusMovedPermanently, w.Code)
	}	
}
