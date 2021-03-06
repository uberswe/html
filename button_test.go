package html

import (
	"strings"
	"testing"
)

func TestButtonRender(t *testing.T) {
	expected := "<button class=\"test\">Test</button>"

	b := Button().Class("test").Content("Test")

	if b.Render() != expected {
		t.Errorf("HTML output (%s) does not match expected output: %s", b.Render(), expected)
	}
}

func TestButtonWithLinkRender(t *testing.T) {
	expected1 := "<button class=\"test2\">Test2</button>"
	expected2 := "<a href=\"https://beubo.com\">"

	b := Button().Class("test2").Content("Test2").Link("https://beubo.com")

	if !strings.Contains(b.Render(), expected1) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected1)
	}

	if !strings.Contains(b.Render(), expected2) {
		t.Errorf("HTML output (%s) does not contain expected output: %s", b.Render(), expected2)
	}
}
