// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// CountOption is a non-required Count option that gets applied to an HTTP request.
type CountOption func(r *transport.Request)

// WithCountIndex - a comma-separated list of indices to restrict the results.
func WithCountIndex(index []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountType - a comma-separated list of types to restrict the results.
func WithCountType(documentType []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithCountAllowNoIndices(allowNoIndices bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountAnalyzeWildcard - specify whether wildcard and prefix queries should be analyzed (default: false).
func WithCountAnalyzeWildcard(analyzeWildcard bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountAnalyzer - the analyzer to use for the query string.
func WithCountAnalyzer(analyzer string) CountOption {
	return func(r *transport.Request) {
	}
}

// CountDefaultOperator - the default operator for query string query (AND or OR).
type CountDefaultOperator int

const (
	// CountDefaultOperatorAND can be used to set CountDefaultOperator to "AND"
	CountDefaultOperatorAND = iota
	// CountDefaultOperatorOR can be used to set CountDefaultOperator to "OR"
	CountDefaultOperatorOR = iota
)

// WithCountDefaultOperator - the default operator for query string query (AND or OR).
func WithCountDefaultOperator(defaultOperator CountDefaultOperator) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountDf - the field to use as default where no field prefix is given in the query string.
func WithCountDf(df string) CountOption {
	return func(r *transport.Request) {
	}
}

// CountExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type CountExpandWildcards int

const (
	// CountExpandWildcardsOpen can be used to set CountExpandWildcards to "open"
	CountExpandWildcardsOpen = iota
	// CountExpandWildcardsClosed can be used to set CountExpandWildcards to "closed"
	CountExpandWildcardsClosed = iota
	// CountExpandWildcardsNone can be used to set CountExpandWildcards to "none"
	CountExpandWildcardsNone = iota
	// CountExpandWildcardsAll can be used to set CountExpandWildcards to "all"
	CountExpandWildcardsAll = iota
)

// WithCountExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithCountExpandWildcards(expandWildcards CountExpandWildcards) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithCountIgnoreUnavailable(ignoreUnavailable bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountLenient - specify whether format-based query failures (such as providing text to a numeric field) should be ignored.
func WithCountLenient(lenient bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountMinScore - include only documents with a specific "_score" value in the result.
func WithCountMinScore(minScore int) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountPreference - specify the node or shard the operation should be performed on (default: random).
func WithCountPreference(preference string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountQ - query in the Lucene query string syntax.
func WithCountQ(q string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountRouting - specific routing value.
func WithCountRouting(routing string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountBody - a query to restrict the results specified with the Query DSL (optional).
func WithCountBody(body map[string]interface{}) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountErrorTrace - include the stack trace of returned errors.
func WithCountErrorTrace(errorTrace bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountFilterPath - a comma-separated list of filters used to reduce the respone.
func WithCountFilterPath(filterPath []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountHuman - return human readable values for statistics.
func WithCountHuman(human bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountIgnore - ignores the specified HTTP status codes.
func WithCountIgnore(ignore []int) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountPretty - pretty format the returned JSON response.
func WithCountPretty(pretty bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithCountSourceParam(sourceParam string) CountOption {
	return func(r *transport.Request) {
	}
}

// Count allows to easily execute a query and get the number of matches for that query. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-count.html for more info.
//
// options: optional parameters.
func (a *API) Count(options ...CountOption) (*CountResponse, error) {
	req := a.transport.NewRequest("POST")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &CountResponse{resp}, err
}

// CountResponse is the response for Count.
type CountResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *CountResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
