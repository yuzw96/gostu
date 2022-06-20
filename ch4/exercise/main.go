package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	Title     string
	User      User
	Milestone Milestone
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Login string
}

type Milestone struct {
	Title       string
	Description string
}

func SearchIssues(params string) (*IssuesSearchResult, error) {
	resp, err := http.Get(IssuesURL + "?q=" + params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	q := r.FormValue("key")
	result, err := SearchIssues(q)
	if err != nil {
		log.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println(err)
	}
}
