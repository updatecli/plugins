package internal

// Spec represents the plugin specification.
type Spec struct {
	// Files is the list of file path to search for manifests.
	Files []string `json:"files"`
	// Ignore allows to specify rule to ignore autodiscovery a specific Kubernetes manifest based on a rule
	Ignore matchingRules `json:",omitempty"`
	// Only allows to specify rule to only autodiscover manifest for a specific Kubernetes manifest based on a rule
	Only matchingRules `json:",omitempty"`
	// versionfilter provides parameters to specify the version pattern used when generating manifest.
	//
	// kind - semver
	//   versionfilter of kind `semver` uses semantic versioning as version filtering
	//   pattern accepts one of:
	//     `patch` - patch only update patch version
	//     `minor` - minor only update minor version
	//     `major` - major only update major versions
	//     `a version constraint` such as `>= 1.0.0`

	// kind - regex
	//   versionfilter of kind `regex` uses regular expression as version filtering
	//   pattern accepts a valid regular expression

	// example:
	// ```
	//  versionfilter:
	//    kind: semver
	//    pattern: minor
	// ```

	//	and its type like regex, semver, or just latest.
	VersionFilter VersionFilter `json:",omitempty"`
}

type VersionFilter struct {
	// Kind defines which kind of version filter to use
	Kind string `json:",omitempty"`
	// Pattern defines the version filter pattern
	Pattern string `json:",omitempty"`
}
