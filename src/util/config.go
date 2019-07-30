package util

import (
"encoding/json"
"io/ioutil"
"log"
)

type Configuration struct {
	VerificationToken string `json:"verification_token"`
	ClientSecret string `json:"client_secret"`
	OAuthAccessToken string `json:"o_auth_access_token"`
}

func ReadConfigFile (configPath string) Configuration {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var botConfiguration Configuration
	err = json.Unmarshal(data, &botConfiguration)
	if err != nil {
		log.Fatal(err)
	}

	return botConfiguration
}
