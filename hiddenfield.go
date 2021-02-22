package html

// HiddenField is a beubo component that can be rendered using HTML templates
type HiddenField struct {
	BaseComponent
	Name  string
	Value string
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (hf HiddenField) Render() string {
	return RenderComponent(hf)
}
