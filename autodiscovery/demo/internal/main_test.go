package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/updatecli/plugins/autodiscovery/demo/internal/filter"
)

type mockHF struct{}

func (m mockHF) GetDockerFilter(image, tag string) (*filter.Spec, error) {
	return &filter.Spec{
		VersionFilter: filter.VersionFilter{
			Kind:    "semver",
			Pattern: "*",
		},
		TagFilter: "*",
	}, nil
}

func (m mockHF) VersionFilterGreaterThanPattern(versionFilter *filter.VersionFilter, pattern string) error {
	versionFilter.Kind = "semver"
	versionFilter.Pattern = ">" + pattern
	return nil
}

func TestRun(t *testing.T) {
	testdata := []struct {
		input          Input
		expectedOutput Output
	}{
		{
			input: Input{
				Spec: Spec{
					Files: []string{
						"testdata/data.txt",
					},
				},
			},
			expectedOutput: Output{
				Manifests: []string{
					`name: 'deps: bump fluent/fluent-bit tag'
sources:
  'fluent/fluent-bit':
    name: 'get latest image tag for "fluent/fluent-bit"'
    kind: 'dockerimage'
    spec:
      image: 'fluent/fluent-bit'
      tagfilter: '*'
      versionfilter:
        kind: 'semver'
        pattern: '*'
targets:
  'fluent/fluent-bit':
    name: 'deps: update Docker image "fluent/fluent-bit" to {{ source "fluent/fluent-bit" }}'
    kind: 'file'
    spec:
      file: 'testdata/data.txt'
      matchpattern: 'fluent/fluent-bit(.*) (harvester,release/harvester/v1.4-head,release/harvester/v1.4.3)'
      replacepattern: 'fluent/fluent-bit:{{ source "fluent/fluent-bit" }} harvester,release/harvester/v1.4-head,release/harvester/v1.4.3'
`,
				},
			},
		},
	}

	for _, td := range testdata {
		mockHF := mockHF{}
		gotOutput, err := Run(td.input, mockHF)
		assert.NoError(t, err)

		assert.Equalf(t, td.expectedOutput, *gotOutput, "strings differ:\nexpected=%s\nactual=%s", td.expectedOutput, *gotOutput)
	}
}
