package source

var Textarea = `{{ define "component.textarea" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IdField }}">{{ .LabelField }}</label>
        <textarea id="{{ .IdField }}" name="{{ .NameField }}" rows="{{ .RowsCount }}" cols="{{ .ColsCount }}">
            {{ .ContentField }}
        </textarea>
    </div>
{{ end }}`
