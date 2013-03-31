package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "time"
    "text/template"
)

func hello(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("secheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello 蛤蟆皮!")
}

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/time", my_time)
    http.HandleFunc("/xss", xss)
    http.HandleFunc("/templates", templates)
    http.HandleFunc("/layout", layout)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func my_time(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    fmt.Printf("Go launched at %s\n", t.Local())
    fmt.Fprintf(w, "Date UTC:%s<br>", t.Local())
    fmt.Fprintf(w, "Time Unix%d<br>", t.Unix())
}

func xss(w http.ResponseWriter, r *http.Request) {
    t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
    t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
}

func layout(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/layout.html", "templates/index.html")
    t.ExecuteTemplate(w, "layout", nil)
    t.Execute(w, nil)
}

func templates(w http.ResponseWriter, r *http.Request) {
    s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
    s1.ExecuteTemplate(w, "header", nil)
    fmt.Printf("header\n")
    s1.ExecuteTemplate(w, "content", nil)
    fmt.Printf("content\n")
    s1.ExecuteTemplate(w, "footer", nil)
    fmt.Printf("footer\n")
    s1.Execute(w, nil)
}