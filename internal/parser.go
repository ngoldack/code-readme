package internal

import (
	"os"
	"text/template"
)

const DateFormat = "02.01.2006"

func Parse(cfg *Config, file []byte) error {

	tmpl := template.New("README")
	Info("Parsing template...")
	tmpl, err := tmpl.Parse(string(file))
	if err != nil {
		return err
	}

	Info("Generating template data...")
	data, err := generateTemplateData(cfg)
	if err != nil {
		return err
	}

	// write file
	Info("Writing finished README...")
	out, err := os.OpenFile("out/README.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	tmpl.Execute(out, data)

	return nil
}
