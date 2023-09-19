package resolution

import (
	"encoding/json"
)

type TokenMetadataAttribute struct {
	DisplayType string      `json:"display_type"`
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
}
type TokenMetadataProperties struct {
	Records map[string]string `json:"records"`
}

type TokenMetadata struct {
	Name            string                   `json:"name"`
	Description     string                   `json:"description"`
	Image           string                   `json:"image"`
	ExternalUrl     string                   `json:"-"`
	ExternalLink    string                   `json:"external_link"`
	ImageData       string                   `json:"image_data"`
	BackgroundColor string                   `json:"background_color"`
	AnimationUrl    string                   `json:"animation_url"`
	YoutubeUrl      string                   `json:"youtube_url"`
	Properties      TokenMetadataProperties  `json:"properties"`
	Attributes      []TokenMetadataAttribute `json:"attributes"`
}

func (tm *TokenMetadata) UnmarshalJSON(data []byte) error {
	// Use a helper struct to aid the unmarshalling
	type Alias TokenMetadata
	aux := &struct {
		ExternalLink1 string `json:"external_link"`
		ExternalLink2 string `json:"url"`
		*Alias
	}{
		Alias: (*Alias)(tm),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Logic to set ExternalUrl based on which one is present
	tm.ExternalUrl = aux.ExternalLink1
	if tm.ExternalLink == "" {
		tm.ExternalLink = aux.ExternalLink2
	}
	return nil
}
