package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"strconv"
)

func TestCreateHash(t *testing.T) {
	t.Run("handles POST requests to /hash", func(t *testing.T) {
		data := url.Values{}
		data.Set("password", "angryMonkey")
		req, _ := http.NewRequest("POST", "/hash", strings.NewReader(data.Encode()))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHash)

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("returns a 405 for non-POST requests to /hash", func(t *testing.T) {
		data := url.Values{}
		data.Set("password", "angryMonkey")
		req, _ := http.NewRequest("GET", "/hash", strings.NewReader(data.Encode()))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHash)

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}
	})

	t.Run("returns a 400 when password field not provided", func(t *testing.T) {
		data := url.Values{}
		req, _ := http.NewRequest("POST", "/hash", strings.NewReader(data.Encode()))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHash)

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
}

func TestGetHashById(t *testing.T) {
	t.Run("returns a 404 for non-existant id", func(t *testing.T) {
		data := url.Values{}
		req, _ := http.NewRequest("GET", "/hash/99", strings.NewReader(data.Encode()))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetHashById)

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

	t.Run("returns a 400 for non int id", func(t *testing.T) {
		data := url.Values{}
		req, _ := http.NewRequest("GET", "/hash/browns", strings.NewReader(data.Encode()))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetHashById)

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
}