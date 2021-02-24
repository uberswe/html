package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// TextStruct is a beubo component that can be rendered using HTML templates
type TextStruct struct {
	BaseComponent
	ClassField       string
	NameField        string
	IdField          string
	ValueField       string
	PlaceholderField string
	LabelField       string
	TypeField        string
}

func Text() TextStruct {
	tmpl, _ := template.New("text").Parse(source.Text)
	return TextStruct{
		TypeField: "text",
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.text",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t TextStruct) Render() string {
	return RenderComponent(t)
}

func (t TextStruct) Class(class string) TextStruct {
	t.ClassField = class
	return t
}

func (t TextStruct) Name(name string) TextStruct {
	t.NameField = name
	return t
}

func (t TextStruct) Id(id string) TextStruct {
	t.IdField = id
	return t
}

func (t TextStruct) Value(value string) TextStruct {
	t.ValueField = value
	return t
}

func (t TextStruct) Placeholdler(placeholder string) TextStruct {
	t.PlaceholderField = placeholder
	return t
}

func (t TextStruct) Label(label string) TextStruct {
	t.LabelField = label
	return t
}

func (t TextStruct) Type(typo string) TextStruct {
	t.TypeField = typo
	return t
}
