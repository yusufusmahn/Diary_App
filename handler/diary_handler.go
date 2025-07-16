package handler

import (
	"diary-app/dto"
	"diary-app/mapper"
	"diary-app/middleware"
	"diary-app/service"
	"diary-app/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var diaryService = service.NewDiaryService()

func RegisterDiary(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterDiaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	diary := mapper.ToDiaryModel(req)
	err := diaryService.RegisterDiary(diary)
	if err != nil {
		utils.RespondError(w, http.StatusConflict, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Diary registered successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	err := diaryService.Login(req.Username, req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateJWT(req.Username)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

func GetDiary(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	diary, err := diaryService.GetDiary(username)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	response := mapper.ToDiaryResponse(diary)
	utils.RespondJSON(w, http.StatusOK, response)
}

func AddEntry(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	var req dto.EntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	entry := mapper.ToEntryModel(req)
	err := diaryService.AddEntry(username, entry)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Entry added successfully")
}

func UpdateEntry(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	vars := mux.Vars(r)
	entryIDStr := vars["id"]

	entryID, err := strconv.Atoi(entryIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid entry ID")
		return
	}

	var req dto.EntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	entry := mapper.ToEntryModelWithID(entryID, req)
	err = diaryService.UpdateEntry(username, entry)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Entry updated successfully")
}


func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	vars := mux.Vars(r)
	entryIDStr := vars["id"]

	entryID, err := strconv.Atoi(entryIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid entry ID")
		return
	}

	err = diaryService.DeleteEntry(username, entryID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Entry deleted successfully")
}


func DeleteDiary(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	err := diaryService.DeleteDiary(username)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Diary deleted successfully")
}

func LockDiary(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	err := diaryService.LockDiary(username)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Diary locked successfully")
}

func UnlockDiary(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsernameFromContext(r)

	err := diaryService.UnlockDiary(username)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Diary unlocked successfully")
}
