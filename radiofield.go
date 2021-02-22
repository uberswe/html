package html

// RadioField is a Component that can be rendered using HTML templates
type RadioField struct {
	BaseComponent
	Name    string
	Value   string
	Checked bool
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (rf RadioField) Render() string {
	return RenderComponent(rf)
}
