// Package client (v2) is the current official Go client for InfluxDB.
package client // import "github.com/aptpod/influxdb1-client/v2"

//go:generate easyjson -all ./${GOFILE}
// Message represents a user message.
type Message struct {
	Level string `json:"level"`
	Text  string `json:"text"`
}
