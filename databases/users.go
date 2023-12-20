package databases

import (
	"fmt"
)

type User struct {
	Id        int64
	UniqueId  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

// albumsByArtist queries for albums that have the specified artist name.
func UsersByEmail(email string) ([]User, error) {
	db := GetDB()
	// An albums slice to hold data from returned rows.
	var users []User

	rows, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, fmt.Errorf("userByEmail %q: %v", email, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.UniqueId, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("userByEmail %q: %v", email, err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("userByEmail %q: %v", email, err)
	}
	return users, nil
}

func AddUser(user User) (int64, error) {
	db := GetDB()
	result, err := db.Exec("INSERT INTO users (email, unique_id, password) VALUES (?, ?, ?)", user.Email, user.UniqueId, user.Password)
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    return id, nil
}

func Test() {
	fmt.Println("Okay");
}