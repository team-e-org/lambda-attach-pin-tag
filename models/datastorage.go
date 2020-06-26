package models

import "database/sql"

type DataStorage interface {
	CreateTag(tag *Tag) (*Tag, error)
	GetTagsByTagNames(tagNames []string) ([]*Tag, error)
}

func NewSQLDataStorage(sqlDB *sql.DB) DataStorage {
	return SQLDataStorage{DB: sqlDB}
}

type SQLDataStorage struct {
	DB *sql.DB
}
