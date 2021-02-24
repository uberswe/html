package source

var Table = `{{ define "component.table" }}
    <table class="{{ .ClassField }}">
        <tr>
            {{range .HeaderArray }}
                <th>{{ .NameField }}</th>
            {{ end }}
        </tr>
        {{range .RowsArray }}
            <tr>
                {{range .ColumnsArray }}
                    <td>{{ .RenderField .ValueField .ComponentField }}</td>
                {{ end }}
            </tr>
        {{ end }}
    </table>
{{ end }}`
