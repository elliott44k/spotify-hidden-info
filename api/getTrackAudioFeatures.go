package handler

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetTrackAudioFeatures(w http.ResponseWriter, r *http.Request) {

	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/audio-features", nil)

	req.Header.Add("Authorization", "Bearer"+" 1POdFZRZbvb...qqillRxMr2z")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	//w.WriteHeader(http.StatusCreated)
	//w.Header().Set("Content-Type", "application/json")
	//resp := make(map[string]string)

	//// Response API Location
	//urlLocationApi := "https://api.spotify.com/v1/audio-features" + "?apikey=" + os.Getenv("SPOTIFY_CLIENT_ID")
	//fmt.Println("Location API URL:", urlLocationApi)
	//userLocationResponse, _ := http.Get(urlLocationApi)
	//// Get the Location body
	//userLocationBody, _ := io.ReadAll(userLocationResponse.Body)
	//userLocationJson := string(userLocationBody)
	//fmt.Println("Location API Json String:", userLocationJson)
	//var userLocation map[string]interface{}
	//json.Unmarshal([]byte(userLocationBody), &userLocation)
	//
	//resp["latitude"] = fmt.Sprint(userLocation["latitude"].(float64))
	//resp["longitude"] = fmt.Sprint(userLocation["longitude"].(float64))
	//
	//resp["github"] = "https://github.com/riccardogiorato/template-go-vercel/blob/main/api/myweather.go"
	//
	//// Response API Weather
	//urlWeatherApi := "https://api.openweathermap.org/data/2.5/weather?lat=" + resp["latitude"] + "&lon=" + resp["longitude"] + "&appid=" + os.Getenv("API_KEY_OPENWEATHER")
	//fmt.Println("Weather API Json String:", urlWeatherApi)
	//weatherApiResponse, _ := http.Get(urlWeatherApi)
	//weatherApiBody, _ := io.ReadAll(weatherApiResponse.Body)
	//weatherApiJson := string(weatherApiBody)
	//
	//resp["weather"] = weatherApiJson
	//
	//jsonResp, err := json.Marshal(resp)
	//if err != nil {
	//	fmt.Println("Error happened in JSON marshal. Err:", err)
	//} else {
	//	w.Write(jsonResp)
	//}
}
