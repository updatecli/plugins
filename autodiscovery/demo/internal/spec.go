package internal

// Spec represents the plugin specification.
type Spec struct {
	// Path is the file path to search for manifests.
	Path string `json:"path"`
}
