package controllers

import (
	"html/template"
	"net/http"

	"github.com/sergei-doroshenko/p014_go_test/models"
	"github.com/sergei-doroshenko/p014_go_test/utils"

	// import for martini
	//"github.com/codegansta/martini"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

// mongoDB collection
var postsCollection *mgo.Collection

// inject render defined in render.Renderer
func indexHandler2(ren render.Render) {
	postDocuments := []models.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)

	posts3 := []models.PostMD{}

	for _, doc := range postDocuments {
		post := models.PostMD{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts3 = append(posts3, post)
	}

	ren.HTML(200, "index", posts3)
}

func writeHandler2(ren render.Render) {

	ren.HTML(200, "write", nil)
}

func editHandler2(ren render.Render, r *http.Request, params martini.Params) {

	id := params["id"]

	postDocument := models.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		ren.Redirect("/")
		return
	}

	post := models.PostMD{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	ren.HTML(200, "write", post)
}

func savePostHandler2(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownToHtml(contentMarkdown)

	postDocument := models.PostDocument{id, title, contentHtml, contentMarkdown}

	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id := utils.GeneratId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	ren.Redirect("/")
}

func deleteHandler2(ren render.Render, r *http.Request, params martini.Params) {
	id := params["id"]
	if id == "" {
		ren.Redirect("/")
	}

	postsCollection.RemoveId(id)

	ren.Redirect("/")
}

func getHtmlHandler2(ren render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkdownToHtml(md)

	ren.JSON(200, map[string]interface{}{"html": html})
}

func RunWithMongo() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	postsCollection = session.DB("blog").C("posts")

	// martini package
	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": utils.Unescape}

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

	// map http requests
	m.Get("/", indexHandler2)
	m.Get("/write", writeHandler2)
	m.Get("/edit/:id", editHandler2)
	m.Get("/delete/:id", deleteHandler2)
	m.Post("/SavePost", savePostHandler2)
	m.Post("/gethtml", getHtmlHandler2)

	m.Run()
}
