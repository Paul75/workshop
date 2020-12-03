package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type UserCollection map[string]User

func (u UserCollection) String() string {
	res := ""
	for _, elem := range u {
		res += elem.FirstName + "\n"
	}
	return res
}

func main() {
	data, err := Get("http://localhost:8080/users")
	if err != nil {
		panic(err)
	}
	payload := make(UserCollection)
	if err := json.Unmarshal(data, &payload); err != nil {
		panic(err)
	}

	t, err := template.New("displayuser").Parse("{{.}}")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, payload)

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

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
