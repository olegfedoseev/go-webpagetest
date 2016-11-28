package webpagetest

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// https://sites.google.com/a/webpagetest.org/docs/advanced-features/raw-test-results

type VideoFrame struct {
	Time  int    `json:"time"`
	Image string `json:"image"`

	VisuallyComplete int `json:"VisuallyComplete"`
}

type Domain struct {
	Bytes       int    `json:"bytes"`
	Requests    int    `json:"requests"`
	CDNProvider string `json:"cdn_provider"`
	Connections int    `json:"connections"`
}

type Breakdown struct {
	Color []int `json:"color"`
	Bytes int   `json:"bytes"`

	Requests int `json:"requests"`
}

type Headers struct {
	Request  []string `json:"request"`
	Response []string `json:"response"`
}

type jsonRequest struct {
	IP           string `json:"ip_addr"`      // "173.194.122.199"
	Method       string `json:"method"`       // "GET"
	Host         string `json:"host"`         // "google.com"
	URL          string `json:"url"`          // "/"
	FullURL      string `json:"full_url"`     // "http://google.com/"
	ResponseCode string `json:"responseCode"` // "302",

	Protocol  string `json:"protocol"`          // "HTTP/2"
	RequestID int    `json:"request_id,string"` // "9"
	Index     int    `json:"index"`             // 0
	Number    int    `json:"number"`            // 1

	Type     int    `json:"type,string"`   // "3"
	Socket   int    `json:"socket,string"` // "22"
	Priority string `json:"priority"`      // "VeryHigh",

	// Network
	BytesOut         int `json:"bytesOut,string"`          // "397"
	BytesIn          int `json:"bytesIn,string"`           // "467"
	ServerCount      int `json:"server_count,string"`      // "11"
	ServerRTT        int `json:"server_rtt,string"`        // "26"
	ClientPort       int `json:"client_port,string"`       // "55276"
	IsSecure         int `json:"is_secure,string"`         // "0"
	CertificateBytes int `json:"certificate_bytes,string"` // "0", "3769",

	// Cache
	Expires         string `json:"expires"`           // "Tue, 14 Nov 2017 22:46:51 GMT", "-1"
	CacheControl    string `json:"cacheControl"`      // "private"
	CacheTime       int    `json:"cache_time,string"` // "0"
	ContentType     string `json:"contentType"`       // "text/html"
	ContentEncoding string `json:"contentEncoding"`   // "gzip"
	ObjectSize      int    `json:"objectSize,string"` // "256"
	CDNProvider     string `json:"cdn_provider"`      // "Google",

	// Timings
	DNSStart json.Number `json:"dns_start"` // "0"
	DNSEnd   json.Number `json:"dns_end"`   // "50"
	DNS      json.Number `json:"dns_ms"`    // "-1",

	ConnectStart json.Number `json:"connect_start"` // "50"
	ConnectEnd   json.Number `json:"connect_end"`   // "76"
	Connect      json.Number `json:"connect_ms"`    // 26,

	SSLStart json.Number `json:"ssl_start"` // "0"
	SSLEnd   json.Number `json:"ssl_end"`   // "0"
	SSL      json.Number `json:"ssl_ms"`    // "-1",

	LoadStart json.Number `json:"load_start,string"` // "76"
	LoadEnd   json.Number `json:"load_end"`          // 119
	Load      json.Number `json:"load_ms,string"`    // "43",

	TTFBStart json.Number `json:"ttfb_start"`     // "76"
	TTFBEnd   json.Number `json:"ttfb_end"`       // 119
	TTFB      json.Number `json:"ttfb_ms,string"` // "43",

	DownloadStart json.Number `json:"download_start"` // 119
	DownloadEnd   json.Number `json:"download_end"`   // 119
	Download      json.Number `json:"download_ms"`    // 0,

	AllStart json.Number `json:"all_start"` // "50"
	AllEnd   json.Number `json:"all_end"`   // 119
	All      json.Number `json:"all_ms"`    // 69,

	// Optimizations
	ScoreCache           json.Number `json:"score_cache"`            // "0"
	ScoreCDN             json.Number `json:"score_cdn"`              // "-1"
	ScoreGZip            json.Number `json:"score_gzip"`             // "-1"
	ScoreCookies         json.Number `json:"score_cookies"`          // "-1"
	ScoreKeepAlive       json.Number `json:"score_keep-alive"`       // "-1"
	ScoreMinify          json.Number `json:"score_minify"`           // "-1"
	ScoreCombine         json.Number `json:"score_combine"`          // "-1"
	ScoreCompress        json.Number `json:"score_compress"`         // "-1"
	ScoreETags           json.Number `json:"score_etags"`            // "-1"
	ScoreProgressiveJpeg json.Number `json:"score_progressive_jpeg"` // -1
	GZipTotal            json.Number `json:"gzip_total"`             // "0"
	GZipSave             json.Number `json:"gzip_save"`              // "0"
	MinifyTotal          json.Number `json:"minify_total"`           // "0"
	MinifySave           json.Number `json:"minify_save"`            // "0"
	ImageTotal           json.Number `json:"image_total"`            // "0"
	ImageSave            json.Number `json:"image_save"`             // "0"
	JpegScanCount        json.Number `json:"jpeg_scan_count"`        // "0",

	// HTTP/2
	HTTP2StreamDependency int `json:"http2_stream_dependency,string"` // "5"
	HTTP2StreamExclusive  int `json:"http2_stream_exclusive,string"`  // "1"
	HTTP2StreamID         int `json:"http2_stream_id,string"`         // "1"
	HTTP2StreamWeight     int `json:"http2_stream_weight,string"`     // "256"
	WasPushed             int `json:"was_pushed,string"`              // "0",

	// Initiator info
	Initiator         string `json:"initiator"`               // "https://www.google.cz/?gfe_rd=cr&ei=JDc5WJ2sDqSE8QfT-5SgBw&gws_rd=ssl"
	InitiatorColumn   int    `json:"initiator_column,string"` // "104"
	InitiatorDetail   string `json:"initiator_detail"`        // "{\"lineNumber\":50,\"type\":\"parser\",\"url\":\"https://www.google.cz/?gfe_rd=cr&ei=JDc5WJ2sDqSE8QfT-5SgBw&gws_rd=ssl\"}"
	InitiatorFunction string `json:"initiator_function"`      // "Xm"
	InitiatorLine     int    `json:"initiator_line,string"`   // "50"
	InitiatorType     string `json:"initiator_type"`          // "other",

	Headers Headers `json:"headers"`
}

