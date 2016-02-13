package models

type PostMD struct {
	Id              string
	Title           string
	ContentHtml     string
	ContentMarkdown string
}

func NewPostMD(id, title, contentHtml, contentMarkdown string) *PostMD {
	return &PostMD{id, title, contentHtml, contentMarkdown}
}
