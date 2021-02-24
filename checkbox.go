package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// CheckBoxStruct is a beubo component that can be rendered using HTML templates
type CheckBoxStruct struct {
	BaseComponent
	ClassField string
	LabelField string
	NameField  string
	IDField    string
	ValueField string
	Checked    bool
}

// Checkbox creates a CheckBoxStruct with default values
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

// Check sets a checkbox checked to true
func (c CheckBoxStruct) Check() CheckBoxStruct {
	c.Checked = true
	return c
}

// Uncheck sets a checkbox checked to false
func (c CheckBoxStruct) Uncheck() CheckBoxStruct {
	c.Checked = false
	return c
}

// Class sets the class of the checkbox
func (c CheckBoxStruct) Class(class string) CheckBoxStruct {
	c.ClassField = class
	return c
}

// Label sets the label of the checkbox
func (c CheckBoxStruct) Label(label string) CheckBoxStruct {
	c.LabelField = label
	return c
}

// Name sets the name of the checkbox
func (c CheckBoxStruct) Name(name string) CheckBoxStruct {
	c.NameField = name
	return c
}

// ID sets the id of the checkbox
func (c CheckBoxStruct) ID(id string) CheckBoxStruct {
	c.IDField = id
	return c
}

// Value sets the value of the checkbox
func (c CheckBoxStruct) Value(value string) CheckBoxStruct {
	c.ValueField = value
	return c
}
