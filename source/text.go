package source

var Text = `{{ define "component.text" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IdField }}">{{ .LabelField }}</label>
        <input type="{{ .TypeField }}" id="{{ .IdField }}" name="{{ .NameField }}" value="{{ .ValueField }}" placeholder="{{ .PlaceholderField }}">
    </div>
{{ end }}`
