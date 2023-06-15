package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestMathHandlers(t *testing.T) {
	r := chi.NewRouter()

	r.Get("/min", handleMathRequest(handleMin))
	r.Get("/max", handleMathRequest(handleMax))
	r.Get("/avg", handleMathRequest(handleAverage))
	r.Get("/median", handleMathRequest(handleMedian))
	r.Get("/percentile", handleMathRequest(handlePercentile))

	testCases := []struct {
		Endpoint      string
		QueryParams   map[string]string
		ExpectedCode  int
		ExpectedBody  string
		ExpectedError string
	}{
		{
			Endpoint:     "/min",
			QueryParams:  map[string]string{"quantifier": "3"},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: `Invalid numbers`,
		},
		{
			Endpoint:     "/min",
			QueryParams:  map[string]string{"numbers": "1,2,3,4,5", "quantifier": "3"},
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"Success","result":[1,2,3]}`,
		},
		{
			Endpoint:     "/max",
			QueryParams:  map[string]string{"numbers": "1,2,3,4,5", "quantifier": "2"},
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"Success","result":[4,5]}`,
		},
		{
			Endpoint:     "/avg",
			QueryParams:  map[string]string{"numbers": "1,2,3,4,5"},
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"Success","result":3}`,
		},
		{
			Endpoint:     "/median",
			QueryParams:  map[string]string{"numbers": "1,2,3,4,5"},
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"Success","result":3}`,
		},
		{
			Endpoint:     "/median",
			QueryParams:  map[string]string{"numbers": "1,2,3,4"},
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"Success","result":2.5}`,
		},
		{
			Endpoint:     "/percentile",
			QueryParams:  map[string]string{"numbers": "1,2,3,4,5", "quantifier": "50"},
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"Success","result":3}`,
		},
		{
			Endpoint:     "/percentile",
			QueryParams:  map[string]string{"numbers": "1,2,3,4,5", "quantifier": "101"},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: `Invalid percentile`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Endpoint, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.Endpoint, nil)

			q := req.URL.Query()
			for key, value := range tc.QueryParams {
				q.Set(key, value)
			}
			req.URL.RawQuery = q.Encode()

			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)

			// Verify status code
			if recorder.Code != tc.ExpectedCode {
				t.Errorf("Expected status code %d, but got %d", tc.ExpectedCode, recorder.Code)
			}

			// Verify response body
			body := strings.TrimSpace(recorder.Body.String())
			if body != tc.ExpectedBody {
				t.Errorf("Expected response body '%s', but got '%s'", tc.ExpectedBody, body)
			}

			// Verify error message (if applicable)
			if tc.ExpectedError != "" {
				var response map[string]interface{}
				if err := json.Unmarshal([]byte(body), &response); err != nil {
					t.Errorf("Failed to decode error response: %v", err)
				} else {
					errMsg := response["error"].(string)
					if errMsg != tc.ExpectedError {
						t.Errorf("Expected error message '%s', but got '%s'", tc.ExpectedError, errMsg)
					}
				}
			}
		})
	}
}
