package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
)

type Facebook struct {
	AppID       string
	AppSecret   string
	AccessToken string
	TokenURL    string
	PlaceURL    string
}

type FacebookPlaceSimple struct {
	OverallStarRating float64 `json:"overall_star_rating"`
	RatingCount       int     `json:"rating_count"`
	Name              string  `json:"name"`
	Location          struct {
		City      string  `json:"city"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		LocatedIn string  `json:"located_in"`
		Longitude float64 `json:"longitude"`
		State     string  `json:"state"`
	} `json:"location"`
	Category     string `json:"category"`
	CategoryList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category_list"`
	ID    string `json:"id"`
	About string `json:"about,omitempty"`
}

type FacebookPlaceList struct {
	Data   []FacebookPlaceSimple `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

func NewFacebook(AppID, AppSecret string) *Facebook {
	return &Facebook{
		AppID:     AppID,
		AppSecret: AppSecret,
		TokenURL:  "https://graph.facebook.com/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials",
		PlaceURL:  "https://graph.facebook.com/search?type=place&center=%s,%s&distance=%s&access_token=%s&fields=overall_star_rating,rating_count,name,about,location,category,category_list",
	}
}

func (f *Facebook) GetToken() error {
	url := fmt.Sprintf(f.TokenURL, f.AppID, f.AppSecret)
	_, body, err := gorequest.New().Get(url).End()
	if err != nil {
		return errors.New("Error getting Facebook AccessToken.")
	}

	f.AccessToken = strings.Replace(body, "access_token=", "", -1)
	return nil
}

func (f *Facebook) Search(radius int, lat, lng string) ([]Place, error) {
	url := fmt.Sprintf(f.PlaceURL, lat, lng, radius, f.AccessToken)
	urlFinal := url
	status := true
	r := []Place{}
	for status {
		_, body, err := gorequest.New().Get(urlFinal).End()
		if err != nil {

		}

		var fpl FacebookPlaceList

		var result = []byte(body)

		err2 := json.Unmarshal(result, &fpl)

		if err2 != nil {
			return nil, err2
		}

		for _, element := range fpl.Data {
			f := Place{
				Name:    element.Name,
				PlaceID: element.ID,
				Lat:     element.Location.Latitude,
				Lng:     element.Location.Longitude,
			}

			if element.Category == "Restaurant" || element.Category == "Bar" {
				r = append(r, f)
			}
		}

		status = false
		if len(fpl.Paging.Next) > 0 {
			status = true
			urlFinal = fpl.Paging.Next
		}
	}
	return r, nil
}
