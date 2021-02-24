package source

var Radio = `{{ define "component.radio" }}
    <div class="{{ .ClassField }}">
        <input type="radio" id="{{ .IdField }}" name="{{ .NameField }}" value="{{ .ValueField }}"{{ if .Checked }} checked{{end}}>
        <label for="{{ .IdField }}">{{ .LabelField }}</label>
    </div>
{{ end }}`
