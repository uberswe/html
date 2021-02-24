package source

var Button = `{{ define "component.button" }}
    {{ if .LinkField }}<a href="{{ .LinkField }}">{{ end }}
    <button class="{{ .ClassField }}">{{ .ContentField }}</button>
    {{ if .LinkField }}</a>{{ end }}
{{ end }}`
