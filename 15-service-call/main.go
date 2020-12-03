package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

func main() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	payload := make(map[string]User)
	if err := json.Unmarshal(data, &payload); err != nil {
		panic(err)
	}

	t, err := template.New("displayuser").Parse("{{.}}")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, payload)

}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
