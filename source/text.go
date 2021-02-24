package source

var Text = `{{ define "component.text" }}
    <div class="{{ .ClassField }}">
        <label for="{{ .IDField }}">{{ .LabelField }}</label>
        <input type="{{ .TypeField }}" id="{{ .IDField }}" name="{{ .NameField }}" value="{{ .ValueField }}" placeholder="{{ .PlaceholderField }}">
    </div>
{{ end }}`
