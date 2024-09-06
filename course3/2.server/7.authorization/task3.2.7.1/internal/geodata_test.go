package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCase struct {
	add  ResponseAddress
	want string
}

func TestGetQuery(t *testing.T) {
	testCases := []TestCase{
		{add: ResponseAddress{
			Address: Address{
				Road:   "San Francisco",
				County: "some County",
			},
		},
			want: "San+Francisco"},
	}
	for _, testCase := range testCases {
		res := GetQuery(testCase.add)
		if res != testCase.want {
			t.Errorf("GetQuery(%s): got %s, want %s", testCase.add, res, testCase.want)
		}
	}
}

func TestHandlerGeocode(t *testing.T) {
	{
		address := Address{
			Road:    "red square",
			Town:    "Moscow",
			State:   "",
			Country: "",
		}

		reqBody, _ := json.Marshal(ResponseAddress{Address: address})
		req := httptest.NewRequest(http.MethodPost, "/geocode", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()

		handlerGeocode(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		result, _ := io.ReadAll(w.Body)

		expected := "Coordinate: Lon:37.62148935239179, Lat:55.7535926"
		if string(result) != expected {
			t.Errorf("expected %v, got %v", expected, string(result))
		}
	}

	{
		req := httptest.NewRequest(http.MethodPost, "/geocode", strings.NewReader("wrJson"))
		w := httptest.NewRecorder()

		handlerGeocode(w, req)
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	}
}

func TestHandlerSearch(t *testing.T) {
	{
		reqAddr := RequestAddressGeocode{Lat: 37.6, Lng: 55.8}

		reqBody, _ := json.Marshal(reqAddr)
		req := httptest.NewRequest(http.MethodPost, "/search", strings.NewReader(string(reqBody)))
		w := httptest.NewRecorder()

		handlerSearch(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	}
	{
		req := httptest.NewRequest(http.MethodPost, "/search", strings.NewReader("wrJson"))
		w := httptest.NewRecorder()

		handlerSearch(w, req)
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	}
}
