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
	IDField          string
	ValueField       string
	PlaceholderField string
	LabelField       string
	TypeField        string
}

// Text creates a TextStruct with default values
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

// Class sets the class of the text field
func (t TextStruct) Class(class string) TextStruct {
	t.ClassField = class
	return t
}

// Name sets the name of the text field
func (t TextStruct) Name(name string) TextStruct {
	t.NameField = name
	return t
}

// ID sets the id of the text struct
func (t TextStruct) ID(id string) TextStruct {
	t.IDField = id
	return t
}

// Value sets the value of the text struct
func (t TextStruct) Value(value string) TextStruct {
	t.ValueField = value
	return t
}

// Placeholder sets the placeholder of the text field
func (t TextStruct) Placeholder(placeholder string) TextStruct {
	t.PlaceholderField = placeholder
	return t
}

// Label sets the label of the text field
func (t TextStruct) Label(label string) TextStruct {
	t.LabelField = label
	return t
}

// Type sets the type of the text field
func (t TextStruct) Type(typo string) TextStruct {
	t.TypeField = typo
	return t
}
