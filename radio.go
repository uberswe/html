package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// RadioStruct is a Component that can be rendered using HTML templates
type RadioStruct struct {
	BaseComponent
	ClassField string
	LabelField string
	NameField  string
	IDField    string
	ValueField string
	Checked    bool
}

// Radio creates a new RadioStruct with default values
func Radio() RadioStruct {
	tmpl, _ := template.New("radio").Parse(source.Radio)
	return RadioStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.radio",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (rf RadioStruct) Render() string {
	return RenderComponent(rf)
}

// Check sets the checked value of the radio field to true
func (rf RadioStruct) Check() RadioStruct {
	rf.Checked = true
	return rf
}

// Uncheck sets the checked value of the radio field to false
func (rf RadioStruct) Uncheck() RadioStruct {
	rf.Checked = false
	return rf
}

// Class sets the class of the radio field
func (rf RadioStruct) Class(class string) RadioStruct {
	rf.ClassField = class
	return rf
}

// Label sets the label of the radio field
func (rf RadioStruct) Label(label string) RadioStruct {
	rf.LabelField = label
	return rf
}

// Name sets the name of the radio field
func (rf RadioStruct) Name(name string) RadioStruct {
	rf.NameField = name
	return rf
}

// ID sets the id of the radio field
func (rf RadioStruct) ID(id string) RadioStruct {
	rf.IDField = id
	return rf
}

// Value sets the value of the radio field
func (rf RadioStruct) Value(value string) RadioStruct {
	rf.ValueField = value
	return rf
}
