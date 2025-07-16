// service/diary_service_test.go
package service

import (
	"diary-app/models"
	"testing"
)

func TestRegisterDiary_Success(t *testing.T) {
	t.Log(">>> Running TestRegisterDiary_Success")

	service := NewDiaryService()
	username := "testuser123"

	_ = service.DeleteDiary(username)
	t.Cleanup(func() {
		_ = service.DeleteDiary(username)
	})

	diary := models.Diary{
		Username: username,
		Password: "mypassword123",
	}

	err := service.RegisterDiary(diary)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	t.Log(">>> Test passed.")
}

