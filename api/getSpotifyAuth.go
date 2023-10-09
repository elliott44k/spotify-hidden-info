package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetSpotifyAuth(w http.ResponseWriter, r *http.Request) {
	var spotifyAuth SpotifyAuth
	
	if spotifyAuth.getKey() == true {
		jsonString := map[string]string{"spotify_auth": spotifyAuth.Key}
		jsonResp, err := json.Marshal(jsonString)
		if err != nil {
			fmt.Fprintf(w, "Error happened in JSON Encode. Err: %s", err)
		} else {
			w.Write(jsonResp)
		}
	}
}

type SpotifyAuth struct {
	Key string
}

func (spot *SpotifyAuth) getKey() bool {
	//check redis
	spot.Key = spot.getRedis()

	// no key in Redis, get key from Spotify
	if spot.Key != "" {
		return true
	} else {
		client := http.Client{Timeout: 5 * time.Second}

		data := url.Values{}
		data.Set("grant_type", "client_credentials")
		data.Set("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
		data.Set("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))
		encodedData := data.Encode()
		req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(encodedData))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		respData, err := io.ReadAll(resp.Body)

		type SpotifyKeyResponse struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			ExpiresIn   string `json:"expires_in"`
		}

		var jsonData SpotifyKeyResponse
		if err := json.Unmarshal(respData, &jsonData); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Error in JSON unmarshal 1")
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
Get Redis Value
*/
func (spot *SpotifyAuth) getRedis() string {
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
		fmt.Println("Error in JSON unmarshal 2")
		return ""
	}

	if jsonData.Result != "" {
		return jsonData.Result
	}

	return ""
}

func (spot *SpotifyAuth) setRedis(value string) bool {
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
		fmt.Println("Error in JSON unmarshal 3")
	}

	if jsonData.Result != "OK" {
		return true
	}

	return false
}
