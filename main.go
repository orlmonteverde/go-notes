package main

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/orlmonteverde/go-notes/config"
	"github.com/orlmonteverde/go-notes/controllers"
)

var port string

type Template struct {
	templates *template.Template
}

func (t Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	if err := config.MakeMigrations(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.Static("static", "public/")
	e.GET("/", controllers.Home)
	e.GET("/post/:id", controllers.GetPost)
	e.GET("/post", controllers.PostRedirect)
	e.GET("/post/", controllers.CreateNote)
	e.POST("/post/", controllers.PostNote)

	e.Logger.Fatal(e.Start(":" + port))
}

func init() {
	if port = os.Getenv("PORT"); port == "" {
		port = "1323"
	}
}
