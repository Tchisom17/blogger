package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)
var tpl *template.Template

type BlogPost struct {
	Id	string
	Title string
	Body string
	//Created time.Time
}
var Posts = []BlogPost{}


func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
//veiw
func index(res http.ResponseWriter, req *http.Request){
	tpl.ExecuteTemplate(res, "index.gohtml", Posts)
}
//form
func blog(res http.ResponseWriter, req *http.Request){
//	p:= BlogPost{
//	Title: "tilte",
//	Body: "this is a boby",
//}
	tpl.ExecuteTemplate(res, "blog.gohtml", Posts)
}
//proccess form
func createBlog(res http.ResponseWriter, req *http.Request){
	//var Posts = []BlogPost{}
	req.ParseForm()
	title := req.FormValue("title")
	bdy := req.FormValue("body")
	d := BlogPost{
		Id: uuid.New().String(),
		Title: title,
		Body: bdy,
	}
	Posts = append(Posts, d)
	//tpl.ExecuteTemplate(res, "create.gohtml", d)
	http.Redirect(res,req,"/", http.StatusMovedPermanently)
}
//delete post
func delete(res http.ResponseWriter, req *http.Request)  {
	param:= chi.URLParam(req, "Id")
	for i,v:= range Posts{
		if param == v.Id{
			Posts= append(Posts[:i], Posts[i+1:]...)
		}
	}
	http.Redirect(res,req,"/", http.StatusMovedPermanently)
}

func main() {
	req := chi.NewRouter()
	req.Use(middleware.Logger)
	req.Get("/", index)
	req.Get("/{{Id}}", delete)
	req.Get("/blog", blog)
	req.Post("/blog/create", createBlog)
	http.ListenAndServe(":2000", req)
}

