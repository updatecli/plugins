# README

This plugin is a proof of concept to see how to automate Docker image tag update
in a custom data file such as:

```
fluent/fluent-bit:2.1.8 updatecli
ghcr.io/kube-vip/kube-vip-iptables:v0.9.2 example2
```

This plugin will parse a file for all Docker images and then try to update to the next version based according
to the autodiscovery configuration.

It leverages WASM Host function `generate_docker_source_spec` and `versionfilter_greater_than_pattern`
To have a similar experience than Updatecli native plugin.

## How To

- `make build` to build the binary named "demo.wasm"
- `make test` to run the UT

## Example

```
autodiscovery:
  crawlers:
    "demo.wasm":
      spec:
        files:
          - data.txt
        # ignore:
        #  - images:
        #    - "fluent/fluent-bit"
        only:
          - images:
              - "fluent/fluent-bit"
        versionfilter:
          kind: "semver"
          # Patch version update only
          # Accept, minor and major
          pattern: "patch"
```
