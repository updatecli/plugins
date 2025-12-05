package internal

var (
	// manifestTemplate is the template used to generate the Updatecli manifest for each discovered image.
	manifestTemplate = `name: 'deps: bump {{ .ImageName }} tag'
{{- if .ActionID }}
actions:
  {{ .ActionID }}:
    title: 'deps: update Docker image "{{ .ImageName }}" to "{{ "{{" }} source "{{ .SourceID }}" {{ "}}" }}"'
{{- end }}
sources:
  '{{ .SourceID }}':
    name: 'get latest image tag for "{{ .ImageName }}"'
    kind: 'dockerimage'
    spec:
      image: '{{ .ImageName }}'
      {{- if .TagFilter }}
      tagfilter: '{{ .TagFilter }}'
      {{- end }}
      versionfilter:
        kind: '{{ .VersionFilterKind }}'
        pattern: '{{ .VersionFilterPattern }}'
targets:
  '{{ .TargetID }}':
    name: 'deps: update Docker image "{{ .ImageName }}" to {{ "{{" }} source "{{ .SourceID }}" {{ "}}" }}'
    kind: 'file'
{{- if .ScmID }}
    scmid: '{{ .ScmID }}'
    disablesourceinput: true
{{- end }}
    spec:
      file: '{{ .TargetFile }}'
      matchpattern: '{{ .ImageName }}:(.*) ({{ .Release }})'
      replacepattern: '{{ .ImageName }}:{{ "{{" }} source "{{ .SourceID }}" {{ "}}" }} {{ .Release }}'
`
)
