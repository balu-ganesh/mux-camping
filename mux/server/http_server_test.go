package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/camping/entities"
)

func TestGetSlot(t *testing.T) {

	testCases := map[string]struct {
		params     map[string]string
		statusCode int
	}{
		"goodParams": {
			map[string]string{
				"startDate": "2022-08-25 12:00",
				"endDate":   "2022-08-25 12:00",
			},
			http.StatusOK,
		},
		"startDateEmpty": {
			map[string]string{
				"startDate": "2022-08-25 12:00",
				"endDate":   "2022-08-25 12:00",
			},
			http.StatusOK,
		},
	}

	for _, tp := range testCases {
		url := fmt.Sprintf("/api/slots")

		req, _ := http.NewRequest("GET", url, nil)
		q := req.URL.Query()

		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		res, err := MakeTestRequest(req)
		if err != nil {
			t.Error("error getting response")
		}
		if res.StatusCode != tp.statusCode {
			t.Errorf("expected status code %d. received %d. test failed. ", tp.statusCode, res.StatusCode)
		}
	}

}

type TestRoundTripper struct {
	res *http.Response
	err error
}

func (rt *TestRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	result := []entities.AvailableSlots{
		{SlotID: 1002, SiteID: 501},
		{SlotID: 1003, SiteID: 501},
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(result); err != nil {
		return nil, err
	}
	rt.res = &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(buf)}
	return rt.res, rt.err
}

func MakeTestRequest(r *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	server := New()
	server.ServeHTTP(w, r)
	if w.Result().StatusCode != 200 {
		return nil, fmt.Errorf("error: HTTP status code %v", w.Result().StatusCode)
	}
	return w.Result(), nil
}
