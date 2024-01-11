package harhar

import (
    "encoding/json"
    "io"
    "os"
    "testing"
)

func TestConvertRequestIntoHttpRequest(t *testing.T) {

    harData, _ := os.ReadFile("api.quobix.com.har")
    var harFile HAR
    err := json.Unmarshal(harData, &harFile)

    if err != nil {
        t.Error("Error: " + err.Error())
        t.Fail()
    }

    request, err := ConvertRequestIntoHttpRequest(harFile.Log.Entries[0].Request)

    b, _ := io.ReadAll(request.Body)

    if request.Method != "POST" {
        t.Error("method does not match", request.Method)
        t.Fail()
    }
    if request.URL.String() != "https://api.quobix.com/report" {
        t.Error("url does not match", request.URL.String())
        t.Fail()
    }
    if len(b) != 46549 {
        t.Error("length of body does not match", len(b))
        t.Fail()
    }
}
