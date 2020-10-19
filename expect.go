package httprunner

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Expectations
type Expectations struct {
	Name         string
	StatusCode   int
	Body         string
	BodyContains []string
	Headers      map[string]string
	Before       func(r *http.Request)
	After        func(r *httptest.ResponseRecorder)

	request  *http.Request
	response *httptest.ResponseRecorder
	t        *testing.T
}

func (e Expectations) assertStatusCode() {
	if e.StatusCode == 0 {
		return
	}

	e.t.Run("assert status code", func(t *testing.T) {
		msg := fmt.Sprintf("Expected status code %d, got %d", e.StatusCode, e.response.Code)
		assert.Equal(t, e.StatusCode, e.response.Code, msg)
	})
}

func (e Expectations) assertHeaders() {
	if len(e.Headers) == 0 {
		return
	}

	e.t.Run("assert response headers", func(t *testing.T) {
		for k, v := range e.Headers {
			headerValue := e.response.Header().Get(k)
			if assert.NotEmpty(t, headerValue, "Response headers does not contain header '%s'", k) {
				assert.Equal(t, v, headerValue, "Header '%s' should equal to '%s', equals '%s'", k, v, headerValue)
			}
		}
	})
}

func (e Expectations) assertResponseBody() {
	if len(e.Body) == 0 {
		return
	}

	e.t.Run("assert response body", func(t *testing.T) {
		msg := fmt.Sprintf("There is a difference between '%s' and '%s'", e.Body, e.response.Body.String())
		contentType := e.response.Header().Get("Content-type")
		if contentType == "" {
			contentType = "application/json"
		}

		switch contentType {
		case "application/json":
			jsonEquals(t, e.Body, e.response.Body.String(), msg)
		case "application/xml":
			xmlEquals(t, e.Body, e.response.Body.String(), msg)
		}
	})
}

func (e Expectations) assertResponseBodyContains() {
	if len(e.BodyContains) == 0 {
		return
	}

	for num, contains := range e.BodyContains {
		testName := fmt.Sprintf("assert response body contains #%d", num)
		e.t.Run(testName, func(t *testing.T) {
			assert.Contains(t, e.response.Body.String(), contains)
		})
	}
}
