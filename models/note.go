package models

import (
	"errors"

	"github.com/orlmonteverde/go-notes/config"
)

type Note struct {
	Id    int
	Title string
	Body  string
}

func GetNoteById(id int) (note Note, err error) {
	db := config.GetConnection()
	defer db.Close()

	q := `SELECT id, title, body FROM notes WHERE id=$1`
	err = db.QueryRow(q, id).Scan(&note.Id, &note.Title, &note.Body)
	return
}

func GetNotes() ([]Note, error) {
	db := config.GetConnection()
	defer db.Close()

	var notes []Note
	q := `SELECT id, title, body FROM notes`
	rows, err := db.Query(q)
	defer rows.Close()
	var n Note
	for rows.Next() {
		rows.Scan(&n.Id, &n.Title, &n.Body)
		notes = append(notes, n)
	}
	return notes, err
}

func CreateNote(n Note) error {
	db := config.GetConnection()
	defer db.Close()

	q := `INSERT INTO
			notes(title, body)
			VALUES($1, $2)`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(n.Title, n.Body)
	if err != nil {
		return err
	}

	i, err := r.RowsAffected()
	if err != nil || i != 1 {
		return errors.New("Expected 1 row affected")
	}
	return nil
}
