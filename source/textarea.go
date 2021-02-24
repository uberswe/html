package source

var Textarea = `{{ define "component.textarea" }}
    <div class="{{ .Class }}">
        <label for="{{ .Identifier }}">{{ .Label }}</label>

        <textarea id="{{ .Identifier }}" name="{{ .Name }}" rows="{{ .Rows }}" cols="{{ .Cols }}">
            {{ .Content }}
        </textarea>
    </div>
{{ end }}`
