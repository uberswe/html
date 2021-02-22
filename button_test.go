package html

import (
	"html/template"
	"testing"
)

func TestButtonRender(t *testing.T) {
	expected := "<button class=\"test\">Test</button>"
	tmpl, err := template.ParseFiles("./html/button.html")

	if err != nil {
		t.Error(err)
	}

	b := Button{
		BaseComponent: BaseComponent{
			Class:          "test",
			Content:        "Test",
			TemplateString: "component.button",
			Template:       tmpl,
		},
	}

	if b.Render() != expected {
		t.Errorf("HTML output (%s) does not match expected output: %s", b.Render(), expected)
	}
}
