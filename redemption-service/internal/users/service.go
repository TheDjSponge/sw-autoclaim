package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)


type Service struct{
	db		*database.Queries
	checkUserURL string
}

func NewService(db *database.Queries, checkUserURL string) *Service {
	return &Service{
		db: db,
		checkUserURL: checkUserURL,
	}
}

func (s *Service) RegisterUser(ctx context.Context, hiveID, server, dicordUsername string, discordID int) error {
	valid, gameUID, err := s.CheckUserValidity(hiveID, server)
	if err != nil || !valid{
		return fmt.Errorf("user validation failed: %w", err)
	}

	err = s.db.AddDiscordUser(ctx, database.AddDiscordUserParams{
		Username: dicordUsername,
		DiscordID: int32(discordID),
	})
	if err != nil{
		return fmt.Errorf("failed to save discord profile: %w", err)
	}

	err = s.db.AddUser(ctx, database.AddUserParams{
		HiveID: hiveID,
		Server: server,
		DiscordID: int32(discordID),
		GameUid: int32(gameUID),
	})

	return err
}

func (s *Service) DeleteUser(ctx context.Context, discordID int, hiveID, server string) error {
	result, err := s.db.DeleteUser(ctx, database.DeleteUserParams{DiscordID: int32(discordID), HiveID: hiveID, Server: server})
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("failed to reference any user to delete")
	}
	return err
} 

func (s *Service) GetAllUsers(ctx context.Context) ([]database.User, error){
	users, err := s.db.GetAllUsers(ctx)
	if err != nil {
		return []database.User{}, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil

}

func (s *Service) CheckUserValidity(hive_id string, server string) (bool, int, error) {
	formData := url.Values{
		"country": {"CH"},
		"lang":    {"fr"},
		"server":  {server},
		"hiveid":  {hive_id},
		"coupon":  {"pouetpouet"},
	}
	encodedData := formData.Encode()

	req, err := http.NewRequest("POST", s.checkUserURL, strings.NewReader(encodedData))
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