package html

// Text is a Component that can be rendered using HTML templates
type Text struct {
	BaseComponent
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t Text) Render() string {
	return RenderComponent(t)
}
