package html

// TextAreaField is a Component that can be rendered using HTML templates
type TextAreaField struct {
	BaseComponent
	Class string
	Label string
	Name  string
	Rows  int
	Cols  int
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t TextAreaField) Render() string {
	return RenderComponent(t)
}
