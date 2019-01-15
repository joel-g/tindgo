package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// TinderAPI API object
type TinderAPI struct{}

var client = http.Client{}

func main() {
	api := API()
	recs := api.Recommendations()
	for _, r := range recs {
		fmt.Printf("%+v", r.User.Jobs)
	}

}

// API returns a new API client
func API() *TinderAPI {
	return &TinderAPI{}
}

// Recommendations fetches reccomendation objects
func (a TinderAPI) Recommendations() []Recommendation {
	resp := tHTTP("GET", "recs/core", nil)
	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)
	var recs recResponse
	json.Unmarshal(byt, &recs)
	return recs.Data.Results
}

func getToken() string {
	dat, err := ioutil.ReadFile(".tindgo")
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func tHTTP(method string, url string, body []byte) *http.Response {
	req, err := http.NewRequest(method, "https://api.gotinder.com/v2/"+url+"?locale-en", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	token := getToken()
	req.Header.Add("x-auth-token", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

type recResponse struct {
	Data struct {
		Results []Recommendation `json:"results"`
	} `json:"data"`
}

// Recommendation is a user object for a recommended user profile
type Recommendation struct {
	Type string `json:"type"`
	User struct {
		ID        string    `json:"_id"`
		Bio       string    `json:"bio"`
		BirthDate time.Time `json:"birth_date"`
		Name      string    `json:"name"`
		Photos    []struct {
			ID       string `json:"id"`
			CropInfo struct {
				ProcessedByBullseye bool `json:"processed_by_bullseye"`
				UserCustomized      bool `json:"user_customized"`
			} `json:"crop_info"`
			URL            string `json:"url"`
			ProcessedFiles []struct {
				URL    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"processedFiles"`
			FileName  string `json:"fileName"`
			Extension string `json:"extension"`
		} `json:"photos"`
		Gender int `json:"gender"`
		Jobs   []struct {
			Title struct {
				Name string `json:"name"`
			} `json:"title"`
		} `json:"jobs"`
		Schools []interface{} `json:"schools"`
	} `json:"user"`
	Facebook struct {
		CommonConnections []interface{} `json:"common_connections"`
		ConnectionCount   int           `json:"connection_count"`
		CommonInterests   []interface{} `json:"common_interests"`
	} `json:"facebook"`
	Spotify struct {
		SpotifyConnected bool `json:"spotify_connected"`
	} `json:"spotify"`
	DistanceMi  int    `json:"distance_mi"`
	ContentHash string `json:"content_hash"`
	SNumber     int    `json:"s_number"`
	Teaser      struct {
		Type   string `json:"type"`
		String string `json:"string"`
	} `json:"teaser"`
	Teasers []struct {
		Type   string `json:"type"`
		String string `json:"string"`
	} `json:"teasers"`
	Instagram struct {
		LastFetchTime         time.Time `json:"last_fetch_time"`
		CompletedInitialFetch bool      `json:"completed_initial_fetch"`
		Photos                []struct {
			Image     string `json:"image"`
			Thumbnail string `json:"thumbnail"`
			Ts        string `json:"ts"`
			Link      string `json:"link"`
		} `json:"photos"`
		MediaCount     int    `json:"media_count"`
		ProfilePicture string `json:"profile_picture"`
		Username       string `json:"username"`
	} `json:"instagram,omitempty"`
}
