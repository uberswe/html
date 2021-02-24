package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// FormStruct is a Component that can be rendered using HTML templates
type FormStruct struct {
	BaseComponent
	ClassField  string
	MethodField string
	ActionField string
	FieldsArray []Component
}

// Form creates a FormStruct with default values
func Form() FormStruct {
	tmpl, _ := template.New("form").Parse(source.Form)
	return FormStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.form",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (f FormStruct) Render() string {
	return RenderComponent(f)
}

// RenderField calls Render to turn a ColumnStruct into a string which is added to the FormStruct Render
func (f FormStruct) RenderField(value string, field Component) template.HTML {
	if field != nil && field.Render() != "" {
		return template.HTML(field.Render())
	}
	return template.HTML(value)
}

// Fields sets the fields of the form, these are components like buttons, checkboxes, and more
func (f FormStruct) Fields(components ...Component) FormStruct {
	f.FieldsArray = components
	return f
}

// Class sets the class of the form
func (f FormStruct) Class(class string) FormStruct {
	f.ClassField = class
	return f
}

// Method sets the method of the form
func (f FormStruct) Method(method string) FormStruct {
	f.MethodField = method
	return f
}

// Action sets the action of the form
func (f FormStruct) Action(action string) FormStruct {
	f.ActionField = action
	return f
}
