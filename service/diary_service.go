package service

import (
	"diary-app/models"
	"diary-app/repository"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type DiaryService struct {
	repo *repository.DiaryRepository
}

func NewDiaryService() *DiaryService {
	return &DiaryService{
		repo: repository.NewDiaryRepository(),
	}
}

func (s *DiaryService) RegisterDiary(diary models.Diary) error {
	diary.Username = strings.ToLower(diary.Username)

	existing, err := s.repo.FindDiaryByUsername(diary.Username)
	if err == nil && existing.Username != "" {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(diary.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	diary.Password = string(hashedPassword)
	return s.repo.CreateDiary(diary)
}

func (s *DiaryService) Login(username, password string) error {
	username = strings.ToLower(username)

	diary, err := s.repo.FindDiaryByUsername(username)
	if err != nil {
		return errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(diary.Password), []byte(password))
	if err != nil {
		return errors.New("invalid username or password")
	}

	return nil
}

func (s *DiaryService) GetDiary(username string) (models.Diary, error) {
	username = strings.ToLower(username)
	return s.repo.FindDiaryByUsername(username)
}

func (s *DiaryService) AddEntry(username string, entry models.Entry) error {
	username = strings.ToLower(username)

	diary, err := s.repo.FindDiaryByUsername(username)
	if err != nil {
		return err
	}
	if diary.IsLocked {
		return errors.New("diary is locked")
	}

	entry.DateCreated = time.Now()
	entry.ID = len(diary.Entries) + 1

	return s.repo.AddEntry(username, entry)
}

func (s *DiaryService) UpdateEntry(username string, entry models.Entry) error {
	username = strings.ToLower(username)

	diary, err := s.repo.FindDiaryByUsername(username)
	if err != nil {
		return err
	}
	if diary.IsLocked {
		return errors.New("diary is locked")
	}

	return s.repo.UpdateEntry(username, entry)
}

func (s *DiaryService) DeleteEntry(username string, entryID int) error {
	username = strings.ToLower(username)

	diary, err := s.repo.FindDiaryByUsername(username)
	if err != nil {
		return err
	}
	if diary.IsLocked {
		return errors.New("diary is locked")
	}

	return s.repo.DeleteEntry(username, entryID)
}

func (s *DiaryService) DeleteDiary(username string) error {
	username = strings.ToLower(username)

	diary, err := s.repo.FindDiaryByUsername(username)
	if err != nil {
		return err
	}
	if diary.IsLocked {
		return errors.New("diary is locked")
	}

	return s.repo.DeleteDiary(username)
}

func (s *DiaryService) LockDiary(username string) error {
	username = strings.ToLower(username)
	return s.repo.LockDiary(username)
}

func (s *DiaryService) UnlockDiary(username string) error {
	username = strings.ToLower(username)
	return s.repo.UnlockDiary(username)
}
