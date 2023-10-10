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

		req, err := http.NewRequest("GET", "https://api.spotify.com/v1/audio-features/"+requestJson.TrackId, nil)

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
