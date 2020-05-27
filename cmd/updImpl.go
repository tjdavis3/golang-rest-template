package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

const implFname = "serverImpl.go"
const genFname = "api.gen.go"
const funcSignature = `func (s *server) %s(w http.ResponseWriter, r *http.Request) {
  http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
`

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid data-type: %s", arr.Kind()))
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func loadImpl(filename string) []string {
	re := regexp.MustCompile(`\Afunc \(s \*server\) ([a-zA-Z0-9]+).*`)

	var functions []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if re.MatchString(line) {
			functions = append(functions, re.ReplaceAllString(line, "$1"))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return functions
}

func loadInterface(filename string) []string {
	var functions []string

	ifRE := regexp.MustCompile(`\Atype ServerInterface interface {`)
	endRE := regexp.MustCompile(`.*}.*`)
	funcRE := regexp.MustCompile(`^[[:blank:]]*([a-zA-Z0-9]+).*`)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	done := false
	for scanner.Scan() {
		if done {
			break
		}
		line := scanner.Text()
		if ifRE.MatchString(line) {
			for scanner.Scan() {
				line = strings.TrimSpace(scanner.Text())
                                if strings.HasPrefix(line, "//") {
                                        continue
				}
				if endRE.MatchString(line) {
					break
				}
				if funcRE.MatchString(line) {
					functions = append(functions, funcRE.ReplaceAllString(line, "$1"))
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return functions
}
func main() {
	functions := loadImpl(implFname)
	intFuncs := loadInterface(genFname)
	fd, err := os.OpenFile(implFname, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	for _, intFunc := range intFuncs {
		if !itemExists(functions, intFunc) {
			line := fmt.Sprintf(funcSignature, intFunc)
			log.Println("Adding ", line)
			_, err = fd.WriteString(line)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
