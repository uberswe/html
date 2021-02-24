package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// HiddenStruct is a beubo component that can be rendered using HTML templates
type HiddenStruct struct {
	BaseComponent
	IdField    string
	NameField  string
	ValueField string
}

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

func (hf HiddenStruct) Id(id string) HiddenStruct {
	hf.IdField = id
	return hf
}

func (hf HiddenStruct) Name(name string) HiddenStruct {
	hf.NameField = name
	return hf
}

func (hf HiddenStruct) Value(value string) HiddenStruct {
	hf.ValueField = value
	return hf
}
