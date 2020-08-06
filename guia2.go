package guia2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

// HTTPClient is the default client to use to communicate with the WebDriver server.
var HTTPClient = http.DefaultClient

// type MobileDriver struct{}
// type MobileElement struct{}
// type MobileSelector struct{}

type RawResponse []byte

var uia2Header = map[string]string{
	"Content-Type": "application/json;charset=UTF-8",
	"accept":       "application/json",
}

func executeHTTP(method string, url string, rawBody []byte) (rawResp RawResponse, err error) {
	debugLog(fmt.Sprintf("--> %s %s\n%s", method, url, rawBody))

	var req *http.Request
	if req, err = http.NewRequest(method, url, bytes.NewBuffer(rawBody)); err != nil {
		return
	}
	for k, v := range uia2Header {
		req.Header.Set(k, v)
	}

	start := time.Now()
	var resp *http.Response
	if resp, err = HTTPClient.Do(req); err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	rawResp, err = ioutil.ReadAll(resp.Body)
	debugLog(fmt.Sprintf("<-- %s %s %d %s\n%s\n", method, url, resp.StatusCode, time.Now().Sub(start), rawResp))
	if err != nil {
		return nil, err
	}

	var reply = new(struct {
		Value struct {
			Err        string `json:"error"`
			Message    string `json:"message"`
			Stacktrace string `json:"stacktrace"`
		}
	})
	if err = json.Unmarshal(rawResp, reply); err != nil {
		if resp.StatusCode == http.StatusOK {
			// 如果遇到 value 直接是 字符串，则报错，但是 http 状态是 200
			// {"sessionId":"b4f2745a-be74-4cb3-8f4c-881cde817a8d","value":"YWJjZDEyMw==\n"}
			return rawResp, nil
		}
		return nil, err
	}
	if reply.Value.Err != "" {
		return nil, fmt.Errorf("%s: %s", reply.Value.Err, reply.Value.Message)
	}

	return
}

const (
	// legacyWebElementIdentifier is the string constant used in the old
	// WebDriver JSON protocol that is the key for the map that contains an
	// unique element identifier.
	legacyWebElementIdentifier = "ELEMENT"

	// webElementIdentifier is the string constant defined by the W3C
	// specification that is the key for the map that contains a unique element identifier.
	webElementIdentifier = "element-6066-11e4-a52e-4f735466cecf"
)

func elementIDFromValue(val map[string]string) string {
	for _, key := range []string{webElementIdentifier, legacyWebElementIdentifier} {
		if v, ok := val[key]; ok && v != "" {
			return v
		}
	}
	return ""
}

type BySelector struct {
	// Set the search criteria to match the given resource ResourceIdID.
	ResourceIdID string `json:"id"`
	// Set the search criteria to match the content-description property for a widget.
	ContentDescription string `json:"accessibility id"`
	XPath              string `json:"xpath"`
	// Set the search criteria to match the class property for a widget (for example, "android.widget.Button").
	ClassName   string `json:"class name"`
	UiAutomator string `json:"-android uiautomator"`
}

func (by BySelector) getMethodAndSelector() (method, selector string) {
	vBy := reflect.ValueOf(by)
	tBy := reflect.TypeOf(by)
	for i := 0; i < vBy.NumField(); i++ {
		vi := vBy.Field(i).Interface()
		// switch vi := vi.(type) {
		// case string:
		// 	selector = vi
		// }
		selector = vi.(string)
		if selector != "" && selector != "UNKNOWN" {
			method = tBy.Field(i).Tag.Get("json")
			return
		}
	}
	return
}

var debugFlag = false

// SetDebug set debug mode
func SetDebug(debug bool) {
	debugFlag = debug
}

func debugLog(msg string) {
	if !debugFlag {
		return
	}
	log.Println("[DEBUG] " + msg)
}
