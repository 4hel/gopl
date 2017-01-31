// small web server using html templates for github issues
package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var result *github.IssuesSearchResult
var issueListTemplate *template.Template
var issueTemplate *template.Template

func init() {
	// make GitHub API Call
	var err error
	result, err = github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// parse issues template
	issueListTemplate = template.Must(template.ParseFiles("issueList"))
	issueTemplate = template.Must(template.ParseFiles("issue"))

}

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/", issueListHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func issueListHandler(w http.ResponseWriter, r *http.Request) {
	if err := issueListTemplate.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	number, err := strconv.Atoi(r.Form["number"][0])
	if err != nil {
		log.Fatal(err)
	}
	var issue *github.Issue
	for _, item := range result.Items {
		if item.Number == number {
			issue = item
		}
	}
	if issue != nil {
		if err := issueTemplate.Execute(w, issue); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprintln(w, "issue not found")
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	name := r.Form["name"][0]
	var user *github.User
	for _, item := range result.Items {
		if current := item.User; current.Login == name {
			user = item.User
			break
		}
	}
	if user != nil {
		fmt.Fprintf(w, "<a href='%s'>%s</a>\n", user.HTMLURL, user.HTMLURL)
	} else {
		fmt.Fprintln(w, "user not found")
	}
}
