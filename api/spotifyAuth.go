package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type SpotifyAuth struct {
	Key string
}

func (spot SpotifyAuth) getKey() bool {
	//check redis
	spot.Key = spot.getRedis()

	// no key in Redis, get key from Spotify
	if spot.Key != "" {
		return true
	} else {
		client := http.Client{Timeout: 5 * time.Second}

		jsonString := map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     os.Getenv("SPOTIFY_CLIENT_ID"),
			"client_secret": os.Getenv("SPOTIFY_CLIENT_SECRET")}
		jsonEncoded := new(bytes.Buffer)
		json.NewEncoder(jsonEncoded).Encode(jsonString)
		req, err := http.NewRequest(http.MethodPost, os.Getenv("https://accounts.spotify.com/api/token")+"/get/spotify_authorization", jsonEncoded)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		respData, err := io.ReadAll(resp.Body)

		type SpotifyKeyResponse struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			ExpiresIn   int    `json:"expires_in"`
		}

		var jsonData SpotifyKeyResponse
		if err := json.Unmarshal(respData, &jsonData); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Error in JSON unmarshal")
		}

		if jsonData.AccessToken == "" {
			spot.Key = ""
			return false
		} else {
			spot.Key = jsonData.AccessToken
			spot.setRedis(jsonData.AccessToken)
			return true
		}
	}

	return false
}

/*
*
Get Redis Value
*/
func (spot SpotifyAuth) getRedis() string {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("REDIS_REST_API_URL")+"/get/spotify_authorization", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("REDIS_REST_API_TOKEN"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	respData, err := io.ReadAll(resp.Body)

	type RedisResponse struct {
		Result string `json:"result"`
	}

	var jsonData RedisResponse
	if err := json.Unmarshal(respData, &jsonData); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Error in JSON unmarshal")
		return ""
	}

	if jsonData.Result != "" {
		return jsonData.Result
	}

	return ""
}

func (spot SpotifyAuth) setRedis(value string) bool {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("REDIS_REST_API_URL")+"/set/spotify_authorization/"+value+"/ex/3600", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("REDIS_REST_API_TOKEN"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	respData, err := io.ReadAll(resp.Body)

	type RedisResponse struct {
		Result string `json:"result"`
	}

	var jsonData RedisResponse
	if err := json.Unmarshal(respData, &jsonData); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Error in JSON unmarshal")
	}

	if jsonData.Result != "OK" {
		return true
	}

	return false
}

// empty handler function for vercel compile
func Handler(w http.ResponseWriter, r *http.Request) {

}
