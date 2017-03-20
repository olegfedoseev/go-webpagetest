package webpagetest

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// getTesters.php

/*
"Dulles_MotoG": {
  "elapsed": 0,
  "status": "OK",
  "testers": [
    {
      "id": "Dedicated_MotoG_1-192.168.0.191",
      "pc": "Dedicated_MotoG_1",
      "ip": "192.168.0.191",

      "ec2": "",
      "version": null,
      "freedisk": null,
      "ie": null,
      "winver": "",
      "isWinServer": "",
      "isWin64": "",
      "dns": "",
      "GPU": null,
      "offline": null,
      "screenwidth": "",
      "screenheight": "",

      "rebooted": false,
      "errors": 14,
      "elapsed": 0,
      "last": 46,
      "busy": 0
    },
    {
      "id": "VM1-01-192.168.10.43",
      "pc": "VM1-01",
      "ec2": "",
      "ip": "192.168.10.43",
      "version": "2.19.0.334",
      "freedisk": "14.729",
      "ie": null,
      "winver": "6.3",
      "isWinServer": "1",
      "isWin64": "1",
      "dns": "192.168.0.1",
      "GPU": "1",
      "offline": null,
      "screenwidth": "1920",
      "screenheight": "1200",
      "rebooted": false,
      "cpu": 28,
      "errors": 6,
      "elapsed": 0,
      "last": 109,
      "busy": 0
    },
    {
      "id": "i-538f6ecb",
      "pc": "IP-AC1F259B",
      "ec2": "i-538f6ecb",
      "ip": "174.129.51.11",
      "version": "2.19.0.334",
      "freedisk": "5.428",
      "ie": null,
      "winver": "6.1",
      "isWinServer": "1",
      "isWin64": "1",
      "dns": "172.31.0.2",
      "GPU": "0",
      "offline": null,
      "screenwidth": "1280",
      "screenheight": "1024",
      "rebooted": false,
      "cpu": 90,
      "errors": 2,
      "elapsed": 0,
      "last": 1,
      "busy": 0
    },
  ]
}
"LosAngeles_IE8": {
  "status": "OFFLINE"
},
"LosAngeles_IE9": {
  "elapsed": 0,
  "status": "OK",
  "testers": [
    {
      "id": "WEBPAGETEST-2-64.183.41.10",
      "pc": "WEBPAGETEST-2",
      "ec2": "",
      "ip": "64.183.41.10",
      "version": "2.19.0.386",
      "freedisk": "32.051",
      "ie": "9.11.9600.18499",
      "winver": "",
      "isWinServer": "",
      "isWin64": "",
      "dns": "192.168.1.1",
      "GPU": null,
      "offline": null,
      "screenwidth": "",
      "screenheight": "",
      "rebooted": false,
      "elapsed": 0,
      "last": 323,
      "busy": 0
    }
  ]
},

*/

type jsonTesterPC struct {
	ID   string `json:"id"`
	Name string `json:"pc"`
	IP   string `json:"ip"`

	ScreenWidth  string `json:"screenwidth"`  // Screen Size
	ScreenHeight string `json:"screenheight"` // Screen Size

	EC2            string `json:"ec2"` // EC2 Instance
	DNS            string `json:"dns"` // DNS Server(s)
	AgentVersion   string `json:"version"`
	IEVersion      string `json:"ie"`       // IE Version
	WindowsVersion string `json:"winver"`   // Windows Version
	FreeDisk       string `json:"freedisk"` // Free Disk (GB)
	IsWinServer    string `json:"isWinServer"`
	IsWin64        string `json:"isWin64"`
	Offline        string `json:"offline"`
	Rebooted       bool   `json:"rebooted"`
	GPU            string `json:"GPU"`
	CPU            int    `json:"cpu"` // CPU Utilization

	Errors  int `json:"errors"` // Error Rate
	Elapsed int `json:"elapsed"`
	Last    int `json:"last"` // Last Work (minutes)
	Busy    int `json:"busy"` // Busy?
}

type jsonTester struct {
	Status  string         `json:"status"`  // "Ok" / "OFFLINE"
	Elapsed int            `json:"elapsed"` // 0
	Testers []jsonTesterPC `json:"testers"`
}

type jsonTesters struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`

	Data map[string]jsonTester `json:"data"`
}

type Tester struct {
	ID           string
	Name         string
	AgentVersion string

	// Status
	ErrorRate  int
	Elapsed    int
	LastWork   int
	IsRebooted bool
	IsOffline  bool
	IsBusy     bool

	// Network
	EC2 string
	IP  string
	DNS string

	// Screen
	ScreenWidth  int64
	ScreenHeight int64

	// Windows
	IEVersion      string
	WindowsVersion string
	IsWinServer    bool
	IsWin64        bool

	// Hardware
	FreeDisk float64
	GPU      bool
	CPU      int
}

type Testers map[string][]Tester

// GetTesters will retrieve all available agents and their status
func (w *WebPageTest) GetTesters() (*Testers, error) {
	body, err := w.query("/getTesters.php", url.Values{"f": []string{"json"}})
	if err != nil {
		return nil, err
	}

	var testers jsonTesters
	if err = json.Unmarshal(body, &testers); err != nil {
		return nil, err
	}

	if testers.StatusCode != 200 {
		return nil, fmt.Errorf("Status != 200: %v", testers.StatusText)
	}

	result := make(Testers, 0)
	for location, data := range testers.Data {
		if data.Status != "OK" {
			continue
		}
		if _, ok := result[location]; !ok {
			result[location] = make([]Tester, 0)
		}

		for _, tester := range data.Testers {
			freedisk, _ := strconv.ParseFloat(tester.FreeDisk, 32)
			width, _ := strconv.ParseInt(tester.ScreenWidth, 10, 32)
			height, _ := strconv.ParseInt(tester.ScreenHeight, 10, 32)

			result[location] = append(result[location], Tester{
				ID:   tester.ID,
				Name: tester.Name,

				AgentVersion: tester.AgentVersion,

				// Status
				ErrorRate:  tester.Errors,
				Elapsed:    tester.Elapsed,
				LastWork:   tester.Last,
				IsRebooted: tester.Rebooted,
				IsOffline:  tester.Offline == "1",
				IsBusy:     tester.Busy == 1,

				// Network
				EC2: tester.EC2,
				IP:  tester.IP,
				DNS: tester.DNS,

				// Screen
				ScreenWidth:  width,
				ScreenHeight: height,

				// Windows
				IEVersion:      tester.IEVersion,
				WindowsVersion: tester.WindowsVersion,
				IsWinServer:    tester.IsWinServer == "1",
				IsWin64:        tester.IsWin64 == "1",

				// Hardware
				FreeDisk: freedisk,
				GPU:      tester.GPU == "1",
				CPU:      tester.CPU,
			})
		}
	}

	return &result, nil
}
