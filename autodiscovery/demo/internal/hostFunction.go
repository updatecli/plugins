package internal

import (
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

//go:wasmimport extism:host/user generate_docker_source_spec
func generate_docker_source_spec(uint64) uint64

func getDockerFilter(image, tag string) (*filterSpec, error) {

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

	filterSpec := filterSpec{}

	err = json.Unmarshal([]byte(response), &filterSpec)

	if err != nil {
		return nil, fmt.Errorf("unable to parse docker filter spec: %w", err)
	}

	return &filterSpec, nil

}

//go:wasmimport extism:host/user versionfilter_greater_than_pattern
func versionfilter_greater_than_pattern(uint64) uint64

// versionFilterGreaterThanPattern modifies the provided version filter pattern to be greater than the specified pattern.
func versionFilterGreaterThanPattern(versionFilter *VersionFilter, pattern string) error {
	input := struct {
		VersionFilter VersionFilter `json:"versionfilter"`
		Pattern       string        `json:"pattern"`
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
