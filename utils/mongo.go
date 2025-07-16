package utils

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

var MongoClient *mongo.Client
var DiaryDB *mongo.Database

func init() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI not found in environment")
    }

    // Use the updated mongo.Connect pattern
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Assign to global vars
    MongoClient = client
    DiaryDB = client.Database("diaryApp") // Use your DB name

    fmt.Println("Connected to MongoDB Atlas (diaryApp)")
}

// GetCollection returns a MongoDB collection from diaryApp database
func GetCollection(name string) *mongo.Collection {
    return DiaryDB.Collection(name)
}


