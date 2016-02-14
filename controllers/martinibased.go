package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sergei-doroshenko/p014_go_test/models"
	"github.com/sergei-doroshenko/p014_go_test/utils"

	// import for martini
	//"github.com/codegansta/martini"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// inject render defined in render.Renderer
func indexHandler1(ren render.Render) {

	fmt.Println(counter)

	ren.HTML(200, "index", postsMD)
}

func writeHandler1(ren render.Render) {

	ren.HTML(200, "write", nil)
}

func editHandler1(ren render.Render, r *http.Request, params martini.Params) {
	// id := r.FormValue("id")
	id := params["id"]
	post, found := postsMD[id]
	if !found {
		ren.Redirect("/")
	}

	ren.HTML(200, "write", post)
}

func savePostHandler1(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownToHtml(contentMarkdown)

	var post *models.PostMD

	if id != "" {
		post = postsMD[id]
		post.Title = title
		post.ContentHtml = contentHtml
		post.ContentMarkdown = contentMarkdown
	} else {
		id := utils.GeneratId()
		post := models.NewPostMD(id, title, contentHtml, contentMarkdown)
		postsMD[post.Id] = post
	}

	ren.Redirect("/")
}

func deleteHandler1(ren render.Render, r *http.Request, params martini.Params) {
	id := params["id"]
	if id == "" {
		ren.Redirect("/")
	}

	delete(postsMD, id)

	ren.Redirect("/")
}

func getHtmlHandler(ren render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkdownToHtml(md)

	ren.JSON(200, map[string]interface{}{"html": html})
}

func unescape(s string) interface{} {
	return template.HTML(s)
}

func RunWithMartini2() {
	// martini package
	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	m.Use(render.Renderer(render.Options{
		Directory:  "views/martini",                     // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		// Delims:          render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
		IndentXML:  true,    // Output human readable XML
		// HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html"
	}))

	// handle static resources
	staticOptions := martini.StaticOptions{Prefix: "public"}
	m.Use(martini.Static("public", staticOptions))

	// some kind of filter
	m.Use(func(r *http.Request) {
		if r.URL.Path == "/write" {
			counter++
		}
	})

	// map http requests
	m.Get("/", indexHandler1)
	m.Get("/write", writeHandler1)
	m.Get("/edit/:id", editHandler1)
	m.Get("/delete/:id", deleteHandler1)
	m.Post("/SavePost", savePostHandler1)
	m.Post("/gethtml", getHtmlHandler)

	m.Get("/test", func() string {
		return "hello-test"
	})

	m.Run()
}

/*func RunWithMartini1() {
	// martini package
	m := martini.Classic()
	// handle static resources
	staticOptions := martini.StaticOptions{Prefix: "public"}
	m.Use(martini.Static("public", staticOptions))

	// some kind of filter
	m.Use(func(r *http.Request) {
		if r.URL.Path == "/write" {
			counter++
		}
	})

	// map http requests
	m.Get("/", indexHandler1)
	m.Get("/write", writeHandler1)
	m.Get("/edit", editHandler1)
	m.Get("/delete", deleteHandler1)
	m.Post("/SavePost", savePostHandler1)

	m.Get("/test", func() string {
		return "hello-test"
	})

	m.Run()
}*/
