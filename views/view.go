package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   string = "views/layouts/"
	templateDir string = "views/"
	templateExt string = ".gohtml"
)

// NewView ...
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
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

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
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

// addTemplatePath takes in a slice of strings representing file paths
// for templates, and it prepends the templateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if templateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = templateDir + f
	}
}

// addTemplateExt takes in a slice of strings
// representing file paths for templates and it appends
// the templateExt extension to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"home.gohtml"} if templateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + templateExt
	}
}
