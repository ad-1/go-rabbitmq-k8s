{{/*
Return a consistent app name (used for metadata.name, labels, etc.)
*/}}
{{- define "producer.name" -}}
{{- default "producer" .Values.nameOverride }}
{{- end }}

{{/*
Simple selector labels
*/}}
{{- define "producer.labels" -}}
app: {{ include "producer.name" . }}
{{- end }}

{{- define "producer.selectorLabels" -}}
app: {{ include "producer.name" . }}
{{- end }}
