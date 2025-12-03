package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	testdata := []struct {
		line            string
		expectedImage   string
		expectedRelease string
		expectedTag     string
	}{
		{
			line:            "kubeovn/kube-ovn:v1.14.10 harvester,release/harvester/master\n",
			expectedImage:   "kubeovn/kube-ovn",
			expectedRelease: "harvester,release/harvester/master",
			expectedTag:     "v1.14.10",
		},
		{
			line:            "longhornio/backing-image-manager:v1.10.1 harvester,release/harvester/master\n",
			expectedImage:   "longhornio/backing-image-manager",
			expectedRelease: "harvester,release/harvester/master",
			expectedTag:     "v1.10.1",
		},
	}

	for _, td := range testdata {
		gotImage, gotImageTag, gotRelease := parseLine(td.line)

		assert.Equal(t, td.expectedImage, gotImage)
		assert.Equal(t, td.expectedRelease, gotRelease)
		assert.Equal(t, td.expectedTag, gotImageTag)
	}
}
