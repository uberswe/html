package html

import (
	"html/template"
)

// Table is a beubo component that can be rendered using HTML templates
type Table struct {
	BaseComponent
	Header           []Column
	Rows             []Row
	PageNumber       int // Current page
	PageDisplayCount int // How many rows per page
}

// Row represents a html table row which can have columns
type Row struct {
	Columns []Column
}

// Column represents a html column in a table which is part of a row
type Column struct {
	Name  string
	Field Component
	Value string
}

// Render calls RenderComponent to turn a Component into a html string for browser output
func (t Table) Render() string {
	return RenderComponent(t)
}

// RenderColumn calls RenderComponent to turn a Column into a html string which is added to the Table Render
func (t Table) RenderColumn(c Column) string {
	return RenderComponent(c.Field)
}

// RenderField calls Render to turn a Column into a string which is added to the Table Render
func (c Column) RenderField(value string, field Component) template.HTML {
	if field != nil && field.Render() != "" {
		return template.HTML(field.Render())
	}
	return template.HTML(value)
}
