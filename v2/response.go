// Package client (v2) is the current official Go client for InfluxDB.
package client // import "github.com/aptpod/influxdb1-client/v2"

import "errors"

//go:generate easyjson -all ./${GOFILE}
// Response represents a list of statement results.
type Response struct {
	Results []Result
	Err     string `json:"error,omitempty"`
}

// Error returns the first error from any statement.
// It returns nil if no errors occurred on any statements.
func (r *Response) Error() error {
	if r.Err != "" {
		return errors.New(r.Err)
	}
	for _, result := range r.Results {
		if result.Err != "" {
			return errors.New(result.Err)
		}
	}
	return nil
}
