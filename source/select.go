package source

var Select = `{{ define "component.select" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IDField }}">{{ .LabelField }}</label>
        <select name="{{ .NameField }}" id="{{ .IDField }}">
            {{range .OptionsArray }}<option value="{{ .ValueField }}">{{ .LabelField }}</option>
            {{ end }}
        </select>
    </div>
{{ end }}`
