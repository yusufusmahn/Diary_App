package models

import "time"

type Entry struct {
    ID          int       `bson:"id"`
    Title       string    `bson:"title"`
    Body        string    `bson:"body"`
    DateCreated time.Time `bson:"date_created"`
}

type Diary struct {
    Username string  `bson:"username"`
    Password string  `bson:"password"`
    IsLocked bool    `bson:"is_locked"`
    Entries  []Entry `bson:"entries"`
}
