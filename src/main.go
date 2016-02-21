package main

import (
  "html/template"
	"net/http"
  "os"
)

type DataPassed struct {
  Title string
}

type MultipleDomains map[string]http.Handler

func (md MultipleDomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  handler := md[r.Host]

  if handler != nil {
    handler.ServeHTTP(w, r)
  } else {
    http.Redirect(w, r, "/", http.StatusFound)
  }
}

//templateHandler assigns each URL to the corresponding template structure.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	data := new(DataPassed)
  layout := "../structure.html"
	content := "../views/" + r.URL.Path + ".html"
	if r.URL.Path == "/" {content = "../views/index.html"}
	if len(r.URL.Path) > 12 && r.URL.Path[1:12] == "cv_created_" {
    data.Title = r.URL.Path[12:]
    content = "../views/cv_created.html"
	}

	info, err := os.Stat(content)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(layout, content)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "structure", data)
}

//main starts the web server and routes URLS.
func main() {
  muxcv := http.NewServeMux()
  muxtk := http.NewServeMux()

  md := make(MultipleDomains)
  md["localhost:8080"] = muxcv
  md["localhost:2015"] = muxtk

  muxcv.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("../public"))))
	muxcv.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("../output"))))
	muxcv.HandleFunc("/createCV/", createCVHandler)
  muxcv.HandleFunc("/", templateHandler)


	http.ListenAndServe(":8080", md)
}
