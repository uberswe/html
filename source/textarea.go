package source

var Textarea = `{{ define "component.textarea" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IDField }}">{{ .LabelField }}</label>
        <textarea id="{{ .IDField }}" name="{{ .NameField }}" rows="{{ .RowsCount }}" cols="{{ .ColsCount }}">
            {{ .ContentField }}
        </textarea>
    </div>
{{ end }}`
