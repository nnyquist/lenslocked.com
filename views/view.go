package views

import "html/template"

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/footer.gohtml",
	) // hardcoded for now, but will be updated later on to be dynamic

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

type View struct {
	Template *template.Template
	Layout   string
}
