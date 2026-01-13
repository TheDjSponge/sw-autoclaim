package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) HandleNewUser(w http.ResponseWriter, r *http.Request){
	type newUserInfo struct {
		HiveID          string `json:"hive_id"`
		Server          string `json:"server"`
		DiscordID       int    `json:"discord_id"`
		DiscordUsername string `json:"discord_username"`
	}
	var userInfo newUserInfo

	
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		RespondWithMessage(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err := h.userService.RegisterUser(r.Context(), userInfo.HiveID, userInfo.Server, userInfo.DiscordUsername, userInfo.DiscordID)
	if err != nil{
		RespondWithMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, nil)
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	type deleteUserInfo struct {
		DiscordID int    `json:"discord_id"`
		HiveID    string `json:"hive_id"`
		Server    string `json:"server"`
	}
	var userInfo deleteUserInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInfo)
	if err != nil {
		log.Printf("Error while trying decode user info before deletion: %w", err)
		RespondWithMessage(w, http.StatusBadRequest, "error while trying to delete user")
		return
	}

	err = h.userService.DeleteUser(r.Context(),userInfo.DiscordID, userInfo.HiveID, userInfo.Server)
	if err != nil {
		log.Printf("Error while trying to delete user from database: %w", err)
		RespondWithMessage(w, http.StatusInternalServerError, "error while trying to delete user")
		return
	}

	RespondWithMessage(w, http.StatusOK, "User deleted successfully.")
}


func (h *Handler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		RespondWithMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, users)
}