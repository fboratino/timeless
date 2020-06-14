package views

import "html/template"

// NewView aggregates all templates together
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View template to use in other pages
type View struct {
	Template *template.Template
}
