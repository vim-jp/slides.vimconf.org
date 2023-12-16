package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"slices"
)

var tmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<meta charset="UTF-8">
<h1>{{.Site}}/{{.Path}}/</h1>
<pre>
<table>
<tr><th>Name</th><th>Size</th></tr>
{{range .Dirs}}<tr><td><a href="{{.Name}}">{{.Name}}</a></td><td align="right">{{.Info.Size}}</td></tr>
{{end}}</table>
</pre>
`))

func genIndex(name string) error {
	dirs, err := readDir(name)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(name, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, map[string]interface{}{
		"Site": "slides.vim.org",
		"Path": filepath.Clean(name),
		"Dirs": dirs,
	})
	return nil
}

func readDir(name string) ([]os.DirEntry, error) {
	// check name pointing a directory.
	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("%s is not directory", name)
	}
	dirs, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}
	return slices.DeleteFunc(dirs, func(d os.DirEntry) bool {
		return d.IsDir() || filepath.Ext(d.Name()) != ".pdf"
	}), nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("require one more directories")
	}
	for _, dir := range flag.Args() {
		err := genIndex(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
}
