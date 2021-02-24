package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// SelectStruct is a beubo component that can be rendered using HTML templates
type SelectStruct struct {
	BaseComponent
	ClassField   string
	LabelField   string
	IDField      string
	NameField    string
	OptionsArray []SelectOption
}

// Select creates a SelectStruct with default values
func Select() SelectStruct {
	tmpl, _ := template.New("select").Parse(source.Select)
	return SelectStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.select",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (s SelectStruct) Render() string {
	return RenderComponent(s)
}

// Options sets one or more options of the select field
func (s SelectStruct) Options(options ...SelectOption) SelectStruct {
	s.OptionsArray = options
	return s
}

// Class sets the class of the select field
func (s SelectStruct) Class(class string) SelectStruct {
	s.ClassField = class
	return s
}

// Label sets the label of the select field
func (s SelectStruct) Label(label string) SelectStruct {
	s.LabelField = label
	return s
}

// ID sets the id of the select field
func (s SelectStruct) ID(id string) SelectStruct {
	s.IDField = id
	return s
}

// Name sets the name of the select field
func (s SelectStruct) Name(name string) SelectStruct {
	s.NameField = name
	return s
}

// SelectOption is part of the SelectField values, there should be one or more of these
type SelectOption struct {
	ValueField string
	LabelField string
}

// Option creates a new SelectOption with default values
func Option() SelectOption {
	return SelectOption{}
}

// Value sets the value of the option
func (s SelectOption) Value(value string) SelectOption {
	s.ValueField = value
	return s
}

// Label sets the label of the option
func (s SelectOption) Label(label string) SelectOption {
	s.LabelField = label
	return s
}
