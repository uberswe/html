package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// HiddenStruct is a beubo component that can be rendered using HTML templates
type HiddenStruct struct {
	BaseComponent
	IDField    string
	NameField  string
	ValueField string
}

// Hidden creates a new HiddenStruct with default values
func Hidden() HiddenStruct {
	tmpl, _ := template.New("hidden").Parse(source.Hidden)
	return HiddenStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.hidden",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (hf HiddenStruct) Render() string {
	return RenderComponent(hf)
}

// ID sets the id of the hidden field
func (hf HiddenStruct) ID(id string) HiddenStruct {
	hf.IDField = id
	return hf
}

// Name sets the name of the hidden field
func (hf HiddenStruct) Name(name string) HiddenStruct {
	hf.NameField = name
	return hf
}

// Value sets the value of the hidden field
func (hf HiddenStruct) Value(value string) HiddenStruct {
	hf.ValueField = value
	return hf
}
