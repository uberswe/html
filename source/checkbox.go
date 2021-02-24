package source

var Checkbox = `{{ define "component.checkbox" }}
    <div class="{{ .ClassField }}">
        <input type="checkbox" id="{{ .IDField }}" name="{{ .NameField }}" value="{{ .ValueField }}"{{ if .Checked }} checked{{end}}>
        <label for="{{ .IDField }}">{{ .LabelField }}</label>
    </div>
{{ end }}`
