package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type ContactDetail struct {
	Email   string
	Subject string
	Message string
}

var tmpl = template.Must(template.ParseFiles("forms.html"))

func main() {

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, struct{ Success bool }{false})
		}

		details := ContactDetail{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		details.string()

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}

func (c *ContactDetail) string() {
	fmt.Printf("email: %v\n", c.Email)
	fmt.Printf("subject: %v\n", c.Subject)
	fmt.Printf("message: %v\n", c.Message)
}
