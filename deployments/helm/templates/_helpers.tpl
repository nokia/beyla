{{/*
Expand the name of the chart.
*/}}
{{- define "beyla.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "beyla.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Allow the release namespace to be overridden for multi-namespace deployments in combined charts
*/}}
{{- define "beyla.namespace" -}}
{{- if .Values.namespaceOverride }}
{{- .Values.namespaceOverride }}
{{- else }}
{{- .Release.Namespace }}
{{- end }}
{{- end }}


{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "beyla.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "beyla.labels" -}}
helm.sh/chart: {{ include "beyla.chart" . }}
{{ include "beyla.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "beyla.selectorLabels" -}}
app.kubernetes.io/name: {{ include "beyla.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "beyla.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "beyla.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}


{{/*
Returns the internal metrics port if set via environment variable or via yaml configuration.
Note: Precedence is given for environment variable
*/}}
{{- define "beyla.internalMetricsPort" -}}
{{- if .Values.env.BEYLA_INTERNAL_METRICS_PROMETHEUS_PORT }}
{{- print .Values.env.BEYLA_INTERNAL_METRICS_PROMETHEUS_PORT }}
{{ else if and (.Values.configmapData.prometheus_export) (ne (.Values.configmapData.prometheus_export.port | quote ) "") }}
{{- print .Values.configmapData.prometheus_export.port }}
{{- else }}
{{- print 0 }}
{{- end }}
{{- end }}
