package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
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
