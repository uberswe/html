package source

var Form = `{{ define "component.form" }}
	<div class="{{ .ClassField }}">
    <form method="{{ .MethodField }}" action="{{ .ActionField }}">
        {{range .FieldsArray }}
            {{ $.RenderField .Render . }}
        {{ end }}
    </form>
    </div>
{{ end }}`
