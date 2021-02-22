package html

// CheckBoxField is a beubo component that can be rendered using HTML templates
type CheckBoxField struct {
	BaseComponent
	Name       string
	Identifier string
	Value      string
	Checked    bool
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (cb CheckBoxField) Render() string {
	return RenderComponent(cb)
}