type TestView struct {
	URL    string `json:"URL"`
	Result int    `json:"result"` // 99999
	Date   int    `json:"date"`   // 1479973600
	Run    int    `json:"run"`    // 1,

	Title          string `json:"title"`           // "Google"
	Tester         string `json:"tester"`          // "74SYNCMASTER-172.28.116.4"
	BrowserName    string `json:"browser_name"`    // "Google Chrome"
	BrowserVersion string `json:"browser_version"` // "54.0.2840.99",

	// Estimated RTT to Server (ms)
	ServerRTT int `json:"server_rtt"`
	// Time to First Byte (ms)
	// The First Byte time (often abbreviated as TTFB) is measured as the time from the start of
	// the initial navigation until the first byte of the base page is received by the browser (after following redirects).
	TTFB int `json:"TTFB"`
	// Time to DOM Loading - From Navigation Timing
	DOMLoading int `json:"domLoading"`
	// First Paint (ms)
	FirstPaint int `json:"firstPaint"`
	// Time to Title (ms)
	TitleTime int `json:"titleTime"`
	// Time to DOM Interactive - From Navigation Timing
	DOMInteractive int `json:"domInteractive"`
	// DOM Content Loaded - From Navigation Timing
	DOMContentLoadedEventStart int `json:"domContentLoadedEventStart"`
	// Browser-reported Load Time (Navigation Timing onload)
	LoadEventStart int `json:"loadEventStart"`
	// Load Time (onload, ms)
	// The Load Time is measured as the time from the start of the initial navigation until the beginning of the window load event (onload).
	DocTime  int `json:"docTime"`
	LoadTime int `json:"loadTime"`
	// Time to Start Render (ms)
	// The Start Render time is measured as the time from the start of the initial
	// navigation until the first non-white content is painted to the browser display.
	Render int `json:"render"`
	// Time to Visually Complete (ms)
	VisualComplete int `json:"visualComplete"`
	// Fully Loaded (ms)
	// The Fully Loaded time is measured as the time from the start of the initial navigation until
	// there was 2 seconds of no network activity after Document Complete.  This will usually
	// include any activity that is triggered by javascript after the main page loads.
	FullyLoaded int `json:"fullyLoaded"`
	// Last Visual Change (ms)
	LastVisualChange int `json:"lastVisualChange"`

	SpeedIndex int `json:"SpeedIndex"`

	// Number of DOM Elements
	// The DOM Elements metric is the count of the DOM elements on the tested page as measured at the end of the test.
	DOMElements int `json:"domElements"`

	LoadEventEnd             int `json:"loadEventEnd"`
	DOMTime                  int `json:"domTime"`
	AboveTheFoldTime         int `json:"aft"`
	DOMContentLoadedEventEnd int `json:"domContentLoadedEventEnd"` // 455,

	// CPU Busy Time (ms)
	DocCPUms         float32 `json:"docCPUms"`         // 951.606
	FullyLoadedCPUms float32 `json:"fullyLoadedCPUms"` // 1294.808,

	DocCPUpct         int `json:"docCPUpct"`         // 39
	FullyLoadedCPUpct int `json:"fullyLoadedCPUpct"` // 19,

	Cached int `json:"cached"` // 0,

	BytesIn          int `json:"bytesIn"`
	BytesOut         int `json:"bytesOut"`
	BytesInDoc       int `json:"bytesInDoc"`
	BytesOutDoc      int `json:"bytesOutDoc"`
	EffectiveBps     int `json:"effectiveBps"`      // 433693
	EffectiveBpsDoc  int `json:"effectiveBpsDoc"`   // 466135
	CertificateBytes int `json:"certificate_bytes"` // 17499,

	Connections int `json:"connections"`

	Requests []jsonRequest `json:"requests"`

	RequestsFull int `json:"requestsFull"`
	RequestsDoc  int `json:"requestsDoc"`

	Responses200   int `json:"responses_200"`
	Responses404   int `json:"responses_404"`
	ResponsesOther int `json:"responses_other"`

	OptimizationChecked  int `json:"optimization_checked"`   // 1
	ScoreCache           int `json:"score_cache"`            // 0
	ScoreCDN             int `json:"score_cdn"`              // -1
	ScoreGZip            int `json:"score_gzip"`             // -1
	ScoreCookies         int `json:"score_cookies"`          // -1
	ScoreKeepAlive       int `json:"score_keep-alive"`       // -1
	ScoreMinify          int `json:"score_minify"`           // -1
	ScoreCombine         int `json:"score_combine"`          // 100
	ScoreCompress        int `json:"score_compress"`         // -1
	ScoreETags           int `json:"score_etags"`            // -1
	ScoreProgressiveJpeg int `json:"score_progressive_jpeg"` // -1,

	GZipTotal   int `json:"gzip_total"`   // 0
	GZipSavings int `json:"gzip_savings"` // 0,

	MinifyTotal   int `json:"minify_total"`   // 0
	MinifySavings int `json:"minify_savings"` // 0,

	ImageTotal   int `json:"image_total"`   // 0
	ImageSavings int `json:"image_savings"` // 0,

	PageSpeedVersion string `json:"pageSpeedVersion"` // "1.9",

	ServerCount int `json:"server_count"` // 16,

	AdultSite int `json:"adult_site"` // 0,

	FixedViewport int `json:"fixed_viewport"` // 0
	IsResponsive  int `json:"isResponsive"`   // -1,

	EventName string `json:"eventName"` // "Step 1"
	NumSteps  int    `json:"numSteps"`  // 1
	Step      int    `json:"step"`      // 1,

	BasePageCDN       string `json:"base_page_cdn"`       // "Google"
	BasePageRedirects int    `json:"base_page_redirects"` // 2
	BasePageTTFB      int    `json:"base_page_ttfb"`      // 524,

	BrowserProcessCount         int `json:"browser_process_count"`           // 8
	BrowserMainMemoryKB         int `json:"browser_main_memory_kb"`          // 69752
	BrowserWorkingSetKB         int `json:"browser_working_set_kb"`          // 136568
	BrowserOtherPrivateMemoryKB int `json:"browser_other_private_memory_kb"` // 66816,

	// "details": "http://webpagetest.app.s/details.php?test=161124_CC_3&run=1&cached=0"
	// "checklist": "http://webpagetest.app.s/performance_optimization.php?test=161124_CC_3&run=1&cached=0"
	// "breakdown": "http://webpagetest.app.s/breakdown.php?test=161124_CC_3&run=1&cached=0"
	// "domains": "http://webpagetest.app.s/domains.php?test=161124_CC_3&run=1&cached=0"
	// "screenShot": "http://webpagetest.app.s/screen_shot.php?test=161124_CC_3&run=1&cached=0"
	Pages map[string]string `json:"pages"`

	// "waterfall": "http://webpagetest.app.s/result/161124_CC_3/1_waterfall_thumb.png"
	// "checklist": "http://webpagetest.app.s/result/161124_CC_3/1_optimization_thumb.png"
	// "screenShot": "http://webpagetest.app.s/result/161124_CC_3/1_screen_thumb.png"
	Thumbnails map[string]string `json:"thumbnails"`

	// "waterfall": "http://webpagetest.app.s/results/16/11/24/CC/3/1_waterfall.png"
	// "connectionView": "http://webpagetest.app.s/results/16/11/24/CC/3/1_connection.png"
	// "checklist": "http://webpagetest.app.s/results/16/11/24/CC/3/1_optimization.png"
	// "screenShot": "http://webpagetest.app.s/getfile.php?test=161124_CC_3&file=1_screen.jpg"
	// "screenShotPng": "http://webpagetest.app.s/getfile.php?test=161124_CC_3&file=1_screen.png"
	Images map[string]string `json:"images"`

	// "headers": "http://webpagetest.app.s/results/16/11/24/CC/3/1_report.txt"
	// "pageData": "http://webpagetest.app.s/results/16/11/24/CC/3/1_IEWPG.txt"
	// "requestsData": "http://webpagetest.app.s/results/16/11/24/CC/3/1_IEWTR.txt"
	// "utilization": "http://webpagetest.app.s/results/16/11/24/CC/3/1_progress.csv"
	RawData map[string]string `json:"rawData"`

	VideoFrames []VideoFrame         `json:"videoFrames"`
	Domains     map[string]Domain    `json:"domains"`
	Breakdown   map[string]Breakdown `json:"breakdown"`
}

