package repository

import (
    "context"
    "diary-app/models"
    "diary-app/utils"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// DiaryRepository handles MongoDB operations for Diary
type DiaryRepository struct {
    collection *mongo.Collection
}

// NewDiaryRepository returns a new instance of DiaryRepository
func NewDiaryRepository() *DiaryRepository {
    return &DiaryRepository{
        collection: utils.GetCollection("diaries"),
    }
}

// CreateDiary inserts a new diary into the collection
func (r *DiaryRepository) CreateDiary(diary models.Diary) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.InsertOne(ctx, diary)
    return err
}

// FindDiaryByUsername finds a diary document by username
func (r *DiaryRepository) FindDiaryByUsername(username string) (models.Diary, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var diary models.Diary
    err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&diary)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return diary, errors.New("diary not found")
        }
        return diary, err
    }
    return diary, nil
}

// AddEntry appends a new entry to the diary of the given username
func (r *DiaryRepository) AddEntry(username string, entry models.Entry) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{"$push": bson.M{"entries": entry}}
    _, err := r.collection.UpdateOne(ctx, bson.M{"username": username}, update)
    return err
}

// UpdateEntry updates an entry by ID in a specific user's diary
func (r *DiaryRepository) UpdateEntry(username string, entry models.Entry) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"username": username, "entries.id": entry.ID}
    update := bson.M{"$set": bson.M{
        "entries.$.title": entry.Title,
        "entries.$.body":  entry.Body,
    }}
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

// DeleteEntry removes an entry from a diary based on entry ID
func (r *DiaryRepository) DeleteEntry(username string, entryID int) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{"$pull": bson.M{"entries": bson.M{"id": entryID}}}
    _, err := r.collection.UpdateOne(ctx, bson.M{"username": username}, update)
    return err
}

// LockDiary sets is_locked to true
func (r *DiaryRepository) LockDiary(username string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"is_locked": true}})
    return err
}

// UnlockDiary sets is_locked to false
func (r *DiaryRepository) UnlockDiary(username string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"is_locked": false}})
    return err
}

// DeleteDiary deletes the entire diary document by username
func (r *DiaryRepository) DeleteDiary(username string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.DeleteOne(ctx, bson.M{"username": username})
    return err
}

