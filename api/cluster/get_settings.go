// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// GetSettingsOption is a non-required GetSettings option that gets applied to an HTTP request.
type GetSettingsOption func(r *transport.Request)

// WithGetSettingsFlatSettings - return settings in flat format (default: false).
func WithGetSettingsFlatSettings(flatSettings bool) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsIncludeDefaults - whether to return all default clusters setting.
func WithGetSettingsIncludeDefaults(includeDefaults bool) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsMasterTimeout - explicit operation timeout for connection to master node.
func WithGetSettingsMasterTimeout(masterTimeout time.Duration) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsTimeout - explicit operation timeout.
func WithGetSettingsTimeout(timeout time.Duration) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsErrorTrace - include the stack trace of returned errors.
func WithGetSettingsErrorTrace(errorTrace bool) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsFilterPath - a comma-separated list of filters used to reduce the respone.
func WithGetSettingsFilterPath(filterPath []string) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsHuman - return human readable values for statistics.
func WithGetSettingsHuman(human bool) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsIgnore - ignores the specified HTTP status codes.
func WithGetSettingsIgnore(ignore []int) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsPretty - pretty format the returned JSON response.
func WithGetSettingsPretty(pretty bool) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// WithGetSettingsSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithGetSettingsSourceParam(sourceParam string) GetSettingsOption {
	return func(r *transport.Request) {
	}
}

// GetSettings - allows to update cluster wide specific settings. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-update-settings.html for more info.
//
// options: optional parameters.
func (c *Cluster) GetSettings(options ...GetSettingsOption) (*GetSettingsResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &GetSettingsResponse{resp}, err
}

// GetSettingsResponse is the response for GetSettings.
type GetSettingsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetSettingsResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
