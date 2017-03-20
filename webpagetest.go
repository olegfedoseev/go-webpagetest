package webpagetest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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

// CancelTest will try to cancel test by it's ID
// With a test ID (and if required, API key) you can cancel a test if it has not started running.
func (w *WebPageTest) CancelTest(testID string) error {
	// http://www.webpagetest.org/cancelTest.php?test=<testId>&k=<API key>
	body, err := w.query("/cancelTest.php", url.Values{"test": []string{testID}})
	if err != nil {
		return err
	}

	// <h3>Sorry, the test could not be cancelled.  It may have already started or been cancelled</h3><form>...
	// <h3 align="center">Test cancelled!</h3><form><i
	if bytes.Contains(body, []byte("Sorry, the test could not be cancelled.")) {
		// Trim left <h3> and split by < to get beginning of message
		return fmt.Errorf("%s", string(bytes.SplitN(bytes.TrimLeft(body, "<h3>"), []byte("<"), 2)[0]))
	}
	if bytes.Contains(body, []byte("Test cancelled!")) {
		return nil
	}

	return fmt.Errorf("Unknown error: %s", string(body))
}

func (w *WebPageTest) query(api string, params url.Values) ([]byte, error) {
	// http://www.webpagetest.org/cancelTest.php?test=<testId>&k=<API key>
	queryUrl := w.Host + api + "?" + params.Encode()
	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to GET \"%s\": %v", queryUrl, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status is no OK: %v [%v]", resp.StatusCode, string(body))
	}

	return body, nil
}

/*
{
  "statusCode": 200,
  "statusText": "Ok",
  "data": {
    "testId": "161128_R3_2",
    "ownerKey": "c9d1754ea6388229093c69adac3740e0339fa100",
    "jsonUrl": "http://webpagetest.app.s/jsonResult.php?test=161128_R3_2",
    "xmlUrl": "http://webpagetest.app.s/xmlResult.php?test=161128_R3_2",
    "userUrl": "http://webpagetest.app.s/results.php?test=161128_R3_2",
    "summaryCSV": "http://webpagetest.app.s/csv.php?test=161128_R3_2",
    "detailCSV": "http://webpagetest.app.s/csv.php?test=161128_R3_2&amp;requests=1"
  }
}
*/

func (w *WebPageTest) RunTest(settings TestSettings) (string, error) {
	resp, err := http.PostForm(w.Host+"/runtest.php", settings.GetFormParams())
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		StatusCode int    `json:"statusCode"`
		StatusText string `json:"statusText"`
		Data       struct {
			TestID  string `json:"testId"`
			UserURL string `json:"userUrl"`
		} `json:"data"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.StatusCode > 200 {
		return "", fmt.Errorf("StatusCode > 200: %v: %v", result.StatusCode, result.StatusText)
	}

	fmt.Printf("Result URL for %v: %v\n", settings.URL, result.Data.UserURL)
	return result.Data.TestID, nil
}

// StatusCallback is helper type for function to be called while waiting for test to complete
type StatusCallback func(testID, status string, duration int)

// RunTestAndWait will start new WebPageTest test run with given TestSettings and will wait for it
// to complete. While it wait, it will poll status updates from server and will call StatusCallback with it
func (w *WebPageTest) RunTestAndWait(settings TestSettings, callback StatusCallback) (*ResultData, error) {
	testID, err := w.RunTest(settings)
	if err != nil {
		return nil, err
	}

	for {
		result, err := w.GetTestStatus(testID)
		if err != nil {
			return nil, err
		}
		// Call callback
		if callback != nil {
			go callback(testID, result.StatusText, result.Elapsed)
		}
		if result.StatusCode < 200 {
			time.Sleep(10 * time.Second)
		}
		if result.StatusCode >= 200 {
			break
		}
	}

	testResult, err := w.GetTestResult(testID)
	if err != nil {
		return nil, err
	}
	return testResult, nil
}

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
