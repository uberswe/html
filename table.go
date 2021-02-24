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

// Table creates a new TableStruct with default values
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

// Header sets one or more header columns for the table
func (t TableStruct) Header(columns ...ColumnStruct) TableStruct {
	t.HeaderArray = columns
	return t
}

// Rows sets one or more rows for the table
func (t TableStruct) Rows(rows ...RowStruct) TableStruct {
	t.RowsArray = rows
	return t
}

// RenderColumn calls RenderComponent to turn a ColumnStruct into a html string which is added to the TableStruct Render
func (t TableStruct) RenderColumn(c ColumnStruct) string {
	return RenderComponent(c.ComponentField)
}

// Class sets the class of the table
func (t TableStruct) Class(class string) TableStruct {
	t.ClassField = class
	return t
}

// RowStruct represents a html table row which can have columns
type RowStruct struct {
	ColumnsArray []ColumnStruct
}

// Row creates a RowStruct with default values
func Row() RowStruct {
	return RowStruct{}
}

// Columns sets one or more columns to a row
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

// Column creates a new ColumnStruct with default values
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

// Name sets the name of the column used for headers only
func (c ColumnStruct) Name(name string) ColumnStruct {
	c.NameField = name
	return c
}

// Value sets the value of the column used for rows only and only used if there is no field
func (c ColumnStruct) Value(value string) ColumnStruct {
	c.ValueField = value
	return c
}

// Field sets the field of the column and will override any value
func (c ColumnStruct) Field(component Component) ColumnStruct {
	c.ComponentField = component
	return c
}
