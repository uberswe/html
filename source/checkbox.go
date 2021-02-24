package source

var Checkbox = `{{ define "component.checkbox" }}
    <div class="{{ .ClassField }}">
        <input type="checkbox" id="{{ .IdField }}" name="{{ .NameField }}" value="{{ .ValueField }}"{{ if .Checked }} checked{{end}}>
        <label for="{{ .IdField }}">{{ .ContentField }}</label>
    </div>
{{ end }}`
