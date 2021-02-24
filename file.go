package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

type FileStruct struct {
	BaseComponent
	ClassField   string
	ContentField string
	NameField    string
	IdField      string
	AcceptField  string
}

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

func (f FileStruct) Class(class string) FileStruct {
	f.ClassField = class
	return f
}

func (f FileStruct) Content(content string) FileStruct {
	f.ContentField = content
	return f
}

func (f FileStruct) Name(name string) FileStruct {
	f.NameField = name
	return f
}

func (f FileStruct) Id(id string) FileStruct {
	f.IdField = id
	return f
}

func (f FileStruct) Accept(accept string) FileStruct {
	f.AcceptField = accept
	return f
}
