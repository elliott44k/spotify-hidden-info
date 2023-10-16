// Package api
// Description: This function is used to get song detailed audio analysis information from Spotify
// Request Type: POST
// Input: JSON {track_id: string}
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"template-go-vercel/api/common"
)

func GetSpotifyTrackAudioFeatures(w http.ResponseWriter, r *http.Request) {
	var spotifyAuth common.SpotifyAuth

	//get spotify auth key
	if spotifyAuth.GetKey() == true {
		client := http.Client{}

		//read request body
		type RequestJson struct {
			TrackId string `json:"track_id"`
		}

		var requestJson RequestJson

		err := json.NewDecoder(r.Body).Decode(&requestJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Request most include a track id
		if requestJson.TrackId == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Spotify endpoint request format: https://api.spotify.com/v1/audio-features/{id}
		req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/audio-features/"+requestJson.TrackId, nil)

		req.Header.Add("Authorization", "Bearer "+spotifyAuth.Key)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
		}

		//read response body, return
		respBody, err := io.ReadAll(resp.Body)
		w.Write(respBody)
	} else {
		fmt.Println("failed to retrieve spotify auth")
	}
}
