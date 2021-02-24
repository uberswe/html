package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// CheckBoxStruct is a beubo component that can be rendered using HTML templates
type CheckBoxStruct struct {
	BaseComponent
	ClassField   string
	ContentField string
	NameField    string
	IdField      string
	ValueField   string
	Checked      bool
}

func Checkbox() CheckBoxStruct {
	tmpl, _ := template.New("checkbox").Parse(source.Checkbox)
	return CheckBoxStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.checkbox",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (c CheckBoxStruct) Render() string {
	return RenderComponent(c)
}

func (c CheckBoxStruct) Check() CheckBoxStruct {
	c.Checked = true
	return c
}

func (c CheckBoxStruct) Uncheck() CheckBoxStruct {
	c.Checked = false
	return c
}

func (c CheckBoxStruct) Class(class string) CheckBoxStruct {
	c.ClassField = class
	return c
}

func (c CheckBoxStruct) Content(content string) CheckBoxStruct {
	c.ContentField = content
	return c
}

func (c CheckBoxStruct) Name(name string) CheckBoxStruct {
	c.NameField = name
	return c
}

func (c CheckBoxStruct) Id(id string) CheckBoxStruct {
	c.IdField = id
	return c
}

func (c CheckBoxStruct) Value(value string) CheckBoxStruct {
	c.ValueField = value
	return c
}