type TestRun struct {
	FirstView  TestView `json:"firstView"`
	RepeatView TestView `json:"repeatView"`
}

type ResultData struct {
	Connectivity

	ID       string `json:"id"`
	URL      string `json:"url"`
	Summary  string `json:"summary"`
	TestUrl  string `json:"testUrl"`
	Location string `json:"location"`
	Label    string `json:"label"`
	From     string `json:"from"`

	Mobile           int    `json:"mobile"`
	Completed        int    `json:"completed"`
	Tester           string `json:"tester"`
	TesterDNS        string `json:"testerDNS"`
	FirstViewOnly    bool   `json:"fvonly"`
	SuccessfulFVRuns int    `json:"successfulFVRuns"`
	SuccessfulRVRuns int    `json:"successfulRVRuns"`

	Runs   map[string]TestRun `json:"runs"`
	Median TestRun            `json:"median"`
	// Average AverageTestRun     `json:"average"`
	// StdDev  StdDevTestRun      `json:"standardDeviation"`
}

type jsonTestResult struct {
	StatusCode int        `json:"statusCode"`
	StatusText string     `json:"statusText"`
	Data       ResultData `json:"data"`
}

func (w *WebPageTest) GetTestResult(testID string) (*ResultData, error) {
	body, err := w.query("/jsonResult.php", url.Values{"test": []string{testID}})
	if err != nil {
		return nil, err
	}

	// fmt.Printf("body: %v\n", string(body))

	var result jsonTestResult
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.StatusCode != 200 {
		return nil, fmt.Errorf("Status != 200: %v", result.StatusText)
	}

	return &result.Data, nil
}