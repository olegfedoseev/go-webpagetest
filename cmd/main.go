package main

import (
	"fmt"
	"github.com/olegfedoseev/webpagetest"
	"os"
)

// https://community.akamai.com/community/web-performance/blog/2016/08/25/using-navigation-timing-apis-to-understand-your-webpage

func main() {
	fmt.Println("WebPageTest")

	if len(os.Args) == 0 {
		fmt.Println("You should provide URL for test!")
		os.Exit(2)
	}

	// fmt.Println("arg", os.Args[0])

	// wpt, _ := webpagetest.NewClient("http://webpagetest.org")
	wpt, _ := webpagetest.NewClient("http://webpagetest.app.s")

	// Cancel Test
	// err := wpt.CancelTest("161124_QZ_1")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }

	// Get Testers
	// result, err := wpt.GetTesters()
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	// Get Locations
	// result, err := wpt.GetLocations()
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	// Test Status
	// result, err := wpt.GetTestStatus("161126_19_12569a3f0de7a2fec98475b5d8bb0d37")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	// Test Result
	// 161124_CC_3 - google.com
	// 161122_K9_A - novosibirsk.n1.ru
	// 161118_62_db87f3f04fe6b52b8cf4481fcf32cc0a
	// 161126_19_12569a3f0de7a2fec98475b5d8bb0d37 (google.cz)
	// result, err := wpt.GetTestResult("161128_S0_1")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// // fmt.Printf("Result: %#v", result)
	// fmt.Printf("ID: %v\n", result.ID)
	// fmt.Printf("URL: %v\n", result.URL)
	// fmt.Printf("Summary: %v\n", result.Summary)
	// fmt.Printf("TestUrl: %v\n", result.TestUrl)
	// fmt.Printf("Location: %v\n", result.Location)
	// fmt.Printf("Label: %v\n", result.Label)
	// fmt.Printf("From: %v\n", result.From)

	// fmt.Printf("Connectivity: %s\n", result.Connectivity.String())

	// fmt.Printf("Mobile: %v\n", result.Mobile)
	// fmt.Printf("Completed: %v\n", result.Completed)
	// fmt.Printf("Tester: %v\n", result.Tester)
	// fmt.Printf("TesterDNS: %v\n", result.TesterDNS)
	// fmt.Printf("FirstViewOnly: %v\n", result.FirstViewOnly)
	// fmt.Printf("SuccessfulFVRuns: %v\n", result.SuccessfulFVRuns)
	// fmt.Printf("SuccessfulRVRuns: %v\n", result.SuccessfulRVRuns)

	// fmt.Printf("\nEstimated RTT to Server (ms) %v\n", result.Median.RepeatView.ServerRTT)
	// fmt.Printf("Time to First Byte (ms) %v\n", result.Median.RepeatView.TTFB)
	// fmt.Printf("DOM Loading %v\n", result.Median.RepeatView.DOMLoading)
	// fmt.Printf("First Paint (ms) %v\n", result.Median.RepeatView.FirstPaint)
	// fmt.Printf("Time to Title (ms) %v\n", result.Median.RepeatView.TitleTime)
	// fmt.Printf("DOM Interactive %v\n", result.Median.RepeatView.DOMInteractive)
	// fmt.Printf("DOM Content Loaded (Navigation Timing) %v\n", result.Median.RepeatView.DOMContentLoadedEventStart)
	// fmt.Printf("Browser-reported Load Time (Navigation Timing onload) %v\n", result.Median.RepeatView.LoadEventStart)
	// fmt.Printf("Load Time (onload, ms) %v\n", result.Median.RepeatView.LoadTime)
	// fmt.Printf("Load Time (DocTime, ms) %v\n", result.Median.RepeatView.DocTime)
	// fmt.Printf("Time to Start Render (ms) %v\n", result.Median.RepeatView.Render)
	// fmt.Printf("Time to Visually Complete (ms) %v\n", result.Median.RepeatView.VisualComplete)
	// fmt.Printf("Fully Loaded (ms) %v\n", result.Median.RepeatView.FullyLoaded)
	// fmt.Printf("Last Visual Change (ms) %v\n", result.Median.RepeatView.LastVisualChange)

	// fmt.Printf("\nSpeedIndex %v\n", result.Median.RepeatView.SpeedIndex)

	// fmt.Printf("CPU Busy Time (ms) %v\n", result.Median.RepeatView.DocCPUms)

	// fmt.Printf("\nNumber of DOM Elements %v\n", result.Median.RepeatView.DOMElements)
	// fmt.Printf("Connections %v\n", result.Median.RepeatView.Connections)
	// fmt.Printf("Requests (Fully Loaded) %v\n", len(result.Median.RepeatView.Requests))
	// fmt.Printf("Requests (onload) %v\n", result.Median.RepeatView.RequestsDoc)

	// fmt.Printf("\nBytes In (onload) %v\n", result.Median.RepeatView.BytesInDoc)
	// fmt.Printf("Bytes In (Fully Loaded) %v\n", result.Median.RepeatView.BytesIn)
	// fmt.Printf("\nBrowser Version %v\n", result.Median.RepeatView.BrowserVersion)

	// fmt.Printf("\nRepeatView.LoadEventEnd %v\n", result.Median.RepeatView.LoadEventEnd)
	// fmt.Printf("RepeatView.DOMTime %v\n", result.Median.RepeatView.DOMTime)
	// fmt.Printf("RepeatView.AboveTheFoldTime %v\n", result.Median.RepeatView.AboveTheFoldTime)

	// fmt.Printf("\nLoad Event %v\n", result.Median.RepeatView.LoadEventEnd-result.Median.RepeatView.LoadEventStart)

	// fmt.Printf("DOMContentLoaded Event %v\n",
	// 	result.Median.RepeatView.DOMContentLoadedEventEnd-result.Median.RepeatView.DOMContentLoadedEventStart)

	// 'basePageSSLTime' => 'Base Page SSL Time (ms)',

	// Runs   map[string]TestRun `json:"runs"`
	// Median TestRun            `json:"median"`

	// result, err := wpt.RunTest(webpagetest.TestSettings{
	// 	URL:      "http://ngs.ru",
	// 	Location: "74RU_wpt:Chrome",
	// 	Runs:     3,
	// 	// ScreenWidth:  1280,
	// 	// ScreenHeight: 720,
	// })
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
	// fmt.Printf("Result: %#v", result)

	result, err := wpt.RunTestAndWait(webpagetest.TestSettings{
		URL:          "http://ngs.ru",
		Location:     "74RU_wpt:Chrome",
		Runs:         3,
		ScreenWidth:  1280,
		ScreenHeight: 720,
	})
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	fmt.Printf("URL: %v\n", result.URL)
	fmt.Printf("Summary: %v\n", result.Summary)
	fmt.Printf("Location: %v\n", result.Location)
	fmt.Printf("Completed: %v\n", result.Completed)
	fmt.Printf("Tester: %v\n", result.Tester)
	fmt.Printf("SuccessfulFVRuns: %v\n", result.SuccessfulFVRuns)

	fmt.Printf("\nEstimated RTT to Server (ms) %v\n", result.Median.RepeatView.ServerRTT)
	fmt.Printf("Time to First Byte (ms) %v\n", result.Median.RepeatView.TTFB)
	fmt.Printf("DOM Loading %v\n", result.Median.RepeatView.DOMLoading)
	fmt.Printf("First Paint (ms) %v\n", result.Median.RepeatView.FirstPaint)
	fmt.Printf("Time to Title (ms) %v\n", result.Median.RepeatView.TitleTime)
	fmt.Printf("DOM Interactive %v\n", result.Median.RepeatView.DOMInteractive)
	fmt.Printf("DOM Content Loaded (Navigation Timing) %v\n", result.Median.RepeatView.DOMContentLoadedEventStart)
	fmt.Printf("Browser-reported Load Time (Navigation Timing onload) %v\n", result.Median.RepeatView.LoadEventStart)
	fmt.Printf("Load Time (onload, ms) %v\n", result.Median.RepeatView.LoadTime)
	fmt.Printf("Load Time (DocTime, ms) %v\n", result.Median.RepeatView.DocTime)
	fmt.Printf("Time to Start Render (ms) %v\n", result.Median.RepeatView.Render)
	fmt.Printf("Time to Visually Complete (ms) %v\n", result.Median.RepeatView.VisualComplete)
	fmt.Printf("Fully Loaded (ms) %v\n", result.Median.RepeatView.FullyLoaded)
	fmt.Printf("Last Visual Change (ms) %v\n", result.Median.RepeatView.LastVisualChange)

	fmt.Printf("\nSpeedIndex %v\n", result.Median.RepeatView.SpeedIndex)

	fmt.Printf("CPU Busy Time (ms) %v\n", result.Median.RepeatView.DocCPUms)

	fmt.Printf("\nNumber of DOM Elements %v\n", result.Median.RepeatView.DOMElements)
	fmt.Printf("Connections %v\n", result.Median.RepeatView.Connections)
	fmt.Printf("Requests (Fully Loaded) %v\n", len(result.Median.RepeatView.Requests))
	fmt.Printf("Requests (onload) %v\n", result.Median.RepeatView.RequestsDoc)

	fmt.Printf("\nBytes In (onload) %v\n", result.Median.RepeatView.BytesInDoc)
	fmt.Printf("Bytes In (Fully Loaded) %v\n", result.Median.RepeatView.BytesIn)
	fmt.Printf("\nBrowser Version %v\n", result.Median.RepeatView.BrowserVersion)

	fmt.Printf("\nRepeatView.LoadEventEnd %v\n", result.Median.RepeatView.LoadEventEnd)
	fmt.Printf("RepeatView.DOMTime %v\n", result.Median.RepeatView.DOMTime)
	fmt.Printf("RepeatView.AboveTheFoldTime %v\n", result.Median.RepeatView.AboveTheFoldTime)

	fmt.Printf("\nLoad Event %v\n", result.Median.RepeatView.LoadEventEnd-result.Median.RepeatView.LoadEventStart)

	fmt.Printf("DOMContentLoaded Event %v\n",
		result.Median.RepeatView.DOMContentLoadedEventEnd-result.Median.RepeatView.DOMContentLoadedEventStart)
}
