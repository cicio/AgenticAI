package html

import (
	"embed"
	"io"
	"strings"
	"html/template"
	"regexp"
	"errors"
)


//go:embed * 
var files embed.FS

var (
	chatui = parse("chatui.html")
	profileShow = parse("profile/show.html")
	profileEdit = parse("profile/edit.html")

)



func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").Funcs(funcs).ParseFS(files, "layout.html", file)
	)
}

type User struct {
	ID int
	Username string
	Email ValidEmail
}

type ValidEmail string

func (e ValidEmail) Validate() error {
	// Simple regex for email validation
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(string(e)) {
		return errors.New("invalid email format")
	}
	return nil
}

type ProfileInfo struct {
	UserID User.ID
	FirstName string
	LastName string
	Email ValidEmail
	Phone string
	Company string
}


type ChatParams struct {
	User User
	Query string
	Answer string
}

func Chat(w io.Writer, params ChatParams, partial string) error {
	if partial =="" {
		partial = "layout.html"
	}
	return chatui.ExecuteTemplate(w, partial, params)
}

type ProfileShowParams struct {
	Title string
	ProfileInfo ProfileInfo
}

func ProfileShow(w io.Writer, params ProfileShowParams, partial string) error {	
	if partial == "" {
		partial = "layout.html"
	}
	return profileShow.ExecuteTemplate(w, partial, params)
}

type ProfileEditParams struct {
	User User
	ProfileInfo ProfileInfo
}

func ProfileEdit(w io.Writer, params ProfileEditParams, partial string) error {	
	if partial == "" {
		partial = "layout.html"
	}
	return profileEdit.ExecuteTemplate(w, partial, params)
}

var funcs = template.FuncMap{
	"uppercase": strings.ToUpper,
	"lowercase": strings.ToLower,
	"trim": strings.TrimSpace,
	"email": func(e ValidEmail) string {
		return string(e)
	}
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").Funcs(funcs).ParseFS(files, "layout.html", file)
	)
}	
