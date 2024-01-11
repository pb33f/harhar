package harhar

import (
    "encoding/json"
    "io"
    "os"
    "testing"
)

func TestConvertRequestIntoHttpResponse(t *testing.T) {

    harData, _ := os.ReadFile("api.quobix.com.har")
    var harFile HAR
    err := json.Unmarshal(harData, &harFile)

    if err != nil {
        t.Error("Error: " + err.Error())
        t.Fail()
    }

    response := ConvertResponseIntoHttpResponse(harFile.Log.Entries[0].Response)

    b, _ := io.ReadAll(response.Body)

    if len(b) != 884940 {
        t.Error("length of body does not match", len(b))
        t.Fail()
    }
}
