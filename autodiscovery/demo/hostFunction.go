package main

import (
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
	"github.com/updatecli/plugins/autodiscovery/demo/internal/filter"
)

type wasmHostFunc struct{}

//go:wasmimport extism:host/user generate_docker_source_spec
func generate_docker_source_spec(uint64) uint64

//go:wasmimport extism:host/user versionfilter_greater_than_pattern
func versionfilter_greater_than_pattern(uint64) uint64

func (w wasmHostFunc) GetDockerFilter(image, tag string) (*filter.Spec, error) {

	input := struct {
		Image string `json:"image"`
		Tag   string `json:"tag"`
	}{
		Image: image,
		Tag:   tag,
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal docker filter spec input: %w", err)
	}

	mem := pdk.AllocateString(string(inputBytes))
	defer mem.Free()

	ptr := generate_docker_source_spec(mem.Offset())
	rmem := pdk.FindMemory(ptr)
	response := string(rmem.ReadBytes())
	pdk.OutputString(response)

	filterSpec := filter.Spec{}

	err = json.Unmarshal([]byte(response), &filterSpec)

	if err != nil {
		return nil, fmt.Errorf("unable to parse docker filter spec: %w", err)
	}

	return &filterSpec, nil

}

// VersionFilterGreaterThanPattern modifies the provided version filter pattern to be greater than the specified pattern.
func (w wasmHostFunc) VersionFilterGreaterThanPattern(versionFilter *filter.VersionFilter, pattern string) error {
	input := struct {
		VersionFilter filter.VersionFilter `json:"versionfilter"`
		Pattern       string               `json:"pattern"`
	}{
		VersionFilter: *versionFilter,
		Pattern:       pattern,
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("unable to marshal docker filter spec input: %w", err)
	}

	mem := pdk.AllocateString(string(inputBytes))
	defer mem.Free()

	ptr := versionfilter_greater_than_pattern(mem.Offset())

	rmem := pdk.FindMemory(ptr)

	response := string(rmem.ReadBytes())

	pdk.OutputString(response)

	result := struct {
		Pattern string `json:"pattern"`
	}{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		return fmt.Errorf("unable to parse docker filter spec: %w", err)
	}

	versionFilter.Pattern = result.Pattern

	return nil
}
