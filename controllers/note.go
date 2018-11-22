package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/orlmonteverde/go-notes/models"
)

func Home(c echo.Context) error {
	notes, err := models.GetNotes()
	if err != nil {
		c.Error(err)
	}
	title := "Home"
	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"Title": title,
		"Data":  notes,
	})
}

func GetPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	note, err := models.GetNoteById(id)
	if err != nil {
		c.Error(err)
	}
	title := "Note"
	m := echo.Map{
		"Title": title,
		"Data":  note,
	}
	return c.Render(http.StatusOK, "show", m)
}

func PostNote(c echo.Context) error {
	var note models.Note
	note.Title = c.FormValue("title")
	note.Body = c.FormValue("body")
	err := models.CreateNote(note)
	if err != nil {
		c.Error(err)
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func PostRedirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/post/")
}

func CreateNote(c echo.Context) error {
	title := "New Note"
	return c.Render(http.StatusOK, "create", map[string]string{"Title": title})
}
