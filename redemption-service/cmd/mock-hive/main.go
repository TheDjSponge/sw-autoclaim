package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Standard Hive Response Structure
type HiveResponse struct {
	RetCode  any         `json:"retCode"`
	RetMsg   string      `json:"retMsg"`
	UserData interface{} `json:"userData,omitempty"` // omitempty handles the claim response which has no user data
}

func main() {
	mux := http.NewServeMux()

	// --- ENDPOINT: User Validation ---
	mux.HandleFunc("/api/check", func(w http.ResponseWriter, r *http.Request) {
		hiveID := r.FormValue("hiveid")
		log.Printf("[CHECK] HiveID: %s", hiveID)

		if hiveID == "non_existent_user" {
			jsonResponse(w, 503, "User not found", nil)
			return
		}

		jsonResponse(w, 100, "", map[string]interface{}{
			"uid":         12345,
			"server_name": r.FormValue("server"),
			"wizard_name": "TestSummoner",
		})
	})

	// --- ENDPOINT: Coupon Claiming ---
	mux.HandleFunc("/api/claim", func(w http.ResponseWriter, r *http.Request) {
		hiveID := r.FormValue("hiveid")
		code := r.FormValue("coupon")
		log.Printf("[CLAIM] HiveID: %s | Code: %s", hiveID, code)

		// logic: Trigger different errors based on the coupon code string
		switch code {
		case "EXPIRED_CODE":
			jsonResponse(w, "(H306)", "This coupon has expired.", nil)
		case "USED_CODE":
			jsonResponse(w, "(H304)", "You have already claimed this coupon.", nil)
		case "INVALID_CODE":
			jsonResponse(w, "(H302)", "This coupon is invalid.", nil)
		case "SERVER_ERROR":
			w.WriteHeader(http.StatusInternalServerError)
		default:
			// Happy Path
			if hiveID == "deleted_user"{
				jsonResponse(w, 503, "Invalid hive ID", nil)
			}
			jsonResponse(w, 100, "Your reward has been sent to your inbox.", nil)
		}
	})

	log.Println("🚀 Mock Hive API starting on :8081...")
	log.Fatal(http.ListenAndServe(":8081", mux))
}

// Helper to send JSON responses consistently
func jsonResponse(w http.ResponseWriter, code any, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HiveResponse{
		RetCode:  code,
		RetMsg:   msg,
		UserData: data,
	})
}