// Package api
// Description: This function is used to get search results from Spotify
// Request Type: POST
// Input: JSON {track_name: string, artist_name: string, type: string}
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"template-go-vercel/api/common"
)

func GetSpotifySearch(w http.ResponseWriter, r *http.Request) {
	var spotifyAuth common.SpotifyAuth

	//get spotify auth key
	if spotifyAuth.GetKey() == true {
		client := http.Client{}
		query := r.URL.Query()

		//read request body
		type RequestJson struct {
			TrackName   string `json:"track_name"`
			ArtistsName string `json:"artist_name"`
			Type        string `json:"type"`
			Market      string `json:"market"`
		}

		var requestJson RequestJson

		err := json.NewDecoder(r.Body).Decode(&requestJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Request most include a track name or an artist name
		if requestJson.TrackName == "" && requestJson.ArtistsName == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// if no market specified, default to US
		if requestJson.Market == "" {
			requestJson.Market = "US"
		}

		// Spotify endpoint uses unique "q" request parameter contains a lot of the search filters
		qString := ""
		if requestJson.TrackName != "" {
			qString += "track:" + requestJson.TrackName
		}

		if requestJson.ArtistsName != "" {
			qString += " artist:" + requestJson.ArtistsName
		}
		query.Add("q", qString)
		query.Add("type", requestJson.Type)
		query.Add("market", requestJson.Market)

		req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/search/", nil)
		req.URL.RawQuery = query.Encode()
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
