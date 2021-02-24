package source

var File = `{{ define "component.file" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IdField }}">{{ .ContentField }}</label>
        <input type="file" id="{{ .IdField }}" name="{{ .NameField }}" accept="{{ .AcceptField }}">
    </div>
{{ end }}`
