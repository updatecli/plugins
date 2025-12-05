package internal

import (
	"bytes"
	"text/template"

	"github.com/updatecli/plugins/autodiscovery/demo/internal/filter"
)

// ManifestParams holds the parameters required to generate the Updatecli manifest.
type ManifestParams struct {
	ActionID      string
	ScmID         string
	ImageName     string
	ImageTag      string
	Release       string
	Spec          Spec
	TagFilter     string
	VersionFilter filter.VersionFilter
}

// templateParams holds the parameters used to populate the manifest template.
type templateParams struct {
	ActionID             string
	ImageName            string
	ImageTag             string
	Release              string
	ScmID                string
	SourceID             string
	TagFilter            string
	TargetID             string
	TargetFile           string
	VersionFilterKind    string
	VersionFilterPattern string
}

func generate(params ManifestParams, targetFile string) (string, error) {

	var tmpl *template.Template

	p := templateParams{
		ActionID:             params.ActionID,
		ImageName:            params.ImageName,
		ImageTag:             params.ImageTag,
		Release:              params.Release,
		ScmID:                params.ScmID,
		SourceID:             params.ImageName,
		TagFilter:            params.TagFilter,
		TargetID:             params.ImageName,
		TargetFile:           targetFile,
		VersionFilterKind:    params.VersionFilter.Kind,
		VersionFilterPattern: params.VersionFilter.Pattern,
	}

	tmpl, err := template.New("manifest").Parse(manifestTemplate)
	if err != nil {
		return "", err
	}

	manifest := bytes.Buffer{}
	if err := tmpl.Execute(&manifest, p); err != nil {
		return "", err
	}

	return manifest.String(), nil
}
