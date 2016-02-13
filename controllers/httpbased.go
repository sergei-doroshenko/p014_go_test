package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sergei-doroshenko/p014_go_test/models"
	"github.com/sergei-doroshenko/p014_go_test/utils"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Hello, Sergei!!!!</h1>")
	t, err := template.ParseFiles("views/index.html", "views/header.html", "views/footer.html") //
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(posts) // write all posts

	fmt.Println(counter)

	t.ExecuteTemplate(w, "index", posts)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/write.html", "views/header.html", "views/footer.html") //
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/write.html", "views/header.html", "views/footer.html") //
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	t.ExecuteTemplate(w, "write", post)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post

	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id := utils.GeneratId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

func RunWithHttp() {
	fmt.Println("HttpBased-1. Server started on port :3000")

	// we need to cut 'public' and left /css/app.css for example
	// static resources handle
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/SavePost", savePostHandler)
	// listen port
	http.ListenAndServe(":3000", nil)
}
