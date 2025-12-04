# README

**This is an experimental Updatecli feature. Feedback is welcome.**

This repository contains the source code for Updatecli plugins.

An Updatecli plugin is a WASM module that extends Updatecli's capabilities.

Currently, Updatecli supports autodiscovery plugins only. More plugin kinds may be added based on user feedback.

Plugins can be written in any language that compiles to WebAssembly (WASM). This project uses the Extism framework: https://extism.org/docs/quickstart/plugin-quickstart.

In this project, the standard Go toolchain is preferred over TinyGo because TinyGo does not support the `text/template` package (see https://tinygo.org/docs/reference/lang-support/stdlib/#texttemplate). `text/template` is useful for generating Updatecli manifests. You may still use TinyGo if you prefer, but be aware of its stdlib limitations.

## Contract

A plugin must export the function `_start` (WASI). It is expected to receive a JSON object as input and return a JSON object containing the generated Updatecli manifests.

### Input

A plugin receives a JSON object with the following fields:

```json
{
  "scmid": "default",
  "actionid": "default",
  "spec": {
    "plugin_param1": "",
    "plugin_param2": ""
  }
}
```

The `scmid` field is provided by Updatecli, it's up to the plugin maintainer to decide how to use that information but typically it's used to link resources to a scm configuration specified in the Updatecli manifest.

The `actionid` field is provided by Updatecli, it's up to the plugin maintainer to decide how to use that information but typically it's used to specify an action title.

The `spec` field is the plugin parameters detected by Updatecli.

### Output

A plugin must output a JSON object such as

```json
{
  "manifests": [
    "Updatecli manifest 1",
    "Updatecli manifest 2"
  ]
}
```

Where `manifests` contains the list of generated Updatecli manifests

Guidelines:
- Include a `version` field in generated manifests to indicate the minimum Updatecli version required.
- Do not set `pipelineid` in generated manifests (Updatecli overrides it).

## Contributing

- This is experimental — please file issues or PRs for improvements.
