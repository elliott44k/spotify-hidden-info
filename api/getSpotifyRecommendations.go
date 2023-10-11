package api

import (
	"fmt"
	"io"
	"net/http"
	"template-go-vercel/api/common"
)

func GetSpotifyRecommendations(w http.ResponseWriter, r *http.Request) {
	var spotifyAuth common.SpotifyAuth

	if spotifyAuth.GetKey() == true {
		client := http.Client{}
		query := r.URL.Query()

		// query must include at least one seed from artists, genres, or tracks
		if len(query["seed_artists"]) == 0 && len(query["seed_genres"]) == 0 && len(query["seed_tracks"]) == 0 {
			http.Error(w, "seed_artists, seed_genres or seed_tracks required", http.StatusBadRequest)
			return
		}

		// if no market specified, default to US
		if len(query["market"]) == 0 {
			query["market"] = []string{"US"}
		}

		req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/recommendations/", nil)
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
		fmt.Println("failed to retrieve spotify recommendations")
	}
}
