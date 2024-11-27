package utils

type ComposeFile struct {
	Version  string                 `json:"version"`
	Services map[string]interface{} `json:"services"`
	Volumes  map[string]interface{} `json:"volumes,omitempty"`
	Networks map[string]interface{} `json:"networks,omitempty"`
}
