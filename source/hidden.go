package source

var Hidden = `{{ define "component.hidden" }}
    <input type="hidden" id="{{ .IDField }}" name="{{ .NameField }}" value="{{ .ValueField }}">
{{ end }}`
