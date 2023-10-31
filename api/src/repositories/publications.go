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

func (p publications) CreatePublication(publication models.Publication) (int64, error) {
	ctx := context.Background()

	result, err := p.queries.CreatePublication(ctx, database.CreatePublicationParams{
		Title:    publication.Title,
		Content:  publication.Content,
		AuthorID: publication.AuthorID,
	})
	if err != nil {
		return 0, nil
	}

	publicationId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(publicationId), nil
}

func (p publications) FindPublications(userId int64) ([]database.FindPublicationsRow, error) {
	ctx := context.Background()

	publications, err := p.queries.FindPublications(ctx, database.FindPublicationsParams{
		ID:         int32(userId),
		FollowerID: int32(userId),
	})
	if err != nil {
		return nil, err
	}

	return publications, nil
}

func (p publications) FindPublicationById(id int64) (database.FindPublicationByIdRow, error) {
	ctx := context.Background()

	publication, err := p.queries.FindPublicationById(ctx, int32(id))
	if err != nil {
		return database.FindPublicationByIdRow{}, nil
	}

	return publication, nil
}

func (p publications) UpdatePublication(publicationId int64, publication models.Publication) error {
	ctx := context.Background()

	if err := p.queries.UpdatePublication(ctx, database.UpdatePublicationParams{
		ID:      int32(publicationId),
		Title:   publication.Title,
		Content: publication.Content,
	}); err != nil {
		return err
	}

	return nil
}

func (p publications) DeletePublicationById(id int64) error {
	ctx := context.Background()

	if err := p.queries.DeletePublicationById(ctx, int32(id)); err != nil {
		return err
	}

	return nil
}

func (p publications) GetUserPublications(userId int64) ([]database.GetUserPublicationsRow, error) {
	ctx := context.Background()

	publications, err := p.queries.GetUserPublications(ctx, int32(userId))
	if err != nil {
		return nil, err
	}

	return publications, err
}

func (p publications) LikePublication(publicationId int64) error {
	ctx := context.Background()

	if err := p.queries.LikePublication(ctx, int32(publicationId)); err != nil {
		return err
	}

	return nil
}

func (p publications) DislikePublication(publicationId int64) error {
	ctx := context.Background()

	if err := p.queries.DislikePublication(ctx, int32(publicationId)); err != nil {
		return err
	}

	return nil
}
