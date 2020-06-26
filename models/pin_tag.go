package models

func (data SQLDataStorage) CreatePinTag(pinID int, tagID int) error {
	const query = `
INSERT INTO pins_tags (pin_id, tag_id) VALUES (?, ?);
`
	stmt, err := data.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(pinID, tagID)
	if err != nil {
		return err
	}

	return nil
}
