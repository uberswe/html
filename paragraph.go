package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// ParagraphStruct is a Component that can be rendered using HTML templates
type ParagraphStruct struct {
	BaseComponent
	ClassField   string
	ContentField string
}

func Paragraph() ParagraphStruct {
	tmpl, _ := template.New("paragraph").Parse(source.Paragraph)
	return ParagraphStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.paragraph",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (p ParagraphStruct) Render() string {
	return RenderComponent(p)
}

func (p ParagraphStruct) Class(class string) ParagraphStruct {
	p.ClassField = class
	return p
}

func (p ParagraphStruct) Content(content string) ParagraphStruct {
	p.ContentField = content
	return p
}
