package httprunner

import (
	"errors"
	"net/http"
	"testing"

	"github.com/devMint/go-restful"
	"github.com/devMint/go-restful/request"
	"github.com/devMint/go-restful/response"
	"github.com/go-chi/chi"
)

func Test_RandomEndpoint(t *testing.T) {
	runner := EndToEndTestRunner{
		Server: exampleServer(),
	}

	runner.GET(t, "/", Expectations{
		Name:       "valid endpoint",
		StatusCode: http.StatusOK,
		Body: `{
			"data": "valid endpoint"
		}`,
	})

	runner.GET(t, "/abcd", Expectations{
		Name:       "invalid endpoint",
		StatusCode: http.StatusBadRequest,
		Body: `{
			"detail": "invalid endpoint",
			"title": "Bad Request",
			"type": "http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html",
			"status": 400
		}`,
		Headers: map[string]string{
			"Content-type": "application/json",
		},
	})

	runner.GET(t, "/abcd", Expectations{
		StatusCode: http.StatusBadRequest,
		Body: `<response>
			<type>http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html</type>
			<title>Bad Request</title>
			<detail>invalid endpoint</detail>
			<status>400</status>
		</response>`,
		BodyContains: []string{
			`<title>Bad Request</title>`,
			`<detail>invalid endpoint</detail>`,
		},
		Headers: map[string]string{
			"Content-type":   "application/xml",
			"Dolor-sit-amet": "xkcd",
		},
		Before: func(r *http.Request) {
			r.Header.Set("Content-type", "application/xml")
		},
	})

	runner.POST(t, "/abcd", "", Expectations{
		StatusCode: http.StatusCreated,
		Body:       `{"data":null}`,
	})

	runner.PUT(t, "/a/b/c/d", "", Expectations{})

	runner.PATCH(t, "/a/b", `{"name": "test"}`, Expectations{
		StatusCode: http.StatusOK,
		Body:       `{"data": { "name": "test" }}`,
	})

	runner.PATCH(t, "/a/b", `{}`, Expectations{
		StatusCode: http.StatusBadRequest,
	})
}

func exampleServer() http.Handler {
	router := restful.NewRouter(chi.NewMux())
	router.Get("/", func(r request.Request) response.Response {
		return response.Ok("valid endpoint")
	})
	router.Get("/abcd", func(r request.Request) response.Response {
		raw := response.BadRequest(errors.New("invalid endpoint"))
		raw.WithHeader("Dolor-sit-amet", "xkcd")

		return raw
	})
	router.Post("/abcd", func(r request.Request) response.Response {
		return response.Created()
	})
	router.Put("/a/b/c/d", func(r request.Request) response.Response {
		return response.NoContent()
	})
	router.Patch("/a/b", func(r request.Request) response.Response {
		body := randomBody{}
		if err := r.Body(&body); err != nil {
			return response.BadRequest(err)
		}

		return response.Ok(body)
	})

	return router
}

type randomBody struct {
	Name string `json:"name" validate:"required"`
}
