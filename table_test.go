package html

import (
	"strings"
	"testing"
)

func TestTableRender(t *testing.T) {

	// <table class="test">
	//                <tr>
	//
	//                        <th>Header 1</th>
	//
	//                        <th>Header 2</th>
	//
	//                        <th>Header 3</th>
	//
	//                </tr>
	//
	//                    <tr>
	//
	//                            <td>Test</td>
	//
	//                            <td><button class="">Click me</button></td>
	//
	//                    </tr>
	//
	//                    <tr>
	//
	//                            <td>Test 432</td>
	//
	//                            <td>Test 123</td>
	//
	//                    </tr>
	//
	//                    <tr>
	//
	//                            <td>Testing Testing</td>
	//
	//                            <td>Testing 123</td>
	//
	//                    </tr>
	//
	//            </table>
	expected1 := "<table class=\"test\">"
	expected2 := "<tr>"
	expected3 := "<th>Header 1</th>"
	expected4 := "<th>Header 2</th>"
	expected5 := "<th>Header 3</th>"
	expected6 := "</tr>"
	expected7 := "<td>Test</td>"
	expected8 := "<td><button class=\"\">Click me</button></td>"
	expected9 := "<td>Test 432</td>"
	expected10 := "<td>Test 123</td>"
	expected11 := "<td>Testing Testing</td>"
	expected12 := "<td>Testing 123</td>"
	expected13 := "</table>"

	header1 := Column().Name("Header 1")
	header2 := Column().Name("Header 2")
	header3 := Column().Name("Header 3")

	row1 := Row().Columns(Column().Value("Test"), Column().Field(Button().Content("Click me")))
	row2 := Row().Columns(Column().Value("Test 432"), Column().Value("Test 123"))
	row3 := Row().Columns(Column().Value("Testing Testing"), Column().Value("Testing 123"))

	f := Table().Class("test").Header(header1, header2, header3).Rows(row1, row2, row3)

	if !strings.Contains(f.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected1)
	}

	if !strings.Contains(f.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected2)
	}

	if !strings.Contains(f.Render(), expected3) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected3)
	}

	if !strings.Contains(f.Render(), expected4) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected4)
	}

	if !strings.Contains(f.Render(), expected5) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected5)
	}

	if !strings.Contains(f.Render(), expected6) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected6)
	}

	if !strings.Contains(f.Render(), expected7) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected7)
	}

	if !strings.Contains(f.Render(), expected8) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected8)
	}

	if !strings.Contains(f.Render(), expected9) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected9)
	}

	if !strings.Contains(f.Render(), expected10) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected10)
	}

	if !strings.Contains(f.Render(), expected11) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected11)
	}

	if !strings.Contains(f.Render(), expected12) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected12)
	}

	if !strings.Contains(f.Render(), expected13) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", f.Render(), expected13)
	}
}
