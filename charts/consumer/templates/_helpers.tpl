{{/*
Return a consistent app name (used for metadata.name, labels, etc.)
*/}}
{{- define "consumer.name" -}}
{{- default "consumer" .Values.nameOverride }}
{{- end }}

{{/*
Simple selector labels
*/}}
{{- define "consumer.labels" -}}
app: {{ include "consumer.name" . }}
{{- end }}

{{- define "consumer.selectorLabels" -}}
app: {{ include "consumer.name" . }}
{{- end }}
