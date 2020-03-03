// Package client (v2) is the current official Go client for InfluxDB.
package client // import "github.com/aptpod/influxdb1-client/v2"

import "github.com/aptpod/influxdb1-client/models"

//go:generate easyjson -all ./${GOFILE}
// Result represents a resultset returned from a single statement.
type Result struct {
	StatementId int `json:"statement_id"`
	Series      []models.Row
	Messages    []*Message
	Err         string `json:"error,omitempty"`
}
