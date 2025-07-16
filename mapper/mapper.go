package mapper

import (
    "diary-app/dto"
    "diary-app/models"
    "strings"
)

func ToDiaryModel(req dto.RegisterDiaryRequest) models.Diary {
    return models.Diary{
        Username: strings.ToLower(req.Username), 
        Password: req.Password,
        IsLocked: false,
        Entries:  []models.Entry{},
    }
}

func ToEntryModel(req dto.EntryRequest) models.Entry {
    return models.Entry{
        Title: req.Title,
        Body:  req.Body,
    }
}

func ToEntryModelWithID(id int, req dto.EntryRequest) models.Entry {
    return models.Entry{
        ID:    id,
        Title: req.Title,
        Body:  req.Body,
    }
}

func ToDiaryResponse(d models.Diary) dto.DiaryResponse {
    entries := make([]dto.EntryResponse, len(d.Entries))
    for i, entry := range d.Entries {
        entries[i] = dto.EntryResponse{
            ID:          entry.ID,
            Title:       entry.Title,
            Body:        entry.Body,
            DateCreated: entry.DateCreated,
        }
    }

    return dto.DiaryResponse{
        Username: d.Username,
        IsLocked: d.IsLocked,
        Entries:  entries,
    }
}
