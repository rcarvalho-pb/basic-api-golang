package repositories

import (
	"api/db/database"
	"api/src/models"
	"context"
)

type publications struct {
	queries *database.Queries
}

func NewPublicationRepository(queries *database.Queries) *publications {
	return &publications{queries}
}

func (p publications) CreatePublication(publication models.Publication) (uint32, error) {
	ctx := context.Background()

	result, err := p.queries.CreatePublication(ctx, database.CreatePublicationParams{
		Title: publication.Title,
		Content: publication.Content,
		AuthorID: publication.AuthorID,
		Likes: publication.Likes,
	})
	if err != nil {
		return 0, nil
	}

	publicationId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(publicationId), nil
}