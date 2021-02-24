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
	IdField    string
	ValueField string
	Checked    bool
}

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

func (rf RadioStruct) Check() RadioStruct {
	rf.Checked = true
	return rf
}

func (rf RadioStruct) Uncheck() RadioStruct {
	rf.Checked = false
	return rf
}

func (rf RadioStruct) Class(class string) RadioStruct {
	rf.ClassField = class
	return rf
}

func (rf RadioStruct) Label(label string) RadioStruct {
	rf.LabelField = label
	return rf
}

func (rf RadioStruct) Name(name string) RadioStruct {
	rf.NameField = name
	return rf
}

func (rf RadioStruct) Id(id string) RadioStruct {
	rf.IdField = id
	return rf
}

func (rf RadioStruct) Value(value string) RadioStruct {
	rf.ValueField = value
	return rf
}
