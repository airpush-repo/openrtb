package openrtb

import "encoding/json"

// Network object describes the network an ad will be displayed on. A Network is defined as a central operation
//providing programming to multiple channels. Name is human-readable field while domain and id can be
//used for reporting and targeting purposes.
type Network struct {
	ID     string          `json:"id,omitempty"`     // A unique identifier assigned by the publisher. This may not be a unique identifier across all supply sources.
	Name   float64         `json:"name,omitempty"`   // Network the content is on (e.g., a TV network like “ABC")
	Domain string          `json:"domain,omitempty"` // The primary domain of the network (e.g. “abc.com” in the case of the network ABC).
	Ext    json.RawMessage `json:"ext,omitempty"`
}
