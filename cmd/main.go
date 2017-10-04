package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/olegfedoseev/go-webpagetest"
)

// https://community.akamai.com/community/web-performance/blog/2016/08/25/using-navigation-timing-apis-to-understand-your-webpage
var (
	wpt *webpagetest.WebPageTest
)

func main() {
	usage := `WebPagetest CLI

Usage:
  webpagetest locations [--server=<url>]
  webpagetest testers [--server=<url>]
  webpagetest status <testID> [--server=<url>]
  webpagetest cancel <testID> [--server=<url>]
  webpagetest results <testID> [--server=<url>] [--step=<stepIdx>]
  webpagetest -h | --help
  webpagetest --version

Options:
  -h --help         Show this screen.
  --version         Show version.
  --server=<url>    URL of private instance of WebPagetest Server
  --step=<stepIdx>  Index of test step to use as source of metrics (1-based)`

	arguments, _ := docopt.Parse(usage, nil, true, "WebPagetest CLI 1.0", false)

	server := "https://webpagetest.org"
	if arguments["--server"] != nil && arguments["--server"].(string) != "" {
		server = arguments["--server"].(string)
	}

	fmt.Printf("Will use server at %s\n", server)
	var err error
	wpt, err = webpagetest.NewClient(server)
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		os.Exit(2)
	}
	if arguments["locations"].(bool) {
		getLocations()
	}

	if arguments["testers"].(bool) {
		getTesters()
	}

	if arguments["status"].(bool) {
		getStatus(arguments["<testID>"].(string))
	}

	if arguments["results"].(bool) {
		var step int64
		if arguments["--step"] != nil && arguments["--step"].(string) != "" {
			step, _ = strconv.ParseInt(arguments["--step"].(string), 10, 32)
			fmt.Printf("Will use step #%d\n", step)
		}

		getResults(arguments["<testID>"].(string), step)
	}

	// TODO: figure out how to specify all test params

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

	// result, err := wpt.RunTestAndWait(webpagetest.TestSettings{
	// 	URL:          "http://ngs.ru",
	// 	Location:     "74RU_wpt:Chrome",
	// 	Runs:         3,
	// 	ScreenWidth:  1280,
	// 	ScreenHeight: 720,
	// })
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	os.Exit(2)
	// }
}

// Get Locations
func getLocations() {
	result, err := wpt.GetLocations()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	for name, locations := range *result {
		fmt.Println(name)
		for _, location := range locations {
			fmt.Printf("  %s [%s] Status: %s\n    Browsers\n", location.Label, location.Location, location.Status)
			for _, browser := range location.Browsers {
				fmt.Printf("    - %s\n", browser)
			}
		}
	}
}

// Get Testers
func getTesters() {
	result, err := wpt.GetTesters()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	for name, testers := range *result {
		fmt.Printf("Testers for '%s':\n", name)
		for idx, tester := range testers {
			fmt.Printf("  [%d] %s v%s %s\n", idx, tester.Name, tester.AgentVersion, tester.IP)
		}
	}
}

// Test Status
func getStatus(testID string) {
	result, err := wpt.GetTestStatus(testID)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
	fmt.Printf("Result: %#v", result)
}

// Cancel Test
func cancelTest(testID string) {
	err := wpt.CancelTest(testID)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}
}

func getResults(testID string, testStep int64) {
	// Test Result
	// 161124_CC_3 - google.com
	// 161122_K9_A - novosibirsk.n1.ru
	// 161118_62_db87f3f04fe6b52b8cf4481fcf32cc0a
	// 161126_19_12569a3f0de7a2fec98475b5d8bb0d37 (google.cz)
	result, err := wpt.GetTestResult(testID)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(2)
	}

	fmt.Printf("ID: %v\n", result.ID)
	fmt.Printf("URL: %v\n", result.URL)
	fmt.Printf("Summary: %v\n", result.Summary)
	fmt.Printf("Location: %v\n", result.Location)
	fmt.Printf("Completed: %v\n", time.Unix(int64(result.Completed), 0))

	fmt.Printf("Runs: %v\n", len(result.Runs))

	runs := make([]string, 0)
	for run := range result.Runs {
		runs = append(runs, run)
	}
	sort.Strings(runs)

	for _, run := range runs {
		views := result.Runs[run]
		fmt.Printf("Run %s\n", run)

		for idx, step := range views.FirstView.Steps {
			if step.Error != "" {
				fmt.Printf("[%d] FirstView Result [%v] %v \n", idx+1, step.Result, step.Error)
				continue
			}
			fmt.Println(stepAsTableRow(&step, idx == 0, "First View "))
		}
		if !result.FirstViewOnly {
			for idx, step := range views.RepeatView.Steps {
				if step.Error != "" {
					fmt.Printf("[%d] RepeatView Result [%v] %v \n", idx+1, step.Result, step.Error)
					continue
				}
				fmt.Println(stepAsTableRow(&step, idx == 0, "Repeat View "))
			}
		}
	}
	if testStep == 0 {
		testStep = 1
	}

	medianRun, err := result.GetMedianRun(int(testStep-1), "loadtime")
	if err != nil {
		fmt.Printf("GetMedianRun failed: %v\n", err)
	}
	fmt.Printf("\nMedian run\n")
	fmt.Println(stepAsTableRow(&medianRun.FirstView.Steps[testStep-1], true,
		fmt.Sprintf("Run: #%d/%d ", medianRun.FirstView.Run, medianRun.RepeatView.Run)))
	fmt.Println(stepAsTableRow(&medianRun.RepeatView.Steps[testStep-1], false, ""))
}

func stepAsTableRow(ts *webpagetest.TestStep, header bool, headerTitle string) string {
	var result string
	if header {
		result += fmt.Sprintf("[%15s| %9s | %10s | %12s | %11s | %17s | %12s]\n",
			headerTitle, "Load Time", "First Byte", "Start Render", "Speed Index", "Document Complete", "Fully Loaded")
	}

	result += fmt.Sprintf("[%15s| %9v | %10v | %12v | %11v | %17v | %12v]",
		fmt.Sprintf("Step %d ", ts.Step),
		time.Duration(ts.LoadTime)*time.Millisecond,
		time.Duration(ts.TTFB)*time.Millisecond,
		time.Duration(ts.StartRender)*time.Millisecond,
		ts.SpeedIndex,
		time.Duration(ts.DocTime)*time.Millisecond,
		time.Duration(ts.FullyLoaded)*time.Millisecond)

	return result
}
