package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type Job struct {
	Employer string
	Role string
	DateStart string
	DateEnd string
	Description string
	Highlights string
	Logo []byte
	LogoLocation string
	Url string
}

type Education struct {
	School string
	Subject string
	Degree string
	DateStart string
	DateEnd string
	Highlights string
	Logo []byte
	LogoLocation string
	Url string
}

type CustomTag struct {
	Name string
	Level string
	Category string
	AdditionalText string
	Url string
}

type CV2 struct {
	Title string
	FullName string
	FaoName string
	Birthday string
	Email string
	Address string
	Phone string
	Picture []byte
	PictureLocation string
	Jobs []Job
	Edu []Education
	Languages []CustomTag
	Skills []CustomTag
	CustomTags []CustomTag
	Categories []string
	PersonalText string
	Body []byte
}

//createCVHandler handles the form input and assigns corresponding functions in order to create a CV 2.0.
func createCVHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(8192)
	cv2 := new(CV2)
	var keys, values = "", ""
	for key, value := range r.Form {
		if strings.Contains(string(key), "School") {
			
		}
		for index, _ := range value {
			values += value[index]
		}
		keys += "---\n" + key + ":\n" + values + "\n---\n\n"
		values = ""
	}
	cv2.Title = r.FormValue("FullName") + "_" + createTimestamp()
	cv2.FullName = r.FormValue("FullName")
	cv2.Body = []byte(keys)

	// file handling
	file, header, err := r.FormFile("LogoJob")
	if err != nil {
		// no logo uploaded
	} else {
		defer file.Close()
		out, _ := os.Create("../output/" + header.Filename)
		defer out.Close()
		//cv2.Jobs[0].PictureLocation = "../output/" + header.Filename
		_, err = io.Copy(out, file)
	}

	// .cv2 file creation
  err_create := cv2.createCV2()
	if err_create != nil {
		http.Error(w, err_create.Error(), http.StatusInternalServerError)
    return
  }

	// .html file creation
	if !createHTML(cv2, "basic") {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// .svg file creation
	if !createSVG(cv2, "comic") {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// remove the output with a delay
	go cv2.removeCV2()
	/* // CODE BEFORE GOROUTINE, go NEEDED FOR TIMER, ERROR HANDLING?
	err_remove := cv2.removeCV2()
	for e := range err_remove {
		if err_remove[e] != nil {
			http.Error(w, "Removing the output didn't quite work.", http.StatusInternalServerError)
	    return
	  }
	}*/

	// redirect
  http.Redirect(w, r, "/cv_created_"+cv2.Title, http.StatusFound)
}

//createHTML creates an .html file with data provided in the *CV2 struct and a template name.
func createHTML(cv2 *CV2, tmplName string) bool {
	cv2_output := "../output/"+ cv2.FullName + "_" + createTimestamp() +".html"
	t, err := template.ParseFiles("../templates/"+ tmplName +".html")
	if err != nil {
		return false
	} else {
		f, _ := os.Create(cv2_output)
		w := new(bytes.Buffer)
		t.ExecuteTemplate(w, "tmplHTML", cv2)
		w.WriteTo(f)
		f.Close()
		return true
	}
}

//createSVG creates an SVG with data provided in the *CV2 struct and a template name.
func createSVG(cv2 *CV2, tmplName string) bool {
	cv2_output := "../output/"+ cv2.FullName + "_" + createTimestamp() +".svg"
	t, err := template.ParseFiles("../templates/"+ tmplName +".svg")
	if err != nil {
		return false
	} else {
		f, _ := os.Create(cv2_output)
		w := new(bytes.Buffer)
		t.ExecuteTemplate(w, "tmplSVG", cv2)
		w.WriteTo(f)
		f.Close()
		return true
	}
}
