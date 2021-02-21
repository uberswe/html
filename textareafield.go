package html

import (
	"html/template"
)

// TextAreaField is a beubo component that can be rendered using HTML templates
type TextAreaField struct {
	Content    string
	Theme      string
	Template   string
	Class      string
	Identifier string
	Label      string
	Name       string
	Rows       int
	Cols       int
	T          *template.Template
}

// GetSection is a getter for the Section property
func (t TextAreaField) GetSection() string {
	return ""
}

// GetTemplateName is a getter for the Template property
func (t TextAreaField) GetTemplateName() string {
	return returnTIfNotEmpty(t.Template, "component.textareafield")
}

// GetTheme is a getter for the Theme property
func (t TextAreaField) GetTheme() string {
	return returnTIfNotEmpty(t.Theme, "default")
}

// GetTemplate is a getter for the T Property
func (t TextAreaField) GetTemplate() *template.Template {
	return t.T
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t TextAreaField) Render() string {
	return RenderComponent(t)
}
