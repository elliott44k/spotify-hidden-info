// Package common
// Path: api/common/spotifyAuth.go
// Description: SpotifyAuth Struct allows for Spotify API Authorization
package common

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
	"context"
	"github.com/redis/go-redis/v9"
)

type SpotifyAuth struct {
	Key string
}

func (spot *SpotifyAuth) GetKey() bool {
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
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-15401.c294.ap-northeast-1-2.ec2.redns.redis-cloud.com:15401",
		Username: "default",
		Password: os.Getenv("REDIS_DATABASE_PASSWORD"),
		DB:       0,
	})

	result, err := rdb.Get(ctx, "spotify_authorization").Result()

	if err != nil {
		panic(err)
	}


	client := redis.NewClient(opt)

	ctx := context.Background()

	val, err := client.Get(ctx, "spotify_authorization").Result()
	if err != nil {
		panic(err)
	}

	if val != "" {
		return val
	}

	return ""
}

func (spot *SpotifyAuth) setRedis(value string) bool {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-15401.c294.ap-northeast-1-2.ec2.redns.redis-cloud.com:15401",
		Username: "default",
		Password: os.Getenv("REDIS_DATABASE_PASSWORD"),
		DB:       0,
	})

	err = rdb.Set(ctx, "spotify_authorization", value, 3600).Err()
	if err == nil {
		return true
	}

	return false
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//handler function required for vercel compile
}
