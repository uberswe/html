package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// Button is a Component that can be rendered using HTML templates
type ButtonStruct struct {
	BaseComponent
	ClassField   string
	ContentField string
	LinkField    template.URL
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (b ButtonStruct) Render() string {
	return RenderComponent(b)
}

func Button() ButtonStruct {
	tmpl, _ := template.New("button").Parse(source.Button)
	return ButtonStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.button",
		},
	}
}

func (b ButtonStruct) Link(link string) ButtonStruct {
	b.LinkField = template.URL(link)
	return b
}

func (b ButtonStruct) Class(class string) ButtonStruct {
	b.ClassField = class
	return b
}

func (b ButtonStruct) Content(content string) ButtonStruct {
	b.ContentField = content
	return b
}
