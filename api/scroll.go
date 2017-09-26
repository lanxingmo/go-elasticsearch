// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// ScrollOption is a non-required Scroll option that gets applied to an HTTP request.
type ScrollOption func(r *transport.Request)

// WithScrollScrollID - the scroll ID.
func WithScrollScrollID(scrollID string) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollScroll - specify how long a consistent view of the index should be maintained for scrolled search.
func WithScrollScroll(scroll time.Duration) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollScrollIDParam - the scroll ID for scrolled search.
func WithScrollScrollIDParam(scrollIDParam string) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollBody - the scroll ID if not passed by URL or query parameter.
func WithScrollBody(body map[string]interface{}) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollErrorTrace - include the stack trace of returned errors.
func WithScrollErrorTrace(errorTrace bool) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollFilterPath - a comma-separated list of filters used to reduce the respone.
func WithScrollFilterPath(filterPath []string) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollHuman - return human readable values for statistics.
func WithScrollHuman(human bool) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollIgnore - ignores the specified HTTP status codes.
func WithScrollIgnore(ignore []int) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollPretty - pretty format the returned JSON response.
func WithScrollPretty(pretty bool) ScrollOption {
	return func(r *transport.Request) {
	}
}

// WithScrollSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithScrollSourceParam(sourceParam string) ScrollOption {
	return func(r *transport.Request) {
	}
}

// Scroll - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-request-scroll.html for more info.
//
// options: optional parameters.
func (a *API) Scroll(options ...ScrollOption) (*ScrollResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &ScrollResponse{resp}, err
}

// ScrollResponse is the response for Scroll.
type ScrollResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ScrollResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
