package internal

import (
	"bytes"
	"html/template"
)

// ManifestParams holds the parameters required to generate the Updatecli manifest.
type ManifestParams struct {
	ActionID  string
	ScmID     string
	ImageName string
	ImageTag  string
	Release   string
	Spec      Spec
}

// TemplateParams holds the parameters used to populate the manifest template.
type TemplateParams struct {
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

func Generate(params ManifestParams) (string, error) {

	var tmpl *template.Template

	p := TemplateParams{
		ActionID:             params.ActionID,
		ImageName:            params.ImageName,
		ImageTag:             params.ImageTag,
		Release:              params.Release,
		ScmID:                params.ScmID,
		SourceID:             params.ImageName,
		TagFilter:            "*",
		TargetID:             params.ImageName,
		TargetFile:           params.Spec.Path,
		VersionFilterKind:    "semver",
		VersionFilterPattern: "*",
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
