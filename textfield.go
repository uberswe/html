package html

// TextField is a beubo component that can be rendered using HTML templates
type TextField struct {
	BaseComponent
	Class       string
	Label       string
	Name        string
	Value       string
	Placeholder string
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t TextField) Render() string {
	return RenderComponent(t)
}
