package dto

import "time"

type EntryResponse struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Body        string    `json:"body"`
    DateCreated time.Time `json:"date_created"`
}

type DiaryResponse struct {
    Username string          `json:"username"`
    IsLocked bool            `json:"is_locked"`
    Entries  []EntryResponse `json:"entries"`
}
