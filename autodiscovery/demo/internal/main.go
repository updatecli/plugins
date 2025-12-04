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
			if imageTag == "" || imageName == "" {
				continue
			}

			manifest, err := Generate(ManifestParams{
				ImageName: imageName,
				ImageTag:  imageTag,
				Release:   release,
				ActionID:  params.ActionID,
				ScmID:     params.ScmID,
				Spec:      params.Spec,
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
