// Package api
// Description: This function is used to get the artist information from Spotify
// Request Type: POST
// Input: JSON{artist_id: string}
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"template-go-vercel/api/common"
)

func GetSpotifyArtist(w http.ResponseWriter, r *http.Request) {
	var spotifyAuth common.SpotifyAuth

	//get spotify auth key
	if spotifyAuth.GetKey() == true {
		client := http.Client{}

		//read request body
		type RequestJson struct {
			ArtistId string `json:"artist_id"`
		}

		var requestJson RequestJson

		err := json.NewDecoder(r.Body).Decode(&requestJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/artists/"+requestJson.ArtistId, nil)

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
