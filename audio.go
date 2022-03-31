package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidAudioNoMIMEs = errors.New("openrtb: audio has no mimes")
)

// Audio object must be included directly in the impression object
type Audio struct {
	MIMEs          []string            `json:"mimes"`                 // Content MIME types supported.
	MinDuration    int                 `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration    int                 `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols      []Protocol          `json:"protocols,omitempty"`   // Video bid response protocols
	StartDelay     StartDelay          `json:"startdelay,omitempty"`  // Indicates the start delay in seconds
	Sequence       int                 `json:"sequence,omitempty"`    // Default: 1
	BlockedAttrs   []CreativeAttribute `json:"battr,omitempty"`       // Blocked creative attributes
	MaxExtended    int                 `json:"maxextended,omitempty"` // Maximum extended video ad duration
	MinBitrate     int                 `json:"minbitrate,omitempty"`  // Minimum bit rate in Kbps
	MaxBitrate     int                 `json:"maxbitrate,omitempty"`  // Maximum bit rate in Kbps
	Delivery       []ContentDelivery   `json:"delivery,omitempty"`    // List of supported delivery methods
	CompanionAds   []Banner            `json:"companionad,omitempty"`
	APIs           []APIFramework      `json:"api,omitempty"`
	CompanionTypes []CompanionType     `json:"companiontype,omitempty"`
	MaxSequence    int                 `json:"maxseq,omitempty"`       // The maximumnumber of ads that canbe played in an ad pod.
	Feed           FeedType            `json:"feed,omitempty"`         // Type of audio feed.
	Stitched       int                 `json:"stitched,omitempty"`     // Indicates if the ad is stitched with audio content or delivered independently
	VolumeNorm     VolumeNorm          `json:"nvol,omitempty"`         // Volume normalization mode.
	PodDuration    int                 `json:"poddur,omitempty"`       // Indicates the total amount of time in seconds that advertisers may fill for a “dynamic” video ad pod.
	PodID          int                 `json:"podid,omitempty"`        // Unique identifier indicating that an impression opportunity belongs to a video ad pod.
	PodSequence    int                 `json:"podseq,omitempty"`       // The sequence (position) of the audio ad pod within a content stream.
	MinCPMPerSec   float64             `json:"mincpmpersec,omitempty"` // Minimum CPM per second.
	SlotInPod      int                 `json:"slotinpod,omitempty"`    // For audio ad pods, this value indicates that the seller can guarantee delivery against the indicated sequence.
	Ext            json.RawMessage     `json:"ext,omitempty"`
}

type jsonAudio Audio

// Validate the object
func (a *Audio) Validate() error {
	if len(a.MIMEs) == 0 {
		return ErrInvalidAudioNoMIMEs
	}
	return nil
}

// MarshalJSON custom marshalling with normalization
func (a *Audio) MarshalJSON() ([]byte, error) {
	a.normalize()
	return json.Marshal((*jsonAudio)(a))
}

// UnmarshalJSON custom unmarshalling with normalization
func (a *Audio) UnmarshalJSON(data []byte) error {
	var h jsonAudio
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*a = (Audio)(h)
	a.normalize()
	return nil
}

func (a *Audio) normalize() {
	if a.Sequence == 0 {
		a.Sequence = 1
	}
}
