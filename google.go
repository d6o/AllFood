package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Google struct {
	searchURL    string
	googleApiKey string
}

type GoogleResult struct {
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Viewport struct {
			Northeast struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"northeast"`
			Southwest struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"viewport"`
	} `json:"geometry"`
	Icon   string `json:"icon"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Photos []struct {
		Height           int      `json:"height"`
		HTMLAttributions []string `json:"html_attributions"`
		PhotoReference   string   `json:"photo_reference"`
		Width            int      `json:"width"`
	} `json:"photos,omitempty"`
	PlaceID      string   `json:"place_id"`
	Reference    string   `json:"reference"`
	Scope        string   `json:"scope"`
	Types        []string `json:"types"`
	Vicinity     string   `json:"vicinity"`
	OpeningHours struct {
		OpenNow     bool          `json:"open_now"`
		WeekdayText []interface{} `json:"weekday_text"`
	} `json:"opening_hours,omitempty"`
	Rating float64 `json:"rating,omitempty"`
}

type GoogleSearch struct {
	HTMLAttributions []interface{}  `json:"html_attributions"`
	NextPageToken    string         `json:"next_page_token"`
	Results          []GoogleResult `json:"results"`
	Status           string         `json:"status"`
}

type GooglePlace struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Result           struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		AdrAddress       string `json:"adr_address"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Viewport struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		Icon   string `json:"icon"`
		ID     string `json:"id"`
		Name   string `json:"name"`
		Photos []struct {
			Height           int      `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos"`
		PlaceID   string   `json:"place_id"`
		Reference string   `json:"reference"`
		Scope     string   `json:"scope"`
		Types     []string `json:"types"`
		URL       string   `json:"url"`
		UtcOffset int      `json:"utc_offset"`
		Vicinity  string   `json:"vicinity"`
	} `json:"result"`
	Status string `json:"status"`
}

func NewGoogle(googleApiKey string) *Google {
	return &Google{
		searchURL:    "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s,%s&radius=%d&key=%s&type=restaurant",
		googleApiKey: googleApiKey,
	}
}

func (g *Google) Search(radius int, lat, lng string) ([]Place, error) {

	url := fmt.Sprintf(g.searchURL, lat, lng, radius, g.googleApiKey)
	urlFinal := url

	status := true

	r := []Place{}

	for status {

		_, body, err := gorequest.New().Get(urlFinal).End()
		if err != nil {
			return nil, errors.New("Error getting Google Places.")
		}

		var gs GoogleSearch
		var result = []byte(body)

		err2 := json.Unmarshal(result, &gs)

		if err2 != nil {
			return nil, err2
		}

		for _, element := range gs.Results {
			f := Place{
				Name:    element.Name,
				PlaceID: element.PlaceID,
				Lat:     element.Geometry.Location.Lat,
				Lng:     element.Geometry.Location.Lng,
			}

			r = append(r, f)
		}

		status = false
		if len(gs.NextPageToken) > 0 {
			status = true
			urlFinal = url + "&pagetoken=" + gs.NextPageToken
		}
	}

	return r, nil
}
