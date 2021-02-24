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
	IdField      string
	ContentField string
	RowsCount    int
	ColsCount    int
}

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

func (t TextAreaStruct) Rows(rows int) TextAreaStruct {
	t.RowsCount = rows
	return t
}

func (t TextAreaStruct) Cols(cols int) TextAreaStruct {
	t.ColsCount = cols
	return t
}

func (t TextAreaStruct) Name(name string) TextAreaStruct {
	t.NameField = name
	return t
}

func (t TextAreaStruct) Id(id string) TextAreaStruct {
	t.IdField = id
	return t
}

func (t TextAreaStruct) Content(content string) TextAreaStruct {
	t.ContentField = content
	return t
}

func (t TextAreaStruct) Label(label string) TextAreaStruct {
	t.LabelField = label
	return t
}

func (t TextAreaStruct) Class(class string) TextAreaStruct {
	t.ClassField = class
	return t
}
