package html

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

// Component is anything that can be rendered on a page. A text field is a component but so is the form the text field is a part of.
type Component interface {
	Render() string
	GetTemplateName() string
	GetTemplate() *template.Template
}

// RenderComponent takes the provided component and finds the relevant template and renders this into a string
func RenderComponent(c Component) string {
	var foundTemplate *template.Template
	if c.GetTemplate() == nil {
		return ""
	}
	if foundTemplate = c.GetTemplate().Lookup(c.GetTemplateName()); foundTemplate == nil {
		log.Printf("Component template not found %s\n", c.GetTemplateName())
		return ""
	}
	buf := &bytes.Buffer{}
	err := foundTemplate.Execute(buf, c)
	if err != nil {
		log.Printf("Component error executing template %s\n", c.GetTemplateName())
		return ""
	}
	return strings.TrimSpace(buf.String())
}

// BaseComponent defines the fields all components must have
type BaseComponent struct {
	TemplateString string
	Template       *template.Template
}

// GetTemplate is a getter for the T Property
func (b BaseComponent) GetTemplate() *template.Template {
	return b.Template
}

// GetTemplateName is a getter for the Template Property
func (b BaseComponent) GetTemplateName() string {
	return b.TemplateString
}

// SetTemplate sets the template used for rendering the html
func (b BaseComponent) SetTemplate(template *template.Template) BaseComponent {
	b.Template = template
	return b
}

// SetTemplateName sets the name of the template for generating the component
func (b BaseComponent) SetTemplateName(template string) BaseComponent {
	b.TemplateString = template
	return b
}
