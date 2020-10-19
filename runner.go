package httprunner

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// EndToEndTestRunner
type EndToEndTestRunner struct {
	Server http.Handler
	Before func(r *http.Request)
	After  func(r *httptest.ResponseRecorder)
}

func (h *EndToEndTestRunner) GET(t *testing.T, endpoint string, expect Expectations) {
	h.makeRequest(t, http.MethodGet, endpoint, "", expect)
}

func (h *EndToEndTestRunner) POST(t *testing.T, endpoint string, body string, expect Expectations) {
	h.makeRequest(t, http.MethodPost, endpoint, body, expect)
}

func (h *EndToEndTestRunner) PUT(t *testing.T, endpoint string, body string, expect Expectations) {
	h.makeRequest(t, http.MethodPut, endpoint, body, expect)
}

func (h *EndToEndTestRunner) PATCH(t *testing.T, endpoint string, body string, expect Expectations) {
	h.makeRequest(t, http.MethodPatch, endpoint, body, expect)
}

func (h *EndToEndTestRunner) HEAD(t *testing.T, endpoint string, expect Expectations) {
	h.makeRequest(t, http.MethodHead, endpoint, "", expect)
}

func (h *EndToEndTestRunner) DELETE(t *testing.T, endpoint string, expect Expectations) {
	h.makeRequest(t, http.MethodDelete, endpoint, "", expect)
}

func (h *EndToEndTestRunner) CONNECT(t *testing.T, endpoint string, expect Expectations) {
	h.makeRequest(t, http.MethodConnect, endpoint, "", expect)
}

func (h *EndToEndTestRunner) OPTIONS(t *testing.T, endpoint string, expect Expectations) {
	h.makeRequest(t, http.MethodOptions, endpoint, "", expect)
}

func (h *EndToEndTestRunner) TRACE(t *testing.T, endpoint string, expect Expectations) {
	h.makeRequest(t, http.MethodTrace, endpoint, "", expect)
}

func (h *EndToEndTestRunner) makeRequest(t *testing.T, method string, endpoint string, body string, expect Expectations) {
	name := expect.Name
	if name == "" {
		name = fmt.Sprintf("%s %s", method, endpoint)
		name = strings.Replace(name, "/", "", 1)
	}

	t.Run(name, func(t *testing.T) {
		request, err := http.NewRequest(method, endpoint, strings.NewReader(body))
		if err != nil {
			t.Fatalf("cannot construct request (%s): %s", t.Name(), err.Error())
		}

		expect.t = t
		expect.response = httptest.NewRecorder()
		expect.request = request

		if h.Before != nil {
			h.Before(expect.request)
		}
		if expect.Before != nil {
			expect.Before(expect.request)
		}

		h.Server.ServeHTTP(expect.response, expect.request)

		expect.assertStatusCode()
		expect.assertHeaders()
		expect.assertResponseBody()
		expect.assertResponseBodyContains()

		if expect.After != nil {
			expect.After(expect.response)
		}
		if h.After != nil {
			h.After(expect.response)
		}
	})
}
