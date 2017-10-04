package webpagetest

import (
	"fmt"
	"net/url"
)

// TestSettings is structure for describing what should be done in test run
type TestSettings struct {
	// URL to be tested
	URL string `json:",omitempty"`
	// Label for the test
	Label string `json:",omitempty"`

	Where    string `json:",omitempty"`
	Browser  string `json:",omitempty"`
	Location string `json:",omitempty"`

	// Viewport Width in css pixels
	ScreenWidth int `json:",omitempty"`
	// Viewport Height in css pixels
	ScreenHeight int `json:",omitempty"`
	// Default metric to use when calculating the median run (loadTime)
	MedianMetric string `json:",omitempty"`
	// Number of test runs (1-10 on the public instance) (1)
	Runs int `json:",omitempty"`
	// Scripted test to execute ("")
	Script string `json:",omitempty"`
	// Custom Headers
	CustomHeaders string `json:",omitempty"`
	// Set to 1 to have Chrome capture the Dev Tools timeline (0)
	Timeline bool `json:",omitempty"`
	// Set to 1 to skip the Repeat View test (0)
	FirstViewOnly bool `json:",omitempty"`
	// Set to 1 to keep the test hidden from the test log (0)
	Private bool `json:",omitempty"`
	// Set to 1 to capture video (video is required for calculating Speed Index) (0)
	CaptureVideo bool `json:",omitempty"`
	// Set to 1 to save a full-resolution version of the fully loaded screen shot as a png (0)
	PNGScreenShot bool `json:",omitempty"`
	// Specify a jpeg compression level (30-100) for the screen shots and video capture (75)
	ImageQuality int `json:",omitempty"`
	// (optional) URL to ping when the test is complete (the test ID will be passed as an "id" parameter)
	Pingback string `json:",omitempty"`
	// (DOM) Element to record for sub-measurement
	DOMElement string `json:",omitempty"`
	// (Override) the number of concurrent connections IE uses (0 to not override)	0
	Connections int `json:",omitempty"`
	// (optional) Set to between 1 - 5 to have Chrome include the Javascript call stack. Must be used in conjunction with "timeline". 	 0
	TimelineStack int `json:",omitempty"`
	// (optional) Set to 1 to force the test to stop at Document Complete (onLoad)	0
	Web10 bool `json:",omitempty"`
	// (optional) space-delimited list of urls to block
	Block string `json:",omitempty"`
	// (optional) User name to use for authenticated tests (http authentication)
	Login string `json:",omitempty"`
	// (optional) Password to use for authenticated tests (http authentication)
	Password string `json:",omitempty"`
	// (optional) Type of authentication to use: 0 = Basic Auth, 1 = SNS	0
	AuthType string `json:",omitempty"`
	// (optional) e-mail address to notify with the test results
	Notify string `json:",omitempty"`
	// (optional) Download bandwidth in Kbps (used when specifying a custom connectivity profile)
	BWDown int `json:",omitempty"`
	// (optional) Upload bandwidth in Kbps (used when specifying a custom connectivity profile)
	BWUp int `json:",omitempty"`
	// (optional) First-hop Round Trip Time in ms (used when specifying a custom connectivity profile)
	Latency int `json:",omitempty"`
	// (optional) Packet loss rate - percent of packets to drop (used when specifying a custom connectivity profile)
	PacketLossRate int `json:",omitempty"`
	// (optional) (required for public instance)	API Key (if assigned) - applies only to runtest.php calls. Contact the site owner for a key if required (http://www.webpagetest.org/getkey.php for the public instance)
	APIKey string `json:",omitempty"`
	// (optional) Set to 1 to enable tcpdump capture	 0
	TCPDump bool `json:",omitempty"`
	// (optional) Set to 1 to disable optimization checks (for faster testing)	0
	NoOpt bool `json:",omitempty"`
	// (optional) Set to 1 to disable screen shot capturing	0
	NoImages bool `json:",omitempty"`
	// (optional) Set to 1 to disable saving of the http headers (as well as browser status messages and CPU utilization)	0
	NoHeaders bool `json:",omitempty"`
	// (optional) Set to 1 to disable javascript (IE, Chrome, Firefox)
	NoScript bool `json:",omitempty"`
	// (optional) Set to 1 to clear the OS certificate caches (causes IE to do OCSP/CRL checks during SSL negotiation if the certificates are not already cached). Added in 2.11	 0
	ClearCerts bool `json:",omitempty"`
	// (optional) Set to 1 to have Chrome emulate a mobile browser (screen resolution, UA string, fixed viewport).  Added in 2.11	 0
	Mobile bool `json:",omitempty"`
	// (optional) Set to 1 to preserve the original browser User Agent string (don't append PTST to it)
	KeepUA bool `json:",omitempty"`
	// (optional) Custom User Agent String to use
	UAString string `json:",omitempty"`
	// (optional) Device Pixel Ratio to use when emulating mobile
	DPR int `json:",omitempty"`
	// (optional) Set to 1 when capturing video to only store the video from the median run.	 0
	MedianRunVideo bool `json:",omitempty"`
	// (optional)  Custom command-line options (Chrome only)
	CmdLine string `json:",omitempty"`
	// (optional) Set to 1 to save the content of the first response (base page) instead of all of the text responses (bodies=1)
	HTMLBody bool `json:",omitempty"`
	// (optional)  Custom metrics to collect at the end of a test
	CustomMetrics string `json:",omitempty"`
	// (optional) Specify a specific tester that the test should run on (must match the PC name in /getTesters.php).  If the tester is not available the job will never run.
	Tester string `json:",omitempty"`
	// (optional) Specify a string that will be used to hash the test to a specific test agent.  The tester will be picked by index among the available testers.  If the number of testers changes then the tests will be distributed to different machines but if the counts remain consistent then the same string will always run the tests on the same test machine.  This can be useful for controlling variability when comparing a given URL over time or different parameters against each other (using the URL as the hash string).
	Affinity string `json:",omitempty"`
	// (optional) Set to 1 to Ignore SSL Certificate Errors e.g. Name mismatch, Self-signed certificates, etc.	 0
	IgnoreSSL bool `json:",omitempty"`
	// (optional)  Device name from mobile_devices.ini to use for mobile emulation (only when mobile=1 is specified to enable emulation and only for Chrome)
	MobileDevice string `json:",omitempty"`
	// (optional)  String to append to the user agent string. This is in addition to the default PTST/ver string. If "keepua" is also specified it will still append. Allows for substitution with some test parameters:
	// %TESTID% - Replaces with the test ID for the current test
	// %RUN% - Replaces with the current run number
	// %CACHED% - Replaces with 1 for repeat view tests and 0 for initial view
	// %VERSION% - Replaces with the current wptdriver version number
	AppendUA string `json:",omitempty"`
}

