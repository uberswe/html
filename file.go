package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// FileStruct represents a html file field
type FileStruct struct {
	BaseComponent
	ClassField  string
	LabelField  string
	NameField   string
	IDField     string
	AcceptField string
}

// File creates a FileStruct with default values
func File() FileStruct {
	tmpl, _ := template.New("file").Parse(source.File)
	return FileStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.file",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (f FileStruct) Render() string {
	return RenderComponent(f)
}

// Class sets the class of the file field
func (f FileStruct) Class(class string) FileStruct {
	f.ClassField = class
	return f
}

// Label sets the label of the file field
func (f FileStruct) Label(label string) FileStruct {
	f.LabelField = label
	return f
}

// Name sets the name of the file field
func (f FileStruct) Name(name string) FileStruct {
	f.NameField = name
	return f
}

// ID sets the id of the file field
func (f FileStruct) ID(id string) FileStruct {
	f.IDField = id
	return f
}

// Accept sets the accept attribute of the file field
func (f FileStruct) Accept(accept string) FileStruct {
	f.AcceptField = accept
	return f
}
