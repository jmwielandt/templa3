package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hola mundo!")

	templateContent, err := os.ReadFile("template.gotmpl")
	if err != nil {
		log.Fatalf("couldn't read template file: %s", err)
	}

	tpl, err := template.New("template").Parse(string(templateContent))
	if err != nil {
		log.Fatalf("couldn't parse template file: %s", err)
	}

	templateVars, err := os.ReadFile("vars.json")
	if err != nil {
		log.Fatalf("couldn't read vars file: %s", err)
	}

	var data any
	err = json.Unmarshal(templateVars, &data)
	if err != nil {
		log.Fatalf("couldn't deserialize json file: %s", err)
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("couldn't execute template: %s", err)
	}
}
