package databases

import (
	"fmt"
)

type Note struct {
	Id        int64
	UniqueId  string
	Title     string
	Content  string
	Metadata  string
	TagId int64
	UserId int64
	CreatedAt string
	UpdatedAt string
}

func AddNote(note Note) (int64, error) {
	db := GetDB()
	result, err := db.Exec("INSERT INTO notes (unique_id, title, content, metadata, tag_id, userId) VALUES (?, ?, ?, ?, ?, ?)", note.UniqueId, note.Title, note.Content, note.Metadata, note.TagId, note.UserId)
    if err != nil {
        return 0, fmt.Errorf("addNote: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addNote: %v", err)
    }
    return id, nil
}


func GetNotes() ([]Note, error) {
	db := GetDB()
	// An albums slice to hold data from returned rows.
	var notes []Note

	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, fmt.Errorf("notes: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.Id, &note.UniqueId, &note.Title, &note.Content, &note.Metadata, &note.TagId, &note.UserId, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, fmt.Errorf("getNotes: %v", err)
		}
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getNotes: %v", err)
	}
	return notes, nil
}