package html

import (
	"html/template"
	"log"
	"net/http"
)

const tmplHome = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

// ImgUpload caricamento di un'immagine dal telefono
func Home(res http.ResponseWriter, req *http.Request) {

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(tmplHome))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(res, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
