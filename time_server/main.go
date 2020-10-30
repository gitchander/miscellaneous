package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//withFormat()
	withTemplate()
}

func withFormat() {

	var formatIndex = `<!DOCTYPE html>
<html>
	<head>
		<title>current time</title>
	</head>
	<body>
		<h1>%s</h1>
	</body>
</html>`

	handler := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, formatIndex, now)
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9090", nil)
	checkError(err)
}

func withTemplate() {

	var pattern = `{{define "turok"}}
<!DOCTYPE html>
<html>
	<head>
		<title>current time</title>
	</head>
	<body>
		<h1>{{.}}</h1>
	</body>
</html>
{{end}}`

	t, err := template.New("master").Parse(pattern)
	checkError(err)

	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("request from client:", r.RemoteAddr)
		os.Stderr.WriteString("error: " + r.RemoteAddr)

		now := time.Now().Format("2006-01-02 15:04:05")

		log.Println("now:", now)

		err = t.ExecuteTemplate(w, "turok", now)
		checkError(err)
	}

	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":9090", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
