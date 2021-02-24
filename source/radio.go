package source

var Radio = `{{ define "component.radio" }}
    <div class="{{ .ClassField }}">
        <input type="radio" id="{{ .IDField }}" name="{{ .NameField }}" value="{{ .ValueField }}"{{ if .Checked }} checked{{end}}>
        <label for="{{ .IDField }}">{{ .LabelField }}</label>
    </div>
{{ end }}`
