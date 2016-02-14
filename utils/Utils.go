package utils

import (
	"crypto/rand"
	"fmt"
	"html/template"

	"github.com/russross/blackfriday"
)

func GeneratId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func ConvertMarkdownToHtml(md string) string {
	return string(blackfriday.MarkdownBasic([]byte(md)))
}

func Unescape(s string) interface{} {
	return template.HTML(s)
}
