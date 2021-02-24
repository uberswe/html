package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// ButtonStruct is a Component that can be rendered using HTML templates
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

// Button initializes a ButtonStruct with default values
func Button() ButtonStruct {
	tmpl, _ := template.New("button").Parse(source.Button)
	return ButtonStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.button",
		},
	}
}

// Link makes a Button a link if provided
func (b ButtonStruct) Link(link string) ButtonStruct {
	b.LinkField = template.URL(link)
	return b
}

// Class sets a class to a button
func (b ButtonStruct) Class(class string) ButtonStruct {
	b.ClassField = class
	return b
}

// Content sets the content of a button
func (b ButtonStruct) Content(content string) ButtonStruct {
	b.ContentField = content
	return b
}
