package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

func main() {
	authRequest := AuthRequest{ClientID: "d66b2fc2-2a5b-4d37-82e5-5356e4e91cd1",
		ClientSecret: "c6d9c4c3-fa83-4998-a836-1409201dbb7c", GrantType: "client_credentials"}
	jsonPayload, err := json.Marshal(authRequest)
	if err != nil {
		log.Fatalf("error creating json payload")
	}

	httpClient := http.Client{}
	request, err := http.NewRequest(http.MethodPost, "https://openapiuat.airtel.africa/auth/oauth2/token", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("error creating http request")
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "*")

	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatalf("error sending request")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error retrieving response body")
	}

	fmt.Println("response is ", string(body))

}
