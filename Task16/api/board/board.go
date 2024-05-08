package board

import (
	"context"
	"graphql_search/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (a *api) GetBoardsByTitle(title string) ([]*models.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var boardsDb []*models.BoardDB

	cursor, err := a.Database.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &boardsDb)
	if err != nil {
		return nil, err
	}

	var boards []*models.Board
	for _, p := range boardsDb {
		boards = append(boards, &models.Board{
			ID:          p.ID.Hex(),
			BoardID:     p.BoardID,
			Visible:     p.Visible,
			Description: p.Description,
			Title:       p.Title,
			Type:        p.Type,
		})
	}

	return boards, nil
}
