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

func (f FormStruct) Fields(components ...Component) FormStruct {
	f.FieldsArray = components
	return f
}

func (f FormStruct) Class(class string) FormStruct {
	f.ClassField = class
	return f
}

func (f FormStruct) Method(method string) FormStruct {
	f.MethodField = method
	return f
}

func (f FormStruct) Action(action string) FormStruct {
	f.ActionField = action
	return f
}
