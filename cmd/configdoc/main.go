package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"

	"github.com/jessevdk/go-flags"
	"../../config"
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
	writer, err := os.Create(opts.Output)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer writer.Close()
	err = GenConfigMD(config.Cfg{}, writer)
	if err != nil {
		fmt.Println(err)
	}
}

func GenConfigMD(config interface{}, output io.Writer) error {
	cfg := reflect.TypeOf(config)
	if cfg.Kind().String() != "struct" {
		log.Fatalln(cfg.Kind())
	}
	_, fname, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("Could not determine path")
	}
	dir := path.Dir(fname)
	template, err := template.ParseFiles(path.Join(dir, "config_md.tmpl"))
	if err != nil {
		return err
	}
	tmplCfg := ConfigStruct{Type: cfg}
	for i := 0; i < cfg.NumField(); i++ {
		tmplCfg.Fields = append(tmplCfg.Fields, cfg.Field(i))
	}
	template.Execute(output, tmplCfg)
	return nil
}
