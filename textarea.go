package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// TextAreaStruct is a Component that can be rendered using HTML templates
type TextAreaStruct struct {
	BaseComponent
	ClassField   string
	NameField    string
	LabelField   string
	IDField      string
	ContentField string
	RowsCount    int
	ColsCount    int
}

// TextArea creates a Textareastruct with default values
func TextArea() TextAreaStruct {
	tmpl, _ := template.New("textarea").Parse(source.Textarea)
	return TextAreaStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.textarea",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t TextAreaStruct) Render() string {
	return RenderComponent(t)
}

// Rows sets the rows attribute of the textarea field
func (t TextAreaStruct) Rows(rows int) TextAreaStruct {
	t.RowsCount = rows
	return t
}

// Cols sets the cols attribute of the textarea field
func (t TextAreaStruct) Cols(cols int) TextAreaStruct {
	t.ColsCount = cols
	return t
}

// Name sets the name of the textarea field
func (t TextAreaStruct) Name(name string) TextAreaStruct {
	t.NameField = name
	return t
}

// ID sets the id of the text area field
func (t TextAreaStruct) ID(id string) TextAreaStruct {
	t.IDField = id
	return t
}

// Content sets the content of the text area field
func (t TextAreaStruct) Content(content string) TextAreaStruct {
	t.ContentField = content
	return t
}

// Label sets the label of the text area field
func (t TextAreaStruct) Label(label string) TextAreaStruct {
	t.LabelField = label
	return t
}

// Class sets the class of the text area field
func (t TextAreaStruct) Class(class string) TextAreaStruct {
	t.ClassField = class
	return t
}
