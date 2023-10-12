package api

import (
	"fmt"
	"io"
	"net/http"
	"template-go-vercel/api/common"
)

func GetSpotifySearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fuck")
	var spotifyAuth common.SpotifyAuth

	if spotifyAuth.GetKey() == true {
		client := http.Client{}
		query := r.URL.Query()

		if len(query["q"]) == 0 {
			http.Error(w, "No search params given", http.StatusBadRequest)
			return
		}

		// if no market specified, default to US
		if len(query["market"]) == 0 {
			query["market"] = []string{"US"}
		}
		query.Add("type", "track")

		fmt.Println(query)

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
