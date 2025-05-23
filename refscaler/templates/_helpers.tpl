{{/*
Expand the name of the chart.
*/}}
{{- define "refscaler.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "refscaler.fullname" -}}
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
Create chart name and version as used by the chart label.
*/}}
{{- define "refscaler.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "refscaler.labels" -}}
helm.sh/chart: {{ include "refscaler.chart" . }}
{{ include "refscaler.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "refscaler.selectorLabels" -}}
app.kubernetes.io/name: {{ include "refscaler.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create a full API path by combining hostname and API prefix.
*/}}
{{- define "refscaler.apiPath" -}}
{{- $hostname := .Values.gateway.hostname -}}
{{- $apiPrefix := .Values.gateway.apiPathPrefix | default "/api" -}}
{{- if not (hasPrefix "/" $apiPrefix) -}}
{{- $apiPrefix = printf "/%s" $apiPrefix -}}
{{- end -}}
{{- printf "https://%s%s" $hostname $apiPrefix -}}
{{- end -}}

{{/*
Create the backend service URL for the frontend to use
*/}}
{{- define "refscaler.backendServiceUrl" -}}
{{- $port := .Values.backend.service.port | int -}}
{{- printf "http://%s-%s:%d" (include "refscaler.fullname" .) "backend" $port -}}
{{- end -}}

{{/*
Create annotations that force a restart when upgraded
*/}}
{{- define "refscaler.restartAnnotations" -}}
kubernetes.io/restartedAt: {{ now | quote }}
{{- end }}
