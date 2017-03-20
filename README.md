# go-webpagetest
Simple, opionated wrapper for WebPagetest API and CLI (mostly as a test for wrapper itself)

# Usage

Get it with:

    go get -d -v github.com/olegfedoseev/go-webpagetest

And than in you code:

    import  "github.com/olegfedoseev/go-webpagetest"

    wpt, err := webpagetest.NewClient("https://webpagetest.org")
    if err != nil {
      log.Fatalf("Failed to create client: %v", err)
    }

    result, err := wpt.RunTest(webpagetest.TestSettings{
      URL:      "https://google.com",
      Location: "Frankfurt_Ruxit:Chrome",
      Runs:     3,
      ScreenWidth:  1280,
      ScreenHeight: 720,
    })
    if err != nil {
      log.Fatalf("Error: %v", err)
    }
    fmt.Printf("Result: %#v", result)

Or you can look at source code of CLI at cmd/main.go
