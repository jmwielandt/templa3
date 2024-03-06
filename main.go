package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jmwielandt/templa3/templa3"
)

func parseCLIArgs() (tplt string, vars string, verbosePrints bool) {
	const (
		templateFlag = "template"
		varsFlag     = "vars"
	)

	templatePath := flag.String(templateFlag, "", "Path to the template file")
	varsPath := flag.String(varsFlag, "", "Path to the vars json file")
	verbose := flag.Bool("verbose", false, "Enables stdout prints after and during template execution")

	flag.Parse()

	if *templatePath == "" {
		slog.Error(fmt.Sprintf("%s can't be an empty string", templateFlag))
		return
	}
	if *varsPath == "" {
		slog.Error(fmt.Sprintf("%s can't be an empty string", varsFlag))
	}

	return *templatePath, *varsPath, *verbose
}

func initializeTemplate(verbose bool) *templa3.Template {
	tpl := templa3.NewTemplate(verbose)
	return tpl
}

func readFiles(templatePath, varsPath string) ([]byte, any, error) {
	// read files
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		err = fmt.Errorf("couldn't read template file: %w", err)
		return nil, nil, err
	}

	templateVars, err := os.ReadFile(varsPath)
	if err != nil {
		err = fmt.Errorf("couldn't read vars file: %w", err)
		return nil, nil, err
	}

	var data any
	err = json.Unmarshal(templateVars, &data)
	if err != nil {
		err = fmt.Errorf("couldn't deserialize json file: %w", err)
		return nil, nil, err
	}
	return templateContent, data, nil
}

func main() {
	templatePath, varsPath, verbose := parseCLIArgs()
	if verbose {
		fmt.Printf("template file: %s\n", templatePath)
		fmt.Printf("vars file: %s\n", varsPath)
	}

	templateContent, data, err := readFiles(templatePath, varsPath)
	if err != nil {
		log.Fatal(err)
	}

	tpl := initializeTemplate(verbose)

	err = tpl.Parse(string(templateContent))
	if err != nil {
		log.Fatalf("couldn't parse template file: %s", err)
	}

	templateResult, err := tpl.ExecuteString(data)
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		fmt.Println("---- TEMPLATE RESULT ----")
	}
	fmt.Println(templateResult)
}
