package html

import (
	"html/template"
)

// Form is a Component that can be rendered using HTML templates
type Form struct {
	BaseComponent
	Fields []Component
	Method string
	Action string
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (f Form) Render() string {
	return RenderComponent(f)
}

// RenderField calls Render to turn a Column into a string which is added to the Form Render
func (f Form) RenderField(value string, field Component) template.HTML {
	if field != nil && field.Render() != "" {
		return template.HTML(field.Render())
	}
	return template.HTML(value)
}
