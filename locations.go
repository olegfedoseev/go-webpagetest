package webpagetest

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// http://webpagetest.org/getLocations.php?f=json
// https://sites.google.com/a/webpagetest.org/docs/advanced-features/webpagetest-restful-apis#TOC-Location-information

/*
"Dulles_MotoG": {
  "Label": "Dulles, VA USA (Android, iOS 9)",
  "location": "Dulles_MotoG",
  "Browsers": "Motorola G - Chrome,Motorola G - Chrome Canary,Motorola G - Chrome Beta,Motorola G - Chrome Dev,Motorola G - UC Browser,Motorola G - UC Mini,Motorola G - Opera Mini,Motorola G - Chrome,Motorola G - Chrome Canary,Motorola G - Chrome Beta,Motorola G - Chrome Dev,Motorola G - UC Browser,Motorola G - UC Mini,Motorola G - Opera Mini",
  "status": "OK",
  "relayServer": null,
  "relayLocation": null,
  "labelShort": "Dulles, VA",
  "group": "Mobile Devices",
  "PendingTests": {
    "p1": 0,
    ...
    "p9": 0,
    "Total": 5,
    "HighPriority": 0,
    "LowPriority": 0,
    "Testing": 5,
    "Idle": 19
  }
},
*/

type jsonLocation struct {
	Label      string `json:"Label"`      // "Dulles, VA USA (Android, iOS 9)"
	LabelShort string `json:"labelShort"` // "Dulles, VA"
	Location   string `json:"location"`   // "Dulles_MotoG"
	Browsers   string `json:"Browsers"`   // "Motorola G - Chrome,Motorola G - Chrome Canary,...",
	Status     string `json:"status"`     // "OK"
	Group      string `json:"group"`      // "Mobile Devices"

	Default bool `json:"default"`

	RelayServer   string `json:"relayServer"`
	RelayLocation string `json:"relayLocation"`

	PendingTests map[string]int `json:"PendingTests"`
}

type jsonLocations struct {
	StatusCode int    `json:"statusCode"` // 200
	StatusText string `json:"statusText"` // "Ok"

	Data map[string]jsonLocation `json:"data"`
}

// Location is where your agents are
type Location struct {
	Label      string
	LabelShort string
	Location   string
	Browsers   []string
	Status     string

	Default bool

	RelayServer   string
	RelayLocation string

	PendingTests map[string]int
}

// Locations grouped by Group
type Locations map[string][]Location

// GetLocations will retrieve all available locations from server
// You can request a list of locations as well as the number of pending tests for each
func (c *Client) GetLocations() (*Locations, error) {
	body, err := c.query("/getLocations.php", url.Values{"f": []string{"json"}})
	if err != nil {
		return nil, err
	}

	var locations jsonLocations
	if err = json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	if locations.StatusCode != 200 {
		return nil, fmt.Errorf("Status != 200: %v", locations.StatusText)
	}

	result := make(Locations, 0)
	for _, l := range locations.Data {
		if _, ok := result[l.Group]; !ok {
			result[l.Group] = make([]Location, 0)
		}

		result[l.Group] = append(result[l.Group], Location{
			Label:         l.Label,
			LabelShort:    l.LabelShort,
			Location:      l.Location,
			Browsers:      strings.Split(l.Browsers, ","),
			Status:        l.Status,
			Default:       l.Default,
			RelayServer:   l.RelayServer,
			RelayLocation: l.RelayLocation,
			PendingTests:  l.PendingTests,
		})
	}

	return &result, nil
}
