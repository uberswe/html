package source

var File = `{{ define "component.file" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IDField }}">{{ .LabelField }}</label>
        <input type="file" id="{{ .IDField }}" name="{{ .NameField }}" accept="{{ .AcceptField }}">
    </div>
{{ end }}`
