package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"template-go-vercel/api/common"
)

func GetSpotifyRecommendations(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fuck")
	var spotifyAuth common.SpotifyAuth

	if spotifyAuth.GetKey() == true {
		client := http.Client{}
		query := r.URL.Query()

		//read request body
		type RequestJson struct {
			Market                  string `json:"market"`
			Seed_Artists            string `json:"seed_artists"`
			Seed_Genres             string `json:"seed_genres"`
			Seed_Tracks             string `json:"seed_tracks"`
			Limit                   string `json:"limit"`
			Min_Key                 string `json:"min_key"`
			Max_Key                 string `json:"max_key"`
			Target_Key              string `json:"target_key"`
			Min_Acousticness        string `json:"min_acousticness"`
			Max_Acousticness        string `json:"max_acousticness"`
			Target_Acousticness     string `json:"target_acousticness"`
			Min_Instrumentalness    string `json:"min_instrumentalness"`
			Max_Instrumentalness    string `json:"max_instrumentalness"`
			Target_Instrumentalness string `json:"target_instrumentalness"`
			Target_Mode             string `json:"target_mode"`
			Min_Tempo               string `json:"min_tempo"`
			Max_Tempo               string `json:"max_tempo"`
			Target_Tempo            string `json:"target_tempo"`
			Min_Time_Signature      string `json:"min_time_signature"`
			Max_Time_Signature      string `json:"max_time_signature"`
			Target_Time_Signature   string `json:"target_time_signature"`
			Min_Duration_ms         string `json:"min_duration"`
			Max_Duration_ms         string `json:"max_duration"`
			Target_Duration_ms      string `json:"target_duration"`
			Min_Energy              string `json:"min_energy"`
			Max_Energy              string `json:"max_energy"`
			Target_Energy           string `json:"target_energy"`
			Min_Danceability        string `json:"min_danceability"`
			Max_Danceability        string `json:"max_danceability"`
			Target_Danceability     string `json:"target_danceability"`
			Min_Liveness            string `json:"min_liveness"`
			Max_Liveness            string `json:"max_liveness"`
			Target_Liveness         string `json:"target_liveness"`
			Min_Speechiness         string `json:"min_speechiness"`
			Max_Speechiness         string `json:"max_speechiness"`
			Target_Speechiness      string `json:"target_speechiness"`
			Min_Popularity          string `json:"min_popularity"`
			Max_Popularity          string `json:"max_popularity"`
			Target_Popularity       string `json:"target_popularity"`
			Min_Loudness            string `json:"min_loudness"`
			Max_Loudness            string `json:"max_loudness"`
			Target_Loudness         string `json:"target_loudness"`
		}

		var requestJson RequestJson

		err := json.NewDecoder(r.Body).Decode(&requestJson)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// query must include at least one seed from artists, genres, or tracks
		if requestJson.Seed_Artists == "" && requestJson.Seed_Genres == "" && requestJson.Seed_Tracks == "" {
			http.Error(w, "seed_artists, seed_genres or seed_tracks required", http.StatusBadRequest)
			return
		}

		// add query params to request
		fields := reflect.TypeOf(requestJson)
		values := reflect.ValueOf(requestJson)
		num := fields.NumField()
		for i := 0; i < num; i++ {
			field := fields.Field(i)
			value := values.Field(i)
			if value.String() != "" {
				query.Add(strings.ToLower(field.Name), value.String())
			}
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
