package source

var Paragraph = `{{ define "component.paragraph" }}
    <div class="{{ .ClassField }}">
        {{ .ContentField }}
    </div>
{{ end }}`
