package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)

func (cfg *apiConfig) HandleNewUser(w http.ResponseWriter, r *http.Request) {
	type newUserInfo struct {
		HiveID          string `json:"hive_id"`
		Server          string `json:"server"`
		DiscordID       int    `json:"discord_id"`
		DiscordUsername string `json:"discord_username"`
	}
	var userInfo newUserInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInfo)
	if err != nil {
		log.Printf("Got error when decoding request body: %v", err)
		respondWithMessage(w, http.StatusBadRequest, "couldn't decode user information, check format in docs.")
		return
	}
	valid, userID, err := cfg.CheckUserValidity(userInfo.HiveID, userInfo.Server)
	if err != nil {
		log.Printf("Got error when checking user validity: %v", err)
		respondWithMessage(w, http.StatusInternalServerError, "error when checking user validity")
		return
	}
	if !valid {
		log.Printf("Invalid user %v", userInfo.HiveID)
		respondWithMessage(w, http.StatusNotFound, "user is not valid")
		return
	}

	err = cfg.database.AddDiscordUser(r.Context(), database.AddDiscordUserParams{Username: userInfo.DiscordUsername, DiscordID: int32(userInfo.DiscordID)})
	if err != nil {
		log.Printf("Couldn't add user to discord users db, got error: %v", err)
		respondWithMessage(w, http.StatusInternalServerError, "couldn't add user to database")
		return
	}

	err = cfg.database.AddUser(r.Context(), database.AddUserParams{HiveID: userInfo.HiveID, Server: userInfo.Server, DiscordID: int32(userInfo.DiscordID), GameUid: int32(userID)})
	if err != nil {
		log.Printf("Couldn't add user to users db, got error: %v", err)
		respondWithMessage(w, http.StatusInternalServerError, "couldn't add user to database")
		return
	}

	respondWithJSON(w, http.StatusCreated, nil)
}

func (cfg *apiConfig) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	type deleteUserInfo struct {
		DiscordID int    `json:"discord_id"`
		HiveID    string `json:"hive_id"`
		Server    string `json:"server"`
	}
	var userInfo deleteUserInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInfo)
	if err != nil {
		log.Printf("Error while trying decode user info before deletion: %v", err)
		respondWithMessage(w, http.StatusBadRequest, "error while trying to delete user")
		return
	}

	result, err := cfg.database.DeleteUser(r.Context(), database.DeleteUserParams{DiscordID: int32(userInfo.DiscordID), HiveID: userInfo.HiveID, Server: userInfo.Server})
	if err != nil {
		log.Printf("Error while trying to delete user from database: %v", err)
		respondWithMessage(w, http.StatusInternalServerError, "error while trying to delete user")
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error while trying to fetch number of affected rows: %v", err)
		respondWithMessage(w, http.StatusInternalServerError, "error while trying to delete user")
		return
	}

	if rowsAffected == 0 {
		respondWithMessage(w, http.StatusInternalServerError, "error while trying to delete user")
		return
	}

	respondWithMessage(w, http.StatusOK, "User deleted successfully.")
}

func (cfg *apiConfig) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := cfg.database.GetAllUsers(r.Context())
	if err != nil {
		log.Printf("Couldn't retrieve users from database, got error: %v", err)
		respondWithMessage(w, http.StatusInternalServerError, "can't get users")
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiConfig) CheckUserValidity(hive_id string, server string) (bool, int, error) {
	formData := url.Values{
		"country": {"CH"},
		"lang":    {"fr"},
		"server":  {server},
		"hiveid":  {hive_id},
		"coupon":  {"pouetpouet"},
	}
	encodedData := formData.Encode()

	req, err := http.NewRequest("POST", cfg.checkUserAPIURL, strings.NewReader(encodedData))
	if err != nil {
		fmt.Println("throwing at request creation")
		return false, -1, err
	}

	req.Header.Set("Referer", "https://event.withhive.com/ci/smon/evt_coupon")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	type userDataPayload struct {
		UID        int    `json:"uid"`
		ServerName string `json:"server_name"`
		WizardName string `json:"wizard_name"`
	}

	type responsePayload struct {
		RetCode  int             `json:"retCode"`
		RetMsg   string          `json:"retMsg"`
		UserData userDataPayload `json:"userData"`
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("throwing at request")
		return false, -1, err
	}
	defer resp.Body.Close()
	var respPayload responsePayload
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respPayload)
	if err != nil {
		fmt.Println("throwing at decoding")
		return false, -1, err
	}
	// display response payload
	if respPayload.RetCode != 100 {
		return false, -1, fmt.Errorf("check user endpoint hit failed with code: %v", respPayload)
	}
	return true, respPayload.UserData.UID, nil
}
