package url

import (
	"net/url"
	"strings"
)

// URLBuilder provides URL building functionality with configurable root URL
type URLBuilder struct {
	rootURL string
}

// NewURLBuilder creates a new URLBuilder with the given root URL
func NewURLBuilder(rootURL string) *URLBuilder {
	return &URLBuilder{rootURL: rootURL}
}

// RootURL returns the configured root URL
func (ub *URLBuilder) RootURL() string {
	return ub.rootURL
}

// BuildURL returns the full URL for a given path with optional query parameters
func (ub *URLBuilder) BuildURL(path string, params map[string]string) string {
	if ub.rootURL == "" {
		// If no root URL is configured, return just the path with query params
		// This is useful for testing environments
		return "/" + strings.TrimPrefix(path, "/") + ub.BuildQuery(params)
	}
	return ub.rootURL + "/" + strings.TrimPrefix(path, "/") + ub.BuildQuery(params)
}

// BuildQuery creates a query string from a map of parameters
func (ub *URLBuilder) BuildQuery(queryData map[string]string) string {
	queryString := ""

	if len(queryData) > 0 {
		v := url.Values{}
		for key, value := range queryData {
			v.Set(key, value)
		}
		queryString += "?" + ub.HttpBuildQuery(v)
	}

	return queryString
}

// HttpBuildQuery converts url.Values to a query string
func (ub *URLBuilder) HttpBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}

// Convenience functions for standalone usage
var defaultBuilder = NewURLBuilder("")

// SetDefaultURL sets the default URL builder's root URL
func SetDefaultURL(rootURL string) {
	defaultBuilder = NewURLBuilder(rootURL)
}

// RootURL returns the default root URL
func RootURL() string {
	return defaultBuilder.RootURL()
}

// BuildURL returns the full URL for a given path using default builder
func BuildURL(path string, params map[string]string) string {
	return defaultBuilder.BuildURL(path, params)
}

// BuildQuery creates a query string using default builder
func BuildQuery(queryData map[string]string) string {
	return defaultBuilder.BuildQuery(queryData)
}

// HttpBuildQuery converts url.Values to a query string using default builder
func HttpBuildQuery(queryData url.Values) string {
	return defaultBuilder.HttpBuildQuery(queryData)
}
