package internal

// Spec represents the plugin specification.
type Spec struct {
	// Files is the list of file path to search for manifests.
	Files []string `json:"files"`
}
