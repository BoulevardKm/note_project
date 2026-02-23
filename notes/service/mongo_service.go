package service

import (
	"context"
	"fmt"
	"notes/database"
	"notes/internal/config"
	"notes/internal/errors"
	"notes/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoService struct {
	db         *mongo.Client
	collection *mongo.Collection
}

var _ Service = (*MongoService)(nil)

func NewService(cfg *config.Config) (Service, error) {
	db, err := database.NewDatabase(cfg)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDatabaseConnection, err)
	}

	collection := db.Database(cfg.DB_NAME).Collection(cfg.DB_COLLECTION)

	return &MongoService{
		db:         db,
		collection: collection,
	}, nil
}

func (m *MongoService) Create(ctx context.Context, note models.Note) (*models.Note, error) {
	result, err := m.collection.InsertOne(ctx, bson.M{
		"name":      note.Name,
		"content":   note.Content,
		"author_id": note.AuthorID,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrNoteCreation, err)
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	note.ID = insertedID.Hex()

	return &note, nil
}

func (m *MongoService) Close() error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoService) GetByID(ctx context.Context, id string) (*models.Note, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrInvalidNoteID, err)
	}

	var note models.Note
	err = m.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&note)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: заметка с ID %s не найдена", errors.ErrNoteNotFound, id)
		}
		return nil, fmt.Errorf("%w: %v", errors.ErrDatabaseOperation, err)
	}

	note.ID = objectID.Hex()
	return &note, nil
}

func (m *MongoService) GetAll(ctx context.Context, authorId int) ([]models.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoService) Update(ctx context.Context, note models.Note) (*models.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoService) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
