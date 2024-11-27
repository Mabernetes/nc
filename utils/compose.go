package utils

type ComposeFile struct {
	Services map[string]interface{} `json:"services"`
	Volumes  map[string]interface{} `json:"volumes,omitempty" yaml:",omitempty"`
	Networks map[string]interface{} `json:"networks,omitempty" yaml:",omitempty"`
}