// GetFormParams returns settings that was set ready to be passed to POST
func (s TestSettings) GetFormParams() url.Values {
	values := url.Values{
		"f":            {"json"},
		"url":          {s.URL},
		"runs":         {fmt.Sprintf("%d", s.Runs)},
		"label":        {s.Label},
		"where":        {s.Where},
		"browser":      {s.Browser},
		"location":     {s.Location},
		"width":        {fmt.Sprintf("%d", s.ScreenWidth)},
		"height":       {fmt.Sprintf("%d", s.ScreenHeight)},
		"medianMetric": {s.MedianMetric},
		"script":       {s.Script},
		"pingback":     {s.Pingback},
		"domelement":   {s.DOMElement},
		"block":        {s.Block},
		"login":        {s.Login},
		"password":     {s.Password},
		"authType":     {s.AuthType},
		"notify":       {s.Notify},
		"k":            {s.APIKey},
		"uastring":     {s.UAString},
		"cmdline":      {s.CmdLine},
		"custom":       {s.CustomMetrics},
		"tester":       {s.Tester},
		"affinity":     {s.Affinity},
		"mobileDevice": {s.MobileDevice},
		"appendua":     {s.AppendUA},
	}

	if s.CustomHeaders != "" {
		values.Add("customHeaders", s.CustomHeaders)
	}
	if s.BWDown > 0 {
		values.Add("bwDown", fmt.Sprintf("%d", s.BWDown))
	}
	if s.BWUp > 0 {
		values.Add("bwUp", fmt.Sprintf("%d", s.BWUp))
	}
	if s.Latency > 0 {
		values.Add("latency", fmt.Sprintf("%d", s.Latency))
	}
	if s.PacketLossRate > 0 {
		values.Add("plr", fmt.Sprintf("%d", s.PacketLossRate))
	}

	// bool
	if s.Timeline {
		values.Add("timeline", "1")
	}
	if s.FirstViewOnly {
		values.Add("fvonly", "1")
	}
	if s.Private {
		values.Add("private", "1")
	}
	if s.CaptureVideo {
		values.Add("video", "1")
	}
	if s.PNGScreenShot {
		values.Add("pngss", "1")
	}
	if s.Web10 {
		values.Add("web10", "1")
	}
	if s.TCPDump {
		values.Add("tcpdump", "1")
	}
	if s.NoOpt {
		values.Add("noopt", "1")
	}
	if s.NoImages {
		values.Add("noimages", "1")
	}
	if s.NoHeaders {
		values.Add("noheaders", "1")
	}
	if s.NoScript {
		values.Add("noscript", "1")
	}
	if s.ClearCerts {
		values.Add("clearcerts", "1")
	}
	if s.Mobile {
		values.Add("mobile", "1")
	}
	if s.KeepUA {
		values.Add("keepua", "1")
	}
	if s.MedianRunVideo {
		values.Add("mv", "1")
	}
	if s.HTMLBody {
		values.Add("htmlbody", "1")
	}
	if s.IgnoreSSL {
		values.Add("ignoreSSL", "1")
	}
	if s.ImageQuality > 0 {
		values.Add("iq", fmt.Sprintf("%d", s.ImageQuality))
	}
	if s.TimelineStack > 0 {
		values.Add("timelineStack", fmt.Sprintf("%d", s.TimelineStack))
	}
	if s.Connections > 0 {
		values.Add("connections", fmt.Sprintf("%d", s.Connections))
	}
	if s.DPR > 0 {
		values.Add("dpr", fmt.Sprintf("%d", s.DPR))
	}

	return values
}
