package filter

type Spec struct {
	VersionFilter VersionFilter `json:"versionfilter"`
	TagFilter     string        `json:"tagfilter"`
}

type VersionFilter struct {
	// Kind defines which kind of version filter to use
	Kind string `json:",omitempty"`
	// Pattern defines the version filter pattern
	Pattern string `json:",omitempty"`
}
