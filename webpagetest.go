package webpagetest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WebPageTest struct {
	Host string
}

func NewClient(host string) (*WebPageTest, error) {
	validURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	return &WebPageTest{
		Host: validURL.String(),
	}, nil
}

type TestStatus struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       struct {
		StatusCode              int    `json:"statusCode"`
		StatusText              string `json:"statusText"`
		BehindCount             int    `json:"behindCount"`
		ID                      string `json:"id"`
		TestID                  string `json:"testId"`
		Runs                    int    `json:"runs"`
		Remote                  bool   `json:"remote"` // Relay Test
		Location                string `json:"location"`
		FirstViewOnly           int    `json:"fvonly"`
		StartTime               string `json:"startTime"`
		CompleteTime            string `json:"completeTime"`
		Elapsed                 int    `json:"elapsed"`
		ElapsedUpdate           int    `json:"elapsedUpdate"`
		TestsExpected           int    `json:"testsExpected"`
		TestsCompleted          int    `json:"testsCompleted"`
		FirstViewRunsCompleted  int    `json:"fvRunsCompleted"`
		RepeatViewRunsCompleted int    `json:"rvRunsCompleted"`
		TestInfo                struct {
			URL            string `json:"url"`
			Runs           int    `json:"runs"`
			FirstViewOnly  int    `json:"fvonly"`
			Web10          int    `json:"web10"`     // Stop Test at Document Complete
			IgnoreSSL      int    `json:"ignoreSSL"` // Ignore SSL Certificate Errors
			Video          string `json:"video"`
			Label          string `json:"label"`
			Priority       int    `json:"priority"`
			Location       string `json:"location"`
			Browser        string `json:"browser"`
			Connectivity   string `json:"connectivity"`
			BandwidthIn    int    `json:"bwIn"`
			BandwidthOut   int    `json:"bwOut"`
			Latency        int    `json:"latency"`
			PacketLossRate string `json:"plr"`
			Tcpdump        int    `json:"tcpdump"`  // Capture network packet trace (tcpdump)
			Timeline       int    `json:"timeline"` // Capture Dev Tools Timeline
			Trace          int    `json:"trace"`    // Capture Chrome Trace (about://tracing)
			Bodies         int    `json:"bodies"`
			NetLog         int    `json:"netlog"`    // Capture Network Log
			Standards      int    `json:"standards"` // Disable Compatibility View (IE Only)
			NoScript       int    `json:"noscript"`  // Disable Javascript
			Pngss          int    `json:"pngss"`
			Iq             int    `json:"iq"`
			KeepUA         int    `json:"keepua"` // Preserve original User Agent string
			Mobile         int    `json:"mobile"`
			Scripted       int    `json:"scripted"`
		} `json:"testInfo"`
	} `json:"data"`
}

// GetTestStatus will return status of test run by given testID
// StatusCode 200 indicates test is completed. 1XX means the test is still
// in progress. And 4XX indicates some error.
func (w *WebPageTest) GetTestStatus(testID string) (*TestStatus, error) {
	statusUrl := w.Host + "/testStatus.php?test=" + testID
	resp, err := http.Get(statusUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to GET \"%s\": %v", statusUrl, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status is not OK: %v [%v]", resp.StatusCode, string(body))
	}
	fmt.Printf("body: %v\n", string(body))

	var result TestStatus
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.StatusCode != 200 {
		return nil, fmt.Errorf("Status != 200: %v", result.StatusText)
	}

	return &result, nil
}

type Location struct {
	Label         string         `json:"Label"`
	LabelShort    string         `json:"labelShort"`
	Location      string         `json:"location"`
	Browser       string         `json:"Browser"`
	RelayServer   string         `json:"relayServer"`
	RelayLocation string         `json:"relayLocation"`
	Default       bool           `json:"default"`
	PendingTests  map[string]int `json:"PendingTests"`
}

type Locations struct {
	StatusCode int                 `json:"statusCode"`
	StatusText string              `json:"statusText"`
	Data       map[string]Location `json:"data`
}

// GetLocations will retrieve all available locations from server
func (w *WebPageTest) GetLocations() (*Locations, error) {
	resultsUrl := w.Host + "/getLocations.php?f=json"
	resp, err := http.Get(resultsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to GET \"%s\": %v", resultsUrl, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status is no OK: %v [%v]", resp.StatusCode, string(body))
	}

	var result Locations
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.StatusCode != 200 {
		return nil, fmt.Errorf("Status != 200: %v", result.StatusText)
	}

	return &result, nil
}

type TestResult struct {
}

func (w *WebPageTest) GetTestResults(testID string) (*TestResult, error) {
	resultsUrl := w.Host + "/jsonResult.php?test=" + testID
	resp, err := http.Get(resultsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to GET \"%s\": %v", resultsUrl, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status is no OK: %v [%v]", resp.StatusCode, string(body))
	}
	fmt.Printf("body: %v\n", string(body))

	// var result struct {
	// 	StatusCode int         `json:"statusCode"`
	// 	StatusText string      `json:"statusText"`
	// 	Data       TestResults `json:"data"`
	// }

	// if err = json.Unmarshal(body, &result); err != nil {
	// 	return nil, err
	// }

	// if result.StatusCode != 200 {
	// 	return nil, fmt.Errorf("Status != 200: %v", result.StatusText)
	// }

	return &TestResult{}, nil
}

// getTestResults(id, options, callback)
// getTesters(options, callback)
// runTest(url_or_script, options, callback)
// cancelTest(id, options, callback)

// getHARData(id, options, callback)
// getPageSpeedData(id, options, callback)
// getUtilizationData(id, options, callback)
// getRequestData(id, options, callback)
// getTimelineData(id, options, callback)
// getNetLogData(id, options, callback)
// getChromeTraceData(id, options, callback)
// getConsoleLogData(id, options, callback)
// getTestInfo(id, options, callback)
// getHistory(days, options, callback)
// getGoogleCsiData(id, options, callback)
// getResponseBody(id, options, callback)
// getWaterfallImage(id, options, callback)
// getScreenshotImage(id, options, callback)
// createVideo(tests, options, callback)
// getEmbedVideoPlayer(id, options, callback)
// scriptToString(script)
