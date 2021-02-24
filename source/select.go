package source

var Select = `{{ define "component.select" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IdField }}">{{ .LabelField }}</label>
        <select name="{{ .NameField }}" id="{{ .IdField }}">
            {{range .OptionsArray }}<option value="{{ .ValueField }}">{{ .LabelField }}</option>
            {{ end }}
        </select>
    </div>
{{ end }}`
