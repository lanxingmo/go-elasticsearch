// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// TermvectorsOption is a non-required Termvectors option that gets applied to an HTTP request.
type TermvectorsOption func(r *transport.Request)

// WithTermvectorsID - the id of the document, when not specified a doc param should be supplied.
func WithTermvectorsID(id string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsFieldStatistics - specifies if document count, sum of document frequencies and sum of total term frequencies should be returned.
func WithTermvectorsFieldStatistics(fieldStatistics bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsFields - a comma-separated list of fields to return.
func WithTermvectorsFields(fields []string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsOffsets - specifies if term offsets should be returned.
func WithTermvectorsOffsets(offsets bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsParent - parent id of documents.
func WithTermvectorsParent(parent string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsPayloads - specifies if term payloads should be returned.
func WithTermvectorsPayloads(payloads bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsPositions - specifies if term positions should be returned.
func WithTermvectorsPositions(positions bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsPreference - specify the node or shard the operation should be performed on (default: random).
func WithTermvectorsPreference(preference string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsRealtime - specifies if request is real-time as opposed to near-real-time (default: true).
func WithTermvectorsRealtime(realtime bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsRouting - specific routing value.
func WithTermvectorsRouting(routing string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsTermStatistics - specifies if total term frequency and document frequency should be returned.
func WithTermvectorsTermStatistics(termStatistics bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsVersion - explicit version number for concurrency control.
func WithTermvectorsVersion(version int) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// TermvectorsVersionType - specific version type.
type TermvectorsVersionType int

const (
	// TermvectorsVersionTypeInternal can be used to set TermvectorsVersionType to "internal"
	TermvectorsVersionTypeInternal = iota
	// TermvectorsVersionTypeExternal can be used to set TermvectorsVersionType to "external"
	TermvectorsVersionTypeExternal = iota
	// TermvectorsVersionTypeExternalGte can be used to set TermvectorsVersionType to "external_gte"
	TermvectorsVersionTypeExternalGte = iota
	// TermvectorsVersionTypeForce can be used to set TermvectorsVersionType to "force"
	TermvectorsVersionTypeForce = iota
)

// WithTermvectorsVersionType - specific version type.
func WithTermvectorsVersionType(versionType TermvectorsVersionType) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsBody - define parameters and or supply a document to get termvectors for. See documentation.
func WithTermvectorsBody(body map[string]interface{}) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsErrorTrace - include the stack trace of returned errors.
func WithTermvectorsErrorTrace(errorTrace bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsFilterPath - a comma-separated list of filters used to reduce the respone.
func WithTermvectorsFilterPath(filterPath []string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsHuman - return human readable values for statistics.
func WithTermvectorsHuman(human bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsIgnore - ignores the specified HTTP status codes.
func WithTermvectorsIgnore(ignore []int) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsPretty - pretty format the returned JSON response.
func WithTermvectorsPretty(pretty bool) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithTermvectorsSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithTermvectorsSourceParam(sourceParam string) TermvectorsOption {
	return func(r *transport.Request) {
	}
}

// Termvectors - returns information and statistics on terms in the fields of a particular document. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-termvectors.html for more info.
//
// index: the index in which the document resides.
//
// documentType: the type of the document.
//
// options: optional parameters.
func (a *API) Termvectors(index string, documentType string, options ...TermvectorsOption) (*TermvectorsResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &TermvectorsResponse{resp}, err
}

// TermvectorsResponse is the response for Termvectors.
type TermvectorsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *TermvectorsResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
