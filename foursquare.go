package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type FourSquare struct {
	searchURL    string
	clientID     string
	clientSecret string
}

type FourSquareResult struct {
	Meta struct {
		Code      int    `json:"code"`
		RequestID string `json:"requestId"`
	} `json:"meta"`
	Response struct {
		SuggestedFilters struct {
			Header  string `json:"header"`
			Filters []struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"filters"`
		} `json:"suggestedFilters"`
		HeaderLocation            string `json:"headerLocation"`
		HeaderFullLocation        string `json:"headerFullLocation"`
		HeaderLocationGranularity string `json:"headerLocationGranularity"`
		TotalResults              int    `json:"totalResults"`
		SuggestedBounds           struct {
			Ne struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"ne"`
			Sw struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"sw"`
		} `json:"suggestedBounds"`
		Groups []struct {
			Type  string `json:"type"`
			Name  string `json:"name"`
			Items []struct {
				Reasons struct {
					Count int `json:"count"`
					Items []struct {
						Summary    string `json:"summary"`
						Type       string `json:"type"`
						ReasonName string `json:"reasonName"`
					} `json:"items"`
				} `json:"reasons"`
				Venue struct {
					ID      string `json:"id"`
					Name    string `json:"name"`
					Contact struct {
					} `json:"contact"`
					Location struct {
						Address        string  `json:"address"`
						Lat            float64 `json:"lat"`
						Lng            float64 `json:"lng"`
						LabeledLatLngs []struct {
							Label string  `json:"label"`
							Lat   float64 `json:"lat"`
							Lng   float64 `json:"lng"`
						} `json:"labeledLatLngs"`
						Distance         int      `json:"distance"`
						PostalCode       string   `json:"postalCode"`
						Cc               string   `json:"cc"`
						City             string   `json:"city"`
						State            string   `json:"state"`
						Country          string   `json:"country"`
						FormattedAddress []string `json:"formattedAddress"`
					} `json:"location"`
					Categories []struct {
						ID         string `json:"id"`
						Name       string `json:"name"`
						PluralName string `json:"pluralName"`
						ShortName  string `json:"shortName"`
						Icon       struct {
							Prefix string `json:"prefix"`
							Suffix string `json:"suffix"`
						} `json:"icon"`
						Primary bool `json:"primary"`
					} `json:"categories"`
					Verified bool `json:"verified"`
					Stats    struct {
						CheckinsCount int `json:"checkinsCount"`
						UsersCount    int `json:"usersCount"`
						TipCount      int `json:"tipCount"`
					} `json:"stats"`
					Price struct {
						Tier     int    `json:"tier"`
						Message  string `json:"message"`
						Currency string `json:"currency"`
					} `json:"price"`
					Rating           float64 `json:"rating"`
					RatingColor      string  `json:"ratingColor"`
					RatingSignals    int     `json:"ratingSignals"`
					AllowMenuURLEdit bool    `json:"allowMenuUrlEdit"`
					BeenHere         struct {
						Count                int  `json:"count"`
						UnconfirmedCount     int  `json:"unconfirmedCount"`
						Marked               bool `json:"marked"`
						LastCheckinExpiredAt int  `json:"lastCheckinExpiredAt"`
					} `json:"beenHere"`
					Hours struct {
						Status         string `json:"status"`
						IsOpen         bool   `json:"isOpen"`
						IsLocalHoliday bool   `json:"isLocalHoliday"`
					} `json:"hours"`
					Photos struct {
						Count  int           `json:"count"`
						Groups []interface{} `json:"groups"`
					} `json:"photos"`
					HereNow struct {
						Count   int    `json:"count"`
						Summary string `json:"summary"`
						Groups  []struct {
							Type  string        `json:"type"`
							Name  string        `json:"name"`
							Count int           `json:"count"`
							Items []interface{} `json:"items"`
						} `json:"groups"`
					} `json:"hereNow"`
				} `json:"venue"`
				Tips []struct {
					ID           string `json:"id"`
					CreatedAt    int    `json:"createdAt"`
					Text         string `json:"text"`
					Type         string `json:"type"`
					CanonicalURL string `json:"canonicalUrl"`
					Likes        struct {
						Count   int           `json:"count"`
						Groups  []interface{} `json:"groups"`
						Summary string        `json:"summary"`
					} `json:"likes"`
					LogView       bool `json:"logView"`
					AgreeCount    int  `json:"agreeCount"`
					DisagreeCount int  `json:"disagreeCount"`
					Todo          struct {
						Count int `json:"count"`
					} `json:"todo"`
					User struct {
						ID        string `json:"id"`
						FirstName string `json:"firstName"`
						LastName  string `json:"lastName"`
						Gender    string `json:"gender"`
						Photo     struct {
							Prefix string `json:"prefix"`
							Suffix string `json:"suffix"`
						} `json:"photo"`
						IsAnonymous bool `json:"isAnonymous"`
					} `json:"user"`
				} `json:"tips"`
				ReferralID string `json:"referralId"`
			} `json:"items"`
		} `json:"groups"`
	} `json:"response"`
}

func NewFourSquare(clientID, clientSecret string) *FourSquare {
	return &FourSquare{
		clientID:     clientID,
		clientSecret: clientSecret,
		searchURL:    "https://api.foursquare.com/v2/venues/explore?ll=%s,%s&intent=browse&radius=%s&limit=50&client_id=%s&client_secret=%s&v=%s&offset=%d&categoryId=4d4b7105d754a06374d81259",
	}
}

func (f *FourSquare) Search(radius int, lat, lng string) ([]Place, error) {
	t := time.Now()
	offset := 0
	status := true
	r := []Place{}

	for status {

		url := fmt.Sprintf(f.searchURL, lat, lng, radius, f.clientID, f.clientSecret, t.Format("20060102"), offset)

		_, body, err := gorequest.New().Get(url).End()
		if err != nil {
			return nil, errors.New("Error getting FourSquare Places.")
		}

		var fs FourSquareResult
		var result = []byte(body)

		err2 := json.Unmarshal(result, &fs)
		if err2 != nil {
			return nil, err2
		}

		for _, element := range fs.Response.Groups {

			if len(element.Items) < 50 {
				status = false
			}

			for _, element2 := range element.Items {

				f := Place{
					Name:    element2.Venue.Name,
					PlaceID: element2.Venue.ID,
					Lat:     element2.Venue.Location.Lat,
					Lng:     element2.Venue.Location.Lng,
				}
				r = append(r, f)

			}
		}
		offset = offset + 50
	}
	return r, nil
}
