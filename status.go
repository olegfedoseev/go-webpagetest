package webpagetest

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// testStatus.php

/*
{
  "statusCode": 200,
  "statusText": "Test Complete",
  "data": {
    "statusCode": 200,
    "statusText": "Test Complete",
    "id": "161118_62_db87f3f04fe6b52b8cf4481fcf32cc0a",
    "testInfo": {
      "url": "http://google.com",
      "runs": 1,
      "fvonly": 0,
      "web10": 0,
      "ignoreSSL": 0,
      "video": "on",
      "label": "",
      "priority": 0,
      "block": "",
      "location": "Dulles",
      "browser": "Chrome",
      "connectivity": "Cable",
      "bwIn": 5000,
      "bwOut": 1000,
      "latency": 28,
      "plr": "0",
      "tcpdump": 0,
      "timeline": 0,
      "trace": 0,
      "bodies": 0,
      "netlog": 0,
      "standards": 0,
      "noscript": 0,
      "pngss": 0,
      "iq": 0,
      "keepua": 0,
      "mobile": 0,
      "addCmdLine": "",
      "scripted": 0
    },
    "testId": "161118_62_db87f3f04fe6b52b8cf4481fcf32cc0a",
    "runs": 1,
    "fvonly": 0,
    "remote": false,
    "testsExpected": 1,
    "location": "Dulles",
    "startTime": "11/18/16 3:30:22",
    "elapsed": 14,
    "completeTime": "11/18/16 3:30:36",
    "testsCompleted": 1,
    "fvRunsCompleted": 1,
    "rvRunsCompleted": 1
  }
}

{
  "statusCode": 100,
  "statusText": "Completed 1 of 3 tests",
  "data": {
    "statusCode": 100,
    "statusText": "Completed 1 of 3 tests",
    "id": "161126_19_12569a3f0de7a2fec98475b5d8bb0d37",
    "testInfo": {
      "url": "http://google.com",
      "runs": 3,
      "fvonly": 0,
      "web10": 0,
      "ignoreSSL": 0,
      "video": "on",
      "label": "test run",
      "priority": 0,
      "block": "",
      "location": "Prague",
      "browser": "Chrome",
      "connectivity": "Cable",
      "bwIn": 5000,
      "bwOut": 1000,
      "latency": 28,
      "plr": "0",
      "tcpdump": 0,
      "timeline": 0,
      "trace": 0,
      "bodies": 0,
      "netlog": 0,
      "standards": 0,
      "noscript": 0,
      "pngss": 0,
      "iq": 0,
      "keepua": 0,
      "mobile": 0,
      "addCmdLine": "",
      "scripted": 0
    },
    "testId": "161126_19_12569a3f0de7a2fec98475b5d8bb0d37",
    "runs": 3,
    "fvonly": 0,
    "remote": false,
    "testsExpected": 3,
    "location": "Prague",
    "startTime": "11/26/16 7:17:53",
    "elapsed": 24,
    "fvRunsCompleted": 2,
    "rvRunsCompleted": 1,
    "testsCompleted": 1
  }
}
*/

type TestInfo struct {
	URL           string `json:"url"`
	Runs          int    `json:"runs"`
	FirstViewOnly int    `json:"fvonly"`
	Web10         int    `json:"web10"`     // Stop Test at Document Complete
	IgnoreSSL     int    `json:"ignoreSSL"` // Ignore SSL Certificate Errors
	Video         string `json:"video"`
	Label         string `json:"label"`
	Priority      int    `json:"priority"`
	Location      string `json:"location"`
	Browser       string `json:"browser"`

	Connectivity   string `json:"connectivity"`
	BandwidthIn    int    `json:"bwIn"`
	BandwidthOut   int    `json:"bwOut"`
	Latency        int    `json:"latency"`
	PacketLossRate int    `json:"plr,string"`

	Tcpdump      int `json:"tcpdump"`  // Capture network packet trace (tcpdump)
	Timeline     int `json:"timeline"` // Capture Dev Tools Timeline
	Trace        int `json:"trace"`    // Capture Chrome Trace (about://tracing)
	Bodies       int `json:"bodies"`
	NetLog       int `json:"netlog"`    // Capture Network Log
	Standards    int `json:"standards"` // Disable Compatibility View (IE Only)
	NoScript     int `json:"noscript"`  // Disable Javascript
	Pngss        int `json:"pngss"`
	ImageQuality int `json:"iq"`
	KeepUA       int `json:"keepua"` // Preserve original User Agent string
	Mobile       int `json:"mobile"`
	Scripted     int `json:"scripted"`
}

type TestStatus struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`

	ID           string `json:"id"`
	TestID       string `json:"testId"`
	Location     string `json:"location"`
	StartTime    string `json:"startTime"`
	CompleteTime string `json:"completeTime"`

	Runs        int `json:"runs"`
	BehindCount int `json:"behindCount"`

	Remote         bool `json:"remote"` // Relay Test
	FirstViewOnly  int  `json:"fvonly"`
	Elapsed        int  `json:"elapsed"`
	ElapsedUpdate  int  `json:"elapsedUpdate"`
	TestsExpected  int  `json:"testsExpected"`
	TestsCompleted int  `json:"testsCompleted"`

	FirstViewRunsCompleted  int `json:"fvRunsCompleted"`
	RepeatViewRunsCompleted int `json:"rvRunsCompleted"`

	TestInfo TestInfo `json:"testInfo"`
}

type jsonTestStatus struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`

	Data TestStatus `json:"data"`
}

// GetTestStatus will return status of test run by given testID
// StatusCode 200 indicates test is completed. 1XX means the test is still
// in progress. And 4XX indicates some error.
func (w *WebPageTest) GetTestStatus(testID string) (*TestStatus, error) {
	body, err := w.query("/testStatus.php", url.Values{"test": []string{testID}})
	if err != nil {
		return nil, err
	}

	var result jsonTestStatus
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.StatusCode > 200 {
		return nil, fmt.Errorf("%s", result.StatusText)
	}

	return &result.Data, nil
}
