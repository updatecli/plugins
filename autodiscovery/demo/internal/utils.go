package internal

import (
	"strings"
)

// parseLine a line from the data file and return the image name and release version.
func parseLine(line string) (imageName, imageTag, release string) {

	lineArray := strings.Split(strings.TrimSuffix(line, "\n"), " ")

	switch len(lineArray) {
	case 2:
		imageName = lineArray[0]
		release = lineArray[1]
	default:
		return "", "", ""
	}

	imageParts := strings.Split(imageName, ":")

	switch len(imageParts) {
	case 2:
		imageName = imageParts[0]
		imageTag = imageParts[1]
	case 1:
		imageName = imageParts[0]
		imageTag = ""
	default:
		return "", "", ""
	}

	return imageName, imageTag, release
}
