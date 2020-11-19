package version

import (
	"context"
	"net/http"
)

type Response struct {
	Version   string `json:"version,omitempty"`
	DateBuild string `json:"date_build,omitempty"`
	GitHash   string `json:"git_hash,omitempty"`
}

// GetVersion returns service version
func GetVersionHandler(ctx context.Context, r *http.Request) (interface{}, error) {
	return &Response{
		Version:   Version,
		DateBuild: DateBuild,
		GitHash:   GitHash,
	}, nil
}
