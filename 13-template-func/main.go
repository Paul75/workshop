package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Any("", home)
	r.Run(":9090")
}

func home(ctx *gin.Context) {
	data, err := Get("http://localhost:8080/users")
	if err != nil {
		panic(err)
	}

	payload := make(map[string]User)
	if err := json.Unmarshal(data, &payload); err != nil {
		panic(err)
	}

	// put this section of code outside of this function
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}
	// end section

	dataTmpl := map[string]interface{}{
		"Title": "Liste des utilisateurs",
		"Users": payload,
	}

	err = tmpl.Execute(ctx.Writer, dataTmpl)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}

var funcMap = template.FuncMap{
	"title": strings.ToUpper,
	// add the function Concat to the list with the key name concat.
}

func Concat(firstName, lastName string) string {
	// implement this function so the result is concatenated and the lastName is in uppercase.
	return ""
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

const templateText = `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="fr" lang="fr" dir="ltr">
  <head>
  <h1>{{ title .Title }}</h1>
  </head>
  <body>
	<ul>
	{{- range .Users}}
		<li>{{ concat .FirstName .LastName }}</li>
	{{- end}}
	</ul>
  </body>
</html>`
