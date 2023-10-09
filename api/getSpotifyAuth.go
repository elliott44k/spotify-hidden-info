package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	/**
	check redis for key
	if key exists, return key for use
	else get key and save for 1 hour and return key
	*/

	//check redis
	client := http.Client{Timeout: 5 * time.Second}
	//client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("REDIS_REST_API_URL")+"/get/spotify_authorization", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("REDIS_REST_API_TOKEN"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	respData, err := io.ReadAll(resp.Body)

	type RedisResponse struct {
		Result string `json:"result"`
	}

	var result RedisResponse
	if err := json.Unmarshal(respData, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	if result.Result != "" {
		fmt.Println(result.Result)
	}

	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	w.Write(respData)

	//
	////We Read the response body on the line below.
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	////Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

}
