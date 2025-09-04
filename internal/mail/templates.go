package mail

import (
	"bytes"
	"text/template"
)

func RenderTemplate(name string, data any) (string, error) {
    tmpl, err := template.ParseFiles("templates/" + name + ".tmpl")
    if err != nil {
        return "", err
    }
    var buf bytes.Buffer
    err = tmpl.Execute(&buf, data)
    return buf.String(), err
}
