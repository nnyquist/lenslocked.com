package views

import (
	"html/template"
	"path/filepath"
)

var (
	layoutDir   string = "views/layouts/"
	templateExt string = ".gohtml"
)

// NewView ...
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...) // ... unpack the slice into strings
	if err != nil {
		// our view should only be used when we are initializing the entire application
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View ...
type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles returns a slice of strings representing
// the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}
