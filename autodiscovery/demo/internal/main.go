package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Output struct {
	Manifests []string
}

// Input represents the JSON input provided by Updatecli.
type Input struct {
	ScmID    string `json:"scmid"`
	ActionID string `json:"actionid"`
	RootDir  string `json:"rootdir"`
	Spec     Spec   `json:"spec"`
}

type filterSpec struct {
	VersionFilter VersionFilter `json:"versionfilter"`
	TagFilter     string        `json:"tagfilter"`
}

func Run(params Input) (*Output, error) {

	var results Output
	var errs []error

	processedFiles := make(map[string]bool)

	if len(params.Spec.Files) == 0 {
		params.Spec.Files = []string{DefaultPath}
	}

	for i := range params.Spec.Files {

		datafile := params.Spec.Files[i]

		if _, ok := processedFiles[datafile]; ok {
			continue
		}

		file, err := os.Open(datafile)
		if err != nil {
			return nil, fmt.Errorf("unable to open data file: %w", err)
		}

		deferClose := func() {
			if err := file.Close(); err != nil {
				fmt.Printf("warning: unable to close data file: %v\n", err)
			}
		}

		defer deferClose()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			imageName, imageTag, release := parseLine(strings.TrimSpace(line))

			// Skip entries with no tag
			if imageTag == "" || imageName == "" || imageTag == "latest" {
				continue
			}

			if len(params.Spec.Ignore) > 0 {
				if params.Spec.Ignore.isMatchingRules(params.RootDir, datafile, imageName) {
					fmt.Printf("Ignoring container %q from %q, as matching ignore rule(s)\n", imageName, datafile)
					continue
				}
			}

			if len(params.Spec.Only) > 0 {
				if !params.Spec.Only.isMatchingRules(params.RootDir, datafile, imageName) {
					fmt.Printf("Ignoring container %q from %q, as not matching only rule(s)\n", imageName, datafile)
					continue
				}
			}

			tagFilter := "*"
			// By default, we use a semver filter with wildcard pattern
			versionFilter := VersionFilter{
				Kind:    "semver",
				Pattern: "*",
			}

			dockerFilterSpec, err := getDockerFilter(imageName, imageTag)
			if err != nil {
				return nil, fmt.Errorf("unable to call getDockerFilter function %v\n%v", err, dockerFilterSpec)
			}

			if dockerFilterSpec != nil {
				versionFilter = dockerFilterSpec.VersionFilter
				tagFilter = dockerFilterSpec.TagFilter
			}

			// Override version filter if specified in the spec
			if params.Spec.VersionFilter.Kind != "" {
				versionFilter.Kind = params.Spec.VersionFilter.Kind
				versionFilter.Pattern = params.Spec.VersionFilter.Pattern

				err = versionFilterGreaterThanPattern(&versionFilter, imageTag)
				if err != nil {
					return nil, fmt.Errorf("unable to call versionFilterGreaterThanPattern function: %w", err)
				}
			}

			if params.Spec.VersionFilter.Kind != "" {
				dockerFilterSpec.VersionFilter = params.Spec.VersionFilter
			}

			manifest, err := generate(ManifestParams{
				ImageName:     imageName,
				ImageTag:      imageTag,
				Release:       release,
				ActionID:      params.ActionID,
				ScmID:         params.ScmID,
				Spec:          params.Spec,
				VersionFilter: versionFilter,
				TagFilter:     tagFilter,
			}, datafile)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			results.Manifests = append(results.Manifests, manifest)
		}

		if _, ok := processedFiles[datafile]; !ok {
			processedFiles[datafile] = true
		}
	}

	if len(errs) > 0 {
		errorMsgs := []string{}
		for _, e := range errs {
			errorMsgs = append(errorMsgs, e.Error())
		}

		return nil, errors.New(strings.Join(errorMsgs, "\n"))
	}

	return &results, nil
}
