package source

var Hidden = `{{ define "component.hidden" }}
    <input type="hidden" id="{{ .IdField }}" name="{{ .NameField }}" value="{{ .ValueField }}">
{{ end }}`
