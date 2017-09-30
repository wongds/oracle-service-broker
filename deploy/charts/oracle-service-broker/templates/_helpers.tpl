{{/* vim: set filetype=mustache: */}}

{{/*
可以通过以下 DNS 域名访问服务.
我们限制字符长度为63字符因为一些Kuberntees属性只支持63字符以内的字符串.例如 DNS 域名解析.
*/}}
{{- define "fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
