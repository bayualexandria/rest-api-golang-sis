package template

import (
    "html/template"
    "io/ioutil"
    "bytes"
)

func LoadHTMLTemplate(path string, data interface{}) (string, error) {
    fileData, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err
    }

    tmpl, err := template.New("email").Parse(string(fileData))
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, data); err != nil {
        return "", err
    }

    return buf.String(), nil
}