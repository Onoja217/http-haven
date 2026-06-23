package handlers

import (
	"html/template"
	"net/http"
)

const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
  <h1>{{.Title}}</h1>
  <p>
    {{if eq .Style "bold"}}
      <strong>{{.Body}}</strong>
    {{else}}
      {{.Body}}
    {{end}}
  </p>
</body>
</html>
`

type PageData struct {
	Title string
	Body  string
	Style string
}

// template.Must panics if template parsing fails at startup.
// It is used when failure should crash the program immediately.

var tmpl = template.Must(template.New("page").Parse(tmplStr))

func Render(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	style := r.URL.Query().Get("style")

	if title == "" || body == "" {
		http.Error(w, "title and body are required", http.StatusBadRequest)
		return
	}

	// IMPORTANT: set header BEFORE Execute
	w.Header().Set("Content-Type", "text/html")

	data := PageData{
		Title: title,
		Body:  body,
		Style: style,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template execution failed", http.StatusInternalServerError)
		return
	}
}