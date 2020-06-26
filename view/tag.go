package view

import "hello/models"

type Tag struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

func NewTag(tag *models.Tag) *Tag {
	t := &Tag{
		ID:  tag.ID,
		Tag: tag.Tag,
	}

	return t
}

func NewTags(tags []*models.Tag) []*Tag {
	b := make([]*Tag, 0, len(tags))

	for _, tag := range tags {
		b = append(b, NewTag(tag))
	}

	return b
}
