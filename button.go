package html

import "html/template"

// Button is a Component that can be rendered using HTML templates
type Button struct {
	BaseComponent
	Link template.URL
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (b Button) Render() string {
	return RenderComponent(b)
}
