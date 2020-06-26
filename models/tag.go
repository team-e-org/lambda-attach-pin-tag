package models

import (
	"time"
)

type Tag struct {
	ID        int
	Tag       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (data SQLDataStorage) CreateTag(tag *Tag) (*Tag, error) {
	const query = `
INSERT INTO tags (tag) VALUES (?);
`

	stmt, err := data.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(tag.Tag)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	tag.ID = int(id64)

	return tag, nil
}

func (data SQLDataStorage) GetTagsByTagNames(tagNames []string) ([]*Tag, error) {
	query := `
SELECT id, tag, created_at, updated_at FROM tags
`

	var whereClause string
	for i, name := range tagNames {
		if i == 0 {
			whereClause = "WHERE tag = '" + name + "'"
		} else {
			whereClause = whereClause + " OR tag = '" + name + "'"
		}
	}

	query = query + whereClause

	stmt, err := data.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	tags := make([]*Tag, 0, len(tagNames))
	for rows.Next() {
		tag := &Tag{}
		err := rows.Scan(
			&tag.ID,
			&tag.Tag,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
