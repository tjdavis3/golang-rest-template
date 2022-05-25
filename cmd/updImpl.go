package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const funcBody = `
func (s *server) %s%s {
	err := buildProblemFromError(http.StatusNotImplemented, errNotImplemented, r)
	err.Render(w, r)
}
`

func main() {
	_, fname, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Could not determine path")
	}
	dir := path.Dir(fname)
	apiDir := path.Join(dir, "../api")
	pkgs, err := parser.ParseDir(token.NewFileSet(), apiDir, filter, 0)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fd, err := os.OpenFile(path.Join(apiDir, "serverImpl.go"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	for _, pkg := range pkgs {

		// Functions
		si, functions := findFuncs(pkg)
		for fn, sig := range si {
			if value, exist := functions[fn]; !exist {
				funcline := strings.Replace(sig, "func", "", 1)
				fmt.Println("Adding ", fn)
				_, err = fmt.Fprintf(fd, funcBody, fn, funcline)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Added function %s\n", fn)
				}
			} else {
				if value != sig {
					fmt.Printf("%s Signatures don't match\n", fn)
					fmt.Println("  Interface: ", sig)
					fmt.Println("  Function: ", value)
				}
			}
		}
	}
}

func findFuncs(pkg *ast.Package) (serverInterface map[string]string, functions map[string]string) {
	functions = make(map[string]string)
	serverInterface = make(map[string]string)
	for _, f := range pkg.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			switch t := n.(type) {
			// find variable declarations
			case *ast.TypeSpec:
				// which are public
				if t.Name.IsExported() {
					switch t.Type.(type) {
					// and are interfaces
					case *ast.InterfaceType:
						if t.Name.Name == "ServerInterface" {
							intf := t.Type.(*ast.InterfaceType)
							for _, m := range intf.Methods.List {
								for _, field := range m.Names {
									sig, err := getFuncSignature(field.Obj.Decl)
									if err == nil {
										serverInterface[field.Obj.Name] = sig
									}
								}
							}
						}
					}
				}
			case *ast.FuncDecl:
				sig, err := getFuncSignature(t)
				if err == nil {
					functions[t.Name.Name] = sig
				}
			}
			return true
		})
	}
	return serverInterface, functions
}

func getFuncSignature(intf interface{}) (signature string, err error) {
	switch intf.(type) {
	// and are interfaces
	case *ast.FuncDecl:
		found := false
		fnc := intf.(*ast.FuncDecl)
		if fnc.Recv != nil {
			for _, fld := range fnc.Recv.List {
				switch fld.Type.(type) {
				case *ast.StarExpr:
					expr := fld.Type.(*ast.StarExpr)
					if fmt.Sprintf("%s", expr.X) == "server" {
						found = true
						break
					}
				}
			}
		}
		var buf bytes.Buffer
		if found {
			body := fnc.Type
			printer.Fprint(&buf, token.NewFileSet(), body)
		} else {
			err = errors.New("Function not found")
		}
		return buf.String(), err

	case *ast.Field:
		fnc := intf.(*ast.Field)
		var buf bytes.Buffer
		body := fnc.Type
		printer.Fprint(&buf, token.NewFileSet(), body)
		return buf.String(), nil
	}
	return
}
func filter(info os.FileInfo) bool {

	name := info.Name()

	if info.IsDir() {
		return false
	}

	if filepath.Ext(name) != ".go" {
		return false
	}

	if strings.HasSuffix(name, "_test.go") {
		return false
	}

	return true

}
