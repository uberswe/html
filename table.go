package html

import (
	"github.com/uberswe/html/source"
	"html/template"
)

// TableStruct is a beubo component that can be rendered using HTML templates
type TableStruct struct {
	BaseComponent
	ClassField  string
	HeaderArray []ColumnStruct
	RowsArray   []RowStruct
}

func Table() TableStruct {
	tmpl, _ := template.New("table").Parse(source.Table)
	return TableStruct{
		BaseComponent: BaseComponent{
			Template:       tmpl,
			TemplateString: "component.table",
		},
	}
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t TableStruct) Render() string {
	return RenderComponent(t)
}

func (t TableStruct) Header(columns ...ColumnStruct) TableStruct {
	t.HeaderArray = columns
	return t
}

func (t TableStruct) Rows(rows ...RowStruct) TableStruct {
	t.RowsArray = rows
	return t
}

// RenderColumn calls RenderComponent to turn a ColumnStruct into a html string which is added to the TableStruct Render
func (t TableStruct) RenderColumn(c ColumnStruct) string {
	return RenderComponent(c.ComponentField)
}

func (t TableStruct) Class(class string) TableStruct {
	t.ClassField = class
	return t
}

// RowStruct represents a html table row which can have columns
type RowStruct struct {
	ColumnsArray []ColumnStruct
}

func Row() RowStruct {
	return RowStruct{}
}

func (r RowStruct) Columns(columns ...ColumnStruct) RowStruct {
	r.ColumnsArray = columns
	return r
}

// ColumnStruct represents a html column in a table which is part of a row
type ColumnStruct struct {
	NameField      string
	ComponentField Component
	// Value is used when there is no Field
	ValueField string
}

func Column() ColumnStruct {
	return ColumnStruct{}
}

// RenderField calls Render to turn a ColumnStruct into a string which is added to the TableStruct Render
func (c ColumnStruct) RenderField(value string, field Component) template.HTML {
	if field != nil && field.Render() != "" {
		return template.HTML(field.Render())
	}
	return template.HTML(value)
}

func (c ColumnStruct) Name(name string) ColumnStruct {
	c.NameField = name
	return c
}

func (c ColumnStruct) Value(value string) ColumnStruct {
	c.ValueField = value
	return c
}

func (c ColumnStruct) Field(component Component) ColumnStruct {
	c.ComponentField = component
	return c
}
