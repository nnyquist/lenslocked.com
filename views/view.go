package views

import "html/template"

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml") // hardcoded for now, but will be updated later on to be dynamic

	t, err := template.ParseFiles(files...) // ... unpack the slice into strings
	if err != nil {
		// our view should only be used when we are initializing the entire application
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
