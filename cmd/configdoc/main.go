// +build ignore

/*
Generates configuration markdown file from the config struct tags.
*/
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"

	"boilerplate/config"

	"github.com/jessevdk/go-flags"
)

type ConfigOpts struct {
	Output string `short:"o" long:"output" default:"docs/config.md"`
}

type ConfigStruct struct {
	reflect.Type
	Fields []reflect.StructField
}

func main() {
	opts := &ConfigOpts{}
	_, err := flags.Parse(opts)
	if err != nil {
		panic(err)
	}
	content, err := GenConfigMD(*config.Config)
	if err != nil {
		fmt.Println("Error generating MD", err)
	}
	fullContent := embedContents(opts.Output, content)
	writer, err := os.Create(opts.Output)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer writer.Close()
	writer.Write([]byte(fullContent))
}

func GenConfigMD(config interface{}) (string, error) {
	var buf bytes.Buffer
	cfg := reflect.TypeOf(config)
	if cfg.Kind().String() != "struct" {
		log.Fatalln("Non struct passed in", cfg.Kind())
	}
	_, fname, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("Could not determine path")
	}
	dir := path.Dir(fname)
	template, err := template.ParseFiles(path.Join(dir, "config_md.tmpl"))
	if err != nil {
		return "", err
	}
	tmplCfg := ConfigStruct{Type: cfg}
	for i := 0; i < cfg.NumField(); i++ {
		tmplCfg.Fields = append(tmplCfg.Fields, cfg.Field(i))
	}
	template.Execute(&buf, tmplCfg)
	return buf.String(), nil
}

var (
	embedStartRegex = regexp.MustCompile(
		`(?m:^ *)<!--\s*config:embed:start\s*-->(?s:.*?)<!--\s*config:embed:end\s*-->(?m:\s*?$)`,
	)
)

func embedContents(fileName string, text string) string {
	embedText := fmt.Sprintf("<!-- config:embed:start -->\n\n%s\n\n<!-- config:embed:end -->", text)

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("unable to find output file %s for embedding. Creating a new file instead", fileName)
		return embedText
	}

	var replacements int
	data = embedStartRegex.ReplaceAllFunc(data, func(_ []byte) []byte {
		replacements++
		return []byte(embedText)
	})

	if replacements == 0 {
		log.Printf("no embed markers found. Appending documentation to the end of the file instead")
		return fmt.Sprintf("%s\n\n%s", string(data), text)
	}

	return string(data)
}
