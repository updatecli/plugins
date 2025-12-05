//go:build std
// +build std

package main

import (
	"fmt"

	pdk "github.com/extism/go-pdk"
	"github.com/updatecli/plugins/autodiscovery/demo/internal"
)

//go:wasmexport autodiscovery
func autodiscovery() int32 {

	params := internal.Input{}
	if err := pdk.InputJSON(&params); err != nil {
		pdk.SetError(fmt.Errorf("unable to parse plugin input: %w", err))
		return -1
	}

	hostFunc := wasmHostFunc{}

	results, err := internal.Run(params, hostFunc)
	if err != nil {
		pdk.SetError(fmt.Errorf("running plugin: %w", err))
		return -1
	}

	if results == nil {
		results = &internal.Output{}
	}

	if err := pdk.OutputJSON(results); err != nil {
		pdk.SetError(err)
		return -1
	}

	return 0
}

// Currently, the standard Go compiler cannot export custom functions and is limited to exporting
// `_start` via WASI. So, `main` functions should contain the plugin behavior, that the host will
// invoke by explicitly calling `_start`.
func main() {
	autodiscovery()
}
