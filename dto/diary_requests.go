package dto

type RegisterDiaryRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type EntryRequest struct {
    Title string `json:"title"`
    Body  string `json:"body"`
}


