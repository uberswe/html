package html

import (
	"testing"
)

func TestHiddenRender(t *testing.T) {
	expected := "<input type=\"hidden\" id=\"test\" name=\"test\" value=\"test\">"

	h := Hidden().Id("test").Value("test").Name("test")

	if h.Render() != expected {
		t.Errorf("HTML output (%s) does not match expected output: %s", h.Render(), expected)
	}
}
