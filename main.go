package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	homeTmpl = template.Must(template.ParseFiles("template/home.html"))
	addTmpl = template.Must(template.ParseFiles("template/add.html"))
	showallTmpl = template.Must(template.ParseFiles("template/showall.html"))
	
)

var list []string

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		homeTmpl.Execute(w, nil)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){
		
		if r.Method == http.MethodPost{
			r.ParseForm()

			newList := r.FormValue("item")
			
			if newList != "" {
				list = append(list, newList)
			}

			addTmpl.Execute(w, list)
			return
		}
			addTmpl.Execute(w, list)
	})

	http.HandleFunc("/showall", func(w http.ResponseWriter, r *http.Request){

			showallTmpl.Execute(w, list)


		
	})


	fmt.Println("Your server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}