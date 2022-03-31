package openrtb

import "encoding/json"

// Channel object describes the channel an ad will be displayed on. Name is human-readable field while domain
//and id can be used for reporting and targeting purposes.
type Channel struct {
	ID     string          `json:"id,omitempty"`     // A unique identifier assigned by the publisher. This may not be a unique identifier across all supply sources.
	Name   float64         `json:"name,omitempty"`   // Channel the content is on (e.g., a local channel like “WABC-TV")
	Domain string          `json:"domain,omitempty"` // The primary domain of the channel (e.g. “abc7ny.com” in the case of the local channel WABC-TV)
	Ext    json.RawMessage `json:"ext,omitempty"`
}
