package tal

import (
	"bytes"
	"strings"
	"testing"
)

func TestTalesDeepPaths(t *testing.T) {
	type cT struct {
		C map[string]string
		D interface{}
		N interface{}
	}
	type aT struct {
		B map[string]cT
	}
	c := cT{
		C: make(map[string]string),
		D: Default,
		N: None,
	}
	c.C["one"] = "two"
	a := aT{
		B: make(map[string]cT),
	}
	a.B["alpha"] = c

	runTalesTest(t, talesTest{
		struct{ A aT }{A: a},
		`<html><body><h1 tal:content="a/b/alpha/C/one">Default header</h1><h2 tal:content="a/b/alpha/D">Default header 2</h2><h3 tal:content="a/b/alpha/N">Default header 3</h3></body></html>`,
		`<html><body><h1>two</h1><h2>Default header 2</h2><h3></h3></body></html>`,
	})
}

type talesTest struct {
	Context  interface{}
	Template string
	Expected string
}

func runTalesTest(t *testing.T, test talesTest, cfg ...RenderConfig) {
	temp, err := CompileTemplate(strings.NewReader(test.Template))
	if err != nil {
		t.Errorf("Error compiling template: %v\n", err)
		return
	}

	resultBuffer := &bytes.Buffer{}
	err = temp.Render(test.Context, resultBuffer, cfg...)

	if err != nil {
		t.Errorf("Error rendering template: %v\n", err)
		return
	}

	resultStr := resultBuffer.String()

	if resultStr != test.Expected {
		t.Errorf("Expected output: \n%v\nActual output: \n%v\nFrom template: \n%v\nCompiled into: \n%v\n", test.Expected, resultStr, test.Template, temp.String())
		return
	}
}
