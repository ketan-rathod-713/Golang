package board

import (
	"context"
	"graphql_search/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *api) GetBoardsByTitle(search string) ([]*models.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var boardsDb []*models.BoardDB

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": search, "$options": "i"}},
			{"description": bson.M{"$regex": search, "$options": "i"}},
			{"type": bson.M{"$regex": search, "$options": "i"}},
		},
	}

	cursor, err := a.Database.Collection("information").Find(ctx, filter)
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
			BoardID:     p.BoardID.Id.Hex(),
			Visible:     p.Visible,
			Description: p.Description,
			Title:       p.Title,
			Type:        p.Type,
		})
	}

	return boards, nil
}

func (a *api) GetBoard(id string) (*models.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objId,
	}
	result := a.Database.Collection("information").FindOne(ctx, filter)

	var boardDb *models.BoardDB
	err := result.Decode(&boardDb)
	if err != nil {
		return nil, err
	}

	var board = &models.Board{
		ID:          boardDb.ID.Hex(),
		BoardID:     boardDb.BoardID.Id.Hex(),
		Visible:     boardDb.Visible,
		Title:       boardDb.Title,
		Description: boardDb.Description,
		Type:        boardDb.Type,
	}

	return board, nil
}
