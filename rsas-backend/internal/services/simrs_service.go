package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// BedData adalah struktur contoh data bed dari SIMRS
type BedData struct {
	RoomName  string `json:"room_name"`
	Class     string `json:"class"`
	Available int    `json:"available"`
	Total     int    `json:"total"`
}

// FetchBedAvailability memanggil API SIMRS untuk mendapatkan data ketersediaan tempat tidur
func FetchBedAvailability() ([]BedData, error) {
	apiURL := os.Getenv("SIMRS_API_URL") + "/beds"
	apiKey := os.Getenv("SIMRS_API_KEY")

	// Setup HTTP Client dengan Timeout
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Menambahkan Header API Key jika diperlukan oleh SIMRS
	req.Header.Add("X-API-KEY", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("SIMRS API connection failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SIMRS API returned status: %d", resp.StatusCode)
	}

	var beds []BedData
	if err := json.NewDecoder(resp.Body).Decode(&beds); err != nil {
		return nil, err
	}

	return beds, nil
}
