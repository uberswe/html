package html

// SelectField is a beubo component that can be rendered using HTML templates
type SelectField struct {
	BaseComponent
	Class   string
	Name    string
	Options []SelectFieldOption
}

// SelectFieldOption is part of the SelectField values, there should be one or more of these
type SelectFieldOption struct {
	Value   string
	Content string
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (sf SelectField) Render() string {
	return RenderComponent(sf)
}
