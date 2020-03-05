//go:generate go run github.com/efskap/statics -f

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"syscall/js"
	"unicode/utf8"
)

var document js.Value // shortcut for convenience
var parsedTmpl *template.Template
var done chan bool // used to keep the program from exiting

func init() {
	done = make(chan bool)
	document = js.Global().Get("document")
	loadTemplate()
}
func loadTemplate() {
	tmplSrc := string(files["parsed.gohtml"])
	funcs := template.FuncMap{}
	parsedTmpl = template.Must(template.New("parsed").Funcs(funcs).Parse(tmplSrc))
}

func main() {
	js.Global().Set("parseFile", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			setStatus("parsing...")

			array := args[0]
			fileBytes := make([]uint8, array.Get("byteLength").Int())
			js.CopyBytesToGo(fileBytes, array)
			reader := bytes.NewBuffer(fileBytes)

			if result, err := processFile(reader); err != nil {
				showError(err)
			} else {
				displayResults(result)
			}
		}()
		return nil
	}))

	js.Global().Call("onReady")

	<-done // prevents the program from exiting (unless you write to this channel)
	println("exiting")
}

type Results struct {
	FirstLine string
	FileSize  int
}

func processFile(file *bytes.Buffer) (out Results, err error) {
	out.FileSize = file.Len()
	str := file.String()
	if utf8.ValidString(str) {
		out.FirstLine = strings.SplitN(str, "\n", 2)[0]
	}
	return
}

func displayResults(results Results) {
	setStatus("")
	getElementByID("parsed").Set("hidden", false)
	buffer := new(bytes.Buffer)
	if err := parsedTmpl.ExecuteTemplate(buffer, "parsed", results); err != nil {
		showError(err)
	}
	getElementByID("parsed").Set("innerHTML", buffer.String())
}

func showError(err error) {
	fmt.Println(err)
	setStatus(err.Error())
}

func setStatus(text string) {
	getElementByID("status-text").Set("hidden", false)
	getElementByID("status-text").Set("textContent", text)
}

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

