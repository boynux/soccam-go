package soccam

import (
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    TwitterConsumerKey string
    TwitterConsumerSecret string
    TwitterAccessToken string
    TwitterAccessTokenSecret string
}

func loadConfig(filename string) Configuration {
    file, _ := os.Open(filename)
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)

    if err != nil {
        fmt.Println("error:", err)
    }

    return configuration
}

func GetConfig() Configuration {
    return loadConfig("config.json")
}
