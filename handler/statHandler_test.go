package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/jcasey214/hashit/stats"
)

func TestGetStats(t *testing.T) {
	t.Run("returns Server Stats", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/stats", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetStats)

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		stats := stats.ServerStats{}

		body, _ := ioutil.ReadAll(rr.Result().Body)
		fmt.Println(string(body))

		err := json.Unmarshal(body, &stats)

		if err != nil {
			t.Error("could not deserialize response")
		}

	})
}
