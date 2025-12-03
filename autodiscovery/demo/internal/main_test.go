package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	testdata := []struct {
		input          Input
		expectedOutput Output
	}{
		{
			input: Input{
				Spec: Spec{
					Path: "testdata/data.txt",
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
      matchpattern: '(.*) (harvester,release/harvester/v1.4-head,release/harvester/v1.4.3)'
      replacepattern: 'fluent/fluent-bit:{{ source "fluent/fluent-bit" }} harvester,release/harvester/v1.4-head,release/harvester/v1.4.3'
`,
				},
			},
		},
	}

	for _, td := range testdata {
		gotOutput, err := Run(td.input)
		assert.NoError(t, err)

		assert.Equalf(t, td.expectedOutput, *gotOutput, "strings differ:\nexpected=%s\nactual=%s", td.expectedOutput, *gotOutput)
	}
}
