package openrtb

import "encoding/json"

// Metric object is associated with an impression as an array of metrics. These metrics can offer insight into
// the impression to assist with decisioning such as average recent viewability, click-through rate, etc. Each
// metric is identified by its type, reports the value of the metric, and optionally identifies the source or
// vendor measuring the value.
type Metric struct {
	Type   string          `json:"type"`             // Type of metric being presented using exchange curated string names which should be published to bidders a priori
	Value  float64         `json:"value"`            // Number representing the value of the metric. Probabilities must be in the range 0.0 – 1.0.
	Vendor string          `json:"vendor,omitempty"` // Source of the value using exchange curated string names which should be published to bidders a priori. If the exchange itself is the source versus a third party, “EXCHANGE” is recommended.
	Ext    json.RawMessage `json:"ext,omitempty"`
}
