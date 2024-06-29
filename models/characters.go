package models

import "database/sql"

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IconUrl     string `json:"icon_url"`
	Class       string `json:"class"`
	Description string `json:"description"`
}

func CreateCharacter(db *sql.DB, c Character) error {
	stmt, err := db.Prepare("INSERT INTO characters(name, icon_url, class, description) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.Name, c.IconUrl, c.Class, c.Description)
	if err != nil {
		return err
	}

	return nil
}

func GetCharacters(db *sql.DB) ([]Character, error) {
	rows, err := db.Query("SELECT * FROM characters")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	characters := []Character{}

	for rows.Next() {
		var c Character
		if err := rows.Scan(&c.ID, &c.Name, &c.IconUrl, &c.Class, &c.Description); err != nil {
			return nil, err
		}

		characters = append(characters, c)
	}

	return characters, nil
}

func GetCharacterById(db *sql.DB, id int) (Character, error) {
	var c Character
	err := db.QueryRow("SELECT * FROM characters WHERE id = ?", id).Scan(&c.ID, &c.Name, &c.IconUrl, &c.Class, &c.Description)
	if err != nil {
		return c, err
	}

	return c, nil
}
