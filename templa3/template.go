package templa3

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

type Template struct {
	verbose     bool
	template    *template.Template
	customFuncs template.FuncMap
}

func NewTemplate(verbose bool) *Template {
	tpl := &Template{
		verbose:     verbose,
		template:    template.New("template"),
		customFuncs: template.FuncMap{},
	}
	tpl.registerFunctions()
	tpl.template = tpl.template.Funcs(tpl.customFuncs)
	return tpl
}

func (c *Template) Parse(text string) error {
	var err error
	c.template, err = c.template.Parse(text)
	return err
}

func (c *Template) Execute(data any) (io.Reader, error) {
	buf := []byte{}
	wr := bytes.NewBuffer(buf)
	err := c.template.Execute(wr, data)
	return wr, err
}

func (c *Template) ExecuteBuf(wr io.Writer, data any) error {
	err := c.template.Execute(wr, data)
	return err
}

func (c *Template) ExecuteString(data any) (string, error) {
	buf := []byte{}
	wr := bytes.NewBuffer(buf)
	err := c.template.Execute(wr, data)
	if err != nil {
		err = fmt.Errorf("couldn't execute template: %s", err)
		return "", err
	}
	templateResult, err := io.ReadAll(wr)
	if err != nil {
		err = fmt.Errorf("couldn't read template result: %s", err)
		return "", err
	}
	return string(templateResult), nil
}
