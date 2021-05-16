package guia2

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/electricbubble/gadb"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

type Driver struct {
	urlPrefix *url.URL
	sessionId string

	// ext.go

	usbDevice gadb.Device
	localPort int
}

func (d *Driver) _requestURL(elem ...string) string {
	tmp, _ := url.Parse(d.urlPrefix.String())
	tmp.Path = path.Join(append([]string{d.urlPrefix.Path}, elem...)...)
	return tmp.String()
}

func (d *Driver) executeGet(pathElem ...string) (rawResp RawResponse, err error) {
	return executeHTTP(http.MethodGet, d._requestURL(pathElem...), nil)
}

func (d *Driver) executePost(data interface{}, pathElem ...string) (rawResp RawResponse, err error) {
	var bsJSON []byte = nil
	if data != nil {
		if bsJSON, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}
	return executeHTTP(http.MethodPost, d._requestURL(pathElem...), bsJSON)
}

func (d *Driver) executeDelete(pathElem ...string) (rawResp RawResponse, err error) {
	return executeHTTP(http.MethodDelete, d._requestURL(pathElem...), nil)
}

type Capabilities map[string]interface{}

func NewEmptyCapabilities() Capabilities {
	return make(Capabilities)
}

func NewDriver(capabilities Capabilities, urlPrefix string) (driver *Driver, err error) {
	if capabilities == nil {
		capabilities = NewEmptyCapabilities()
	}
	driver = new(Driver)
	if driver.urlPrefix, err = url.Parse(urlPrefix); err != nil {
		return nil, err
	}
	if driver.sessionId, err = driver.NewSession(capabilities); err != nil {
		return nil, err
	}
	return
}

func (d *Driver) NewSession(capabilities Capabilities) (sessionID string, err error) {
	// register(postHandler, new NewSession("/wd/hub/session"))
	var rawResp RawResponse
	data := map[string]interface{}{"capabilities": capabilities}
	if rawResp, err = d.executePost(data, "/session"); err != nil {
		return "", err
	}
	var reply = new(struct{ Value struct{ SessionId string } })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return "", err
	}
	sessionID = reply.Value.SessionId
	// d.sessionIdCache[sessionID] = true
	return
}

func (d *Driver) Quit() (err error) {
	// register(deleteHandler, new DeleteSession("/wd/hub/session/:sessionId"))
	if d.sessionId == "" {
		return nil
	}
	if _, err = d.executeDelete("/session", d.sessionId); err == nil {
		d.sessionId = ""
	}

	return err
}

func (d *Driver) ActiveSessionID() string {
	return d.sessionId
}

func (d *Driver) SessionIDs() (sessionIDs []string, err error) {
	// register(getHandler, new GetSessions("/wd/hub/sessions"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/sessions"); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value []struct{ SessionId string } })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}

	sessionIDs = make([]string, len(reply.Value))
	for i := range reply.Value {
		sessionIDs[i] = reply.Value[i].SessionId
	}
	return
}

func (d *Driver) SessionDetails() (scrollData map[string]interface{}, err error) {
	// register(getHandler, new GetSessionDetails("/wd/hub/session/:sessionId"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value map[string]interface{} })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}

	scrollData = reply.Value
	return
}

func (d *Driver) Status() (ready bool, err error) {
	// register(getHandler, new Status("/wd/hub/status"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/status"); err != nil {
		return false, err
	}
	var reply = new(struct {
		Value struct {
			// Message string
			Ready bool
		}
	})
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return false, err
	}
	ready = reply.Value.Ready
	return
}

// Screenshot grab device screenshot
func (d *Driver) Screenshot() (raw *bytes.Buffer, err error) {
	// register(getHandler, new CaptureScreenshot("/wd/hub/session/:sessionId/screenshot"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "screenshot"); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}

	var decodeStr []byte
	if decodeStr, err = base64.StdEncoding.DecodeString(reply.Value); err != nil {
		return nil, err
	}

	raw = bytes.NewBuffer(decodeStr)
	return
}

type Orientation string

const (
	OrientationLandscape Orientation = "LANDSCAPE"
	OrientationPortrait  Orientation = "PORTRAIT"
)

func (d *Driver) Orientation() (orientation Orientation, err error) {
	// register(getHandler, new GetOrientation("/wd/hub/session/:sessionId/orientation"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "orientation"); err != nil {
		return "", err
	}
	var reply = new(struct{ Value Orientation })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return "", err
	}

	orientation = reply.Value
	return
}

type Rotation struct {
	X, Y, Z int
}

func (d *Driver) Rotation() (rotation Rotation, err error) {
	// register(getHandler, new GetRotation("/wd/hub/session/:sessionId/rotation"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "rotation"); err != nil {
		return Rotation{}, err
	}
	var reply = new(struct{ Value Rotation })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return Rotation{}, err
	}

	rotation = reply.Value
	return
}

type Size struct {
	Width, Height int
}

// DeviceSize get window size of the device
func (d *Driver) DeviceSize() (deviceSize Size, err error) {
	// register(getHandler, new GetDeviceSize("/wd/hub/session/:sessionId/window/:windowHandle/size"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "window/:windowHandle/size"); err != nil {
		return Size{}, err
	}
	var reply = new(struct{ Value Size })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return Size{}, err
	}

	deviceSize = reply.Value
	return
}

// Source get page source
func (d *Driver) Source() (sXML string, err error) {
	// register(getHandler, new Source("/wd/hub/session/:sessionId/source"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "source"); err != nil {
		return "", err
	}
	var reply = new(struct{ Value string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return "", err
	}

	sXML = reply.Value
	return
}

// StatusBarHeight get status bar height of the device
func (d *Driver) StatusBarHeight() (height int, err error) {
	// register(getHandler, new GetSystemBars("/wd/hub/session/:sessionId/appium/device/system_bars"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "appium/device/system_bars"); err != nil {
		return 0, err
	}
	var reply = new(struct{ Value struct{ StatusBar int } })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return 0, err
	}

	height = reply.Value.StatusBar
	return
}

type BatteryStatus int

const (
	_                                  = iota
	BatteryStatusUnknown BatteryStatus = iota
	BatteryStatusCharging
	BatteryStatusDischarging
	BatteryStatusNotCharging
	BatteryStatusFull
)

func (bs BatteryStatus) String() string {
	switch bs {
	case BatteryStatusUnknown:
		return "unknown"
	case BatteryStatusCharging:
		return "charging"
	case BatteryStatusDischarging:
		return "discharging"
	case BatteryStatusNotCharging:
		return "not charging"
	case BatteryStatusFull:
		return "full"
	default:
		return fmt.Sprintf("unknown status code (%d)", bs)
	}
}

type BatteryInfo struct {
	// Battery level in range [0.0, 1.0],
	// where 1.0 means 100% charge.
	Level  float64
	Status BatteryStatus
}

func (d *Driver) BatteryInfo() (info BatteryInfo, err error) {
	// register(getHandler, new GetBatteryInfo("/wd/hub/session/:sessionId/appium/device/battery_info"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "appium/device/battery_info"); err != nil {
		return BatteryInfo{}, err
	}
	var reply = new(struct{ Value BatteryInfo })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return BatteryInfo{}, err
	}

	info = reply.Value
	if info.Level == -1 || info.Status == -1 {
		return info, errors.New("cannot be retrieved from the system")
	}
	return
}

func (d *Driver) GetAppiumSettings() (settings map[string]interface{}, err error) {
	// register(getHandler, new GetSettings("/wd/hub/session/:sessionId/appium/settings"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "appium/settings"); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value map[string]interface{} })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}

	settings = reply.Value
	return
}

// DeviceScaleRatio get device pixel ratio
func (d *Driver) DeviceScaleRatio() (scale float64, err error) {
	// register(getHandler, new GetDevicePixelRatio("/wd/hub/session/:sessionId/appium/device/pixel_ratio"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "appium/device/pixel_ratio"); err != nil {
		return 0, err
	}
	var reply = new(struct{ Value float64 })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return 0, err
	}

	scale = reply.Value
	return
}

type (
	DeviceInfo struct {
		// ANDROID_ID A 64-bit number (as a hex string) that is uniquely generated when the user
		// first sets up the device and should remain constant for the lifetime of the user's device. The value
		// may change if a factory reset is performed on the device.
		AndroidID string `json:"androidId"`
		// Build.MANUFACTURER value
		Manufacturer string `json:"manufacturer"`
		// Build.MODEL value
		Model string `json:"model"`
		// Build.BRAND value
		Brand string `json:"brand"`
		// Current running OS's API VERSION
		APIVersion string `json:"apiVersion"`
		// The current version string, for example "1.0" or "3.4b5"
		PlatformVersion string `json:"platformVersion"`
		// the name of the current celluar network carrier
		CarrierName string `json:"carrierName"`
		// the real size of the default display
		RealDisplaySize string `json:"realDisplaySize"`
		// The logical density of the display in Density Independent Pixel units.
		DisplayDensity int `json:"displayDensity"`
		// available networks
		Networks []networkInfo `json:"networks"`
		// current system locale
		Locale string `json:"locale"`
		// current system timezone
		// e.g. "Asia/Tokyo", "America/Caracas", "Asia/Shanghai"
		TimeZone  string `json:"timeZone"`
		Bluetooth struct {
			State string `json:"state"`
		} `json:"bluetooth"`
	}
	networkCapabilities struct {
		TransportTypes            string `json:"transportTypes"`
		NetworkCapabilities       string `json:"networkCapabilities"`
		LinkUpstreamBandwidthKbps int    `json:"linkUpstreamBandwidthKbps"`
		LinkDownBandwidthKbps     int    `json:"linkDownBandwidthKbps"`
		SignalStrength            int    `json:"signalStrength"`
		SSID                      string `json:"SSID"`
	}
	networkInfo struct {
		Type          int                 `json:"type"`
		TypeName      string              `json:"typeName"`
		Subtype       int                 `json:"subtype"`
		SubtypeName   string              `json:"subtypeName"`
		IsConnected   bool                `json:"isConnected"`
		DetailedState string              `json:"detailedState"`
		State         string              `json:"state"`
		ExtraInfo     string              `json:"extraInfo"`
		IsAvailable   bool                `json:"isAvailable"`
		IsRoaming     bool                `json:"isRoaming"`
		IsFailover    bool                `json:"isFailover"`
		Capabilities  networkCapabilities `json:"capabilities"`
	}
)

func (d *Driver) DeviceInfo() (info DeviceInfo, err error) {
	// register(getHandler, new GetDeviceInfo("/wd/hub/session/:sessionId/appium/device/info"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "appium/device/info"); err != nil {
		return DeviceInfo{}, err
	}
	var reply = new(struct{ Value DeviceInfo })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return DeviceInfo{}, err
	}

	info = reply.Value
	return
}

// AlertText get text of the on-screen dialog
func (d *Driver) AlertText() (text string, err error) {
	// register(getHandler, new GetAlertText("/wd/hub/session/:sessionId/alert/text"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "alert/text"); err != nil {
		return "", err
	}
	var reply = new(struct{ Value string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return "", err
	}

	text = reply.Value
	return
}

type (
	Point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	PointF struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
)

// Tap perform a click at arbitrary coordinates specified
func (d *Driver) Tap(x, y int) (err error) {
	return d.TapFloat(float64(x), float64(y))
}

func (d *Driver) TapFloat(x, y float64) (err error) {
	// register(postHandler, new Tap("/wd/hub/session/:sessionId/appium/tap"))
	data := map[string]interface{}{
		"x": x,
		"y": y,
	}
	_, err = d.executePost(data, "/session", d.sessionId, "appium/tap")
	return
}

func (d *Driver) TapPoint(point Point) (err error) {
	return d.Tap(point.X, point.Y)
}

func (d *Driver) TapPointF(point PointF) (err error) {
	return d.TapFloat(point.X, point.Y)
}

func (d *Driver) _swipe(startX, startY, endX, endY interface{}, steps int, elementID ...string) (err error) {
	// register(postHandler, new Swipe("/wd/hub/session/:sessionId/touch/perform"))
	data := map[string]interface{}{
		"startX": startX,
		"startY": startY,
		"endX":   endX,
		"endY":   endY,
		"steps":  steps,
	}
	if len(elementID) != 0 {
		data["elementId"] = elementID[0]
	}
	_, err = d.executePost(data, "/session", d.sessionId, "touch/perform")
	return
}

// Swipe performs a swipe from one coordinate to another using the number of steps
// to determine smoothness and speed. Each step execution is throttled to 5ms
// per step. So for a 100 steps, the swipe will take about 1/2 second to complete.
//  `steps` is the number of move steps sent to the system
func (d *Driver) Swipe(startX, startY, endX, endY int, steps ...int) (err error) {
	return d.SwipeFloat(float64(startX), float64(startY), float64(endX), float64(endY), steps...)
}

func (d *Driver) SwipeFloat(startX, startY, endX, endY float64, steps ...int) (err error) {
	if len(steps) == 0 {
		steps = []int{12}
	}
	return d._swipe(startX, startY, endX, endY, steps[0])
}

func (d *Driver) SwipePoint(startPoint, endPoint Point, steps ...int) (err error) {
	return d.Swipe(startPoint.X, startPoint.Y, endPoint.X, endPoint.Y, steps...)
}

func (d *Driver) SwipePointF(startPoint, endPoint PointF, steps ...int) (err error) {
	return d.SwipeFloat(startPoint.X, startPoint.Y, endPoint.X, endPoint.Y, steps...)
}

func (d *Driver) _drag(data map[string]interface{}) (err error) {
	// register(postHandler, new Drag("/wd/hub/session/:sessionId/touch/drag"))
	_, err = d.executePost(data, "/session", d.sessionId, "touch/drag")
	return
}

// Drag performs a swipe from one coordinate to another coordinate. You can control
// the smoothness and speed of the swipe by specifying the number of steps.
// Each step execution is throttled to 5 milliseconds per step, so for a 100
// steps, the swipe will take around 0.5 seconds to complete.
func (d *Driver) Drag(startX, startY, endX, endY int, steps ...int) (err error) {
	return d.DragFloat(float64(startX), float64(startY), float64(endX), float64(endY), steps...)
}

func (d *Driver) DragFloat(startX, startY, endX, endY float64, steps ...int) error {
	if len(steps) == 0 {
		steps = []int{12}
	}
	data := map[string]interface{}{
		"startX": startX,
		"startY": startY,
		"endX":   endX,
		"endY":   endY,
		"steps":  steps[0],
	}
	return d._drag(data)
}

func (d *Driver) DragPoint(startPoint Point, endPoint Point, steps ...int) error {
	return d.Drag(startPoint.X, startPoint.Y, endPoint.X, endPoint.Y, steps...)
}

func (d *Driver) DragPointF(startPoint PointF, endPoint PointF, steps ...int) (err error) {
	return d.DragFloat(startPoint.X, startPoint.Y, endPoint.X, endPoint.Y, steps...)
}

func (d *Driver) TouchLongClick(x, y int, duration ...float64) (err error) {
	if len(duration) == 0 {
		duration = []float64{1.0}
	}
	// register(postHandler, new TouchLongClick("/wd/hub/session/:sessionId/touch/longclick"))
	data := map[string]interface{}{
		"params": map[string]interface{}{
			"x":        x,
			"y":        y,
			"duration": int(duration[0] * 1000),
		},
	}
	_, err = d.executePost(data, "/session", d.sessionId, "touch/longclick")
	return
}

func (d *Driver) TouchLongClickPoint(point Point, duration ...float64) (err error) {
	return d.TouchLongClick(point.X, point.Y, duration...)
}

func (d *Driver) SendKeys(text string, isReplace ...bool) (err error) {
	if len(isReplace) == 0 {
		isReplace = []bool{true}
	}
	// register(postHandler, new SendKeysToElement("/wd/hub/session/:sessionId/keys"))
	// https://github.com/appium/appium-uiautomator2-server/blob/master/app/src/main/java/io/appium/uiautomator2/handler/SendKeysToElement.java#L76-L85
	data := map[string]interface{}{
		"text":    text,
		"replace": isReplace[0],
	}
	_, err = d.executePost(data, "/session", d.sessionId, "keys")
	return
}

// PressBack simulates a short press on the BACK button.
func (d *Driver) PressBack() (err error) {
	// register(postHandler, new PressBack("/wd/hub/session/:sessionId/back"))
	_, err = d.executePost(nil, "/session", d.sessionId, "back")
	return
}

// public class KeyCodeModel extends BaseModel {
//    @RequiredField
//    public Integer keycode;
//    public Integer metastate;
//    public Integer flags;
// }
func (d *Driver) LongPressKeyCode(keyCode KeyCode, metaState KeyMeta, flags ...KeyFlag) (err error) {
	if len(flags) == 0 {
		flags = []KeyFlag{KFFromSystem}
	}
	data := map[string]interface{}{
		"keycode":   keyCode,
		"metastate": metaState,
		"flags":     flags[0],
	}
	// register(postHandler, new LongPressKeyCode("/wd/hub/session/:sessionId/appium/device/long_press_keycode"))
	_, err = d.executePost(data, "/session", d.sessionId, "/appium/device/long_press_keycode")
	return
}

func (d *Driver) _pressKeyCode(keyCode KeyCode, metaState KeyMeta, flags ...KeyFlag) (err error) {
	// register(postHandler, new PressKeyCodeAsync("/wd/hub/session/:sessionId/appium/device/press_keycode"))
	data := map[string]interface{}{
		"keycode": keyCode,
	}
	if metaState != KMEmpty {
		data["metastate"] = metaState
	}
	if len(flags) != 0 {
		data["flags"] = flags[0]
	}
	_, err = d.executePost(data, "/session", d.sessionId, "appium/device/press_keycode")
	return
}

func (d *Driver) PressKeyCode(keyCode KeyCode, metaState KeyMeta, flags ...KeyFlag) (err error) {
	if len(flags) == 0 {
		flags = []KeyFlag{KFFromSystem}
	}
	return d._pressKeyCode(keyCode, metaState, KFFromSystem)
}

// PressKeyCodeAsync simulates a short press using a key code.
func (d *Driver) PressKeyCodeAsync(keyCode KeyCode, metaState ...KeyMeta) (err error) {
	if len(metaState) == 0 {
		metaState = []KeyMeta{KMEmpty}
	}
	return d._pressKeyCode(keyCode, metaState[0])
}

func (d *Driver) TouchDown(x, y int) (err error) {
	// register(postHandler, new TouchDown("/wd/hub/session/:sessionId/touch/down"))
	data := map[string]interface{}{
		"params": map[string]interface{}{
			"x": x,
			"y": y,
		},
	}
	_, err = d.executePost(data, "/session", d.sessionId, "touch/down")
	return
}

func (d *Driver) TouchDownPoint(point Point) error {
	return d.TouchDown(point.X, point.Y)
}

func (d *Driver) TouchUp(x, y int) (err error) {
	// register(postHandler, new TouchUp("/wd/hub/session/:sessionId/touch/up"))
	data := map[string]interface{}{
		"params": map[string]interface{}{
			"x": x,
			"y": y,
		},
	}
	_, err = d.executePost(data, "/session", d.sessionId, "touch/up")
	return
}

func (d *Driver) TouchUpPoint(point Point) error {
	return d.TouchUp(point.X, point.Y)
}

func (d *Driver) TouchMove(x, y int) (err error) {
	// register(postHandler, new TouchMove("/wd/hub/session/:sessionId/touch/move"))
	data := map[string]interface{}{
		"params": map[string]interface{}{
			"x": x,
			"y": y,
		},
	}
	_, err = d.executePost(data, "/session", d.sessionId, "touch/move")
	return
}

func (d *Driver) TouchMovePoint(point Point) error {
	return d.TouchMove(point.X, point.Y)
}

// OpenNotification opens the notification shade.
func (d *Driver) OpenNotification() (err error) {
	// register(postHandler, new OpenNotification("/wd/hub/session/:sessionId/appium/device/open_notifications"))
	_, err = d.executePost(nil, "/session", d.sessionId, "appium/device/open_notifications")
	return
}

func (d *Driver) _flick(data map[string]interface{}) (err error) {
	// register(postHandler, new Flick("/wd/hub/session/:sessionId/touch/flick"))
	_, err = d.executePost(data, "/session", d.sessionId, "touch/flick")
	return
}

func (d *Driver) Flick(xSpeed, ySpeed int) (err error) {
	data := map[string]interface{}{
		"xspeed": xSpeed,
		"yspeed": ySpeed,
	}
	if xSpeed == 0 && ySpeed == 0 {
		return errors.New("both 'xSpeed' and 'ySpeed' cannot be zero")
	}

	return d._flick(data)
}

func (d *Driver) _scrollTo(method, selector string, maxSwipes int, elementID ...string) (err error) {
	// register(postHandler, new ScrollTo("/wd/hub/session/:sessionId/touch/scroll"))
	params := map[string]interface{}{
		"strategy": method,
		"selector": selector,
	}
	if maxSwipes > 0 {
		params["maxSwipes"] = maxSwipes
	}
	data := map[string]interface{}{"params": params}
	if len(elementID) != 0 {
		data["origin"] = map[string]string{
			legacyWebElementIdentifier: elementID[0],
			webElementIdentifier:       elementID[0],
		}
	}
	_, err = d.executePost(data, "/session", d.sessionId, "touch/scroll")
	return
}

func (d *Driver) ScrollTo(by BySelector, maxSwipes ...int) (err error) {
	if len(maxSwipes) == 0 {
		maxSwipes = []int{0}
	}
	method, selector := by.getMethodAndSelector()
	return d._scrollTo(method, selector, maxSwipes[0])
}

type touchGesture struct {
	Touch PointF  `json:"touch"`
	Time  float64 `json:"time"`
}

type TouchAction []touchGesture

func NewTouchAction(cap ...int) *TouchAction {
	if len(cap) == 0 || cap[0] <= 0 {
		cap = []int{8}
	}
	tmp := make(TouchAction, 0, cap[0])
	return &tmp
}

func (ta *TouchAction) Add(x, y int, startTime ...float64) *TouchAction {
	return ta.AddFloat(float64(x), float64(y), startTime...)
}

func (ta *TouchAction) AddFloat(x, y float64, startTime ...float64) *TouchAction {
	if len(startTime) == 0 {
		var tmp float64 = 0
		if len(*ta) != 0 {
			g := (*ta)[len(*ta)-1]
			tmp = g.Time + 0.05
		}
		startTime = []float64{tmp}
	}
	*ta = append(*ta, touchGesture{Touch: PointF{x, y}, Time: startTime[0]})
	return ta
}

func (ta *TouchAction) AddPoint(point Point, startTime ...float64) *TouchAction {
	return ta.AddFloat(float64(point.X), float64(point.Y), startTime...)
}

func (ta *TouchAction) AddPointF(point PointF, startTime ...float64) *TouchAction {
	return ta.AddFloat(point.X, point.Y, startTime...)
}

func (d *Driver) MultiPointerGesture(gesture1 *TouchAction, gesture2 *TouchAction, tas ...*TouchAction) (err error) {
	// Must provide coordinates for at least 2 pointers
	actions := make([]*TouchAction, 0)
	actions = append(actions, gesture1, gesture2)
	if len(tas) != 0 {
		actions = append(actions, tas...)
	}
	data := map[string]interface{}{
		"actions": actions,
	}
	// register(postHandler, new MultiPointerGesture("/wd/hub/session/:sessionId/touch/multi/perform"))
	_, err = d.executePost(data, "/session", d.sessionId, "/touch/multi/perform")
	return
}

type w3cGesture map[string]interface{}

func _newW3CGesture() w3cGesture {
	return make(w3cGesture)
}

func (g w3cGesture) _set(key string, value interface{}) w3cGesture {
	g[key] = value
	return g
}

func (g w3cGesture) pause(duration float64) w3cGesture {
	return g._set("type", "pause").
		_set("duration", duration)
}

func (g w3cGesture) keyDown(value string) w3cGesture {
	return g._set("type", "keyDown").
		_set("value", value)
}

func (g w3cGesture) keyUp(value string) w3cGesture {
	return g._set("type", "keyUp").
		_set("value", value)
}

func (g w3cGesture) pointerDown(button int) w3cGesture {
	return g._set("type", "pointerDown")._set("button", button)
}

func (g w3cGesture) pointerUp(button int) w3cGesture {
	return g._set("type", "pointerUp")._set("button", button)
}

func (g w3cGesture) pointerMove(x, y float64, origin string, duration float64, pressureAndSize ...float64) w3cGesture {
	switch len(pressureAndSize) {
	case 1:
		g._set("pressure", pressureAndSize[0])
	case 2:
		g._set("pressure", pressureAndSize[0])
		g._set("size", pressureAndSize[1])
	}
	return g._set("type", "pointerMove").
		_set("duration", duration).
		_set("origin", origin).
		_set("x", x).
		_set("y", y)

}

func (g w3cGesture) size(size ...float64) w3cGesture {
	if len(size) == 0 {
		size = []float64{1.0}
	}
	return g._set("size", size[0])
}

func (g w3cGesture) pressure(pressure ...float64) w3cGesture {
	if len(pressure) == 0 {
		pressure = []float64{1.0}
	}
	return g._set("pressure", pressure[0])
}

type W3CGestures []w3cGesture

func NewW3CGestures(cap ...int) *W3CGestures {
	if len(cap) == 0 || cap[0] <= 0 {
		cap = []int{8}
	}
	tmp := make(W3CGestures, 0, cap[0])
	return &tmp
}

func (g *W3CGestures) Pause(duration ...float64) *W3CGestures {
	if len(duration) == 0 || duration[0] < 0 {
		duration = []float64{0.5}
	}
	*g = append(*g, _newW3CGesture().pause(duration[0]*1000))
	return g
}

func (g *W3CGestures) KeyDown(value string) *W3CGestures {
	*g = append(*g, _newW3CGesture().keyDown(value))
	return g
}

func (g *W3CGestures) KeyUp(value string) *W3CGestures {
	*g = append(*g, _newW3CGesture().keyUp(value))
	return g
}

func (g *W3CGestures) SendKeys(text string) *W3CGestures {
	ss := strings.Split(text, "")
	for i := range ss {
		g.KeyDown(ss[i])
		g.KeyUp(ss[i])
	}
	return g
}

type W3CMouseButtonType int

const (
	MBTLeft   W3CMouseButtonType = 0
	MBTMiddle W3CMouseButtonType = 1
	MBTRight  W3CMouseButtonType = 2
)

func (g *W3CGestures) PointerDown(button ...W3CMouseButtonType) *W3CGestures {
	if len(button) == 0 {
		button = []W3CMouseButtonType{MBTLeft}
	}
	*g = append(*g, _newW3CGesture().pointerDown(int(button[0])))
	return g
}

func (g *W3CGestures) PointerUp(button ...W3CMouseButtonType) *W3CGestures {
	if len(button) == 0 {
		button = []W3CMouseButtonType{MBTLeft}
	}
	*g = append(*g, _newW3CGesture().pointerUp(int(button[0])))
	return g
}

type W3CPointerMoveType string

const (
	PMTViewport W3CPointerMoveType = "viewport"
	PMTPointer  W3CPointerMoveType = "pointer"
)

func (g *W3CGestures) PointerMove(x, y float64, origin interface{}, duration float64, pressure, size float64) *W3CGestures {
	val := ""
	switch v := origin.(type) {
	case string:
		val = v
	case W3CPointerMoveType:
		val = string(v)
	case *Element:
		val = v.id
	default:
		val = string(PMTViewport)
	}
	*g = append(*g, _newW3CGesture().pointerMove(x, y, val, duration, pressure, size))
	return g
}

func (g *W3CGestures) PointerMoveTo(x, y float64, duration ...float64) *W3CGestures {
	if len(duration) == 0 || duration[0] < 0 {
		duration = []float64{0.5}
	}
	*g = append(*g, _newW3CGesture().pointerMove(x, y, string(PMTViewport), duration[0]*1000))
	return g
}

func (g *W3CGestures) PointerMoveRelative(x, y float64, duration ...float64) *W3CGestures {
	if len(duration) == 0 || duration[0] < 0 {
		duration = []float64{0.5}
	}
	*g = append(*g, _newW3CGesture().pointerMove(x, y, string(PMTPointer), duration[0]*1000))
	return g
}

func (g *W3CGestures) PointerMouseOver(x, y float64, element *Element, duration ...float64) *W3CGestures {
	if len(duration) == 0 || duration[0] < 0 {
		duration = []float64{0.5}
	}
	*g = append(*g, _newW3CGesture().pointerMove(x, y, element.id, duration[0]*1000))
	return g
}

type W3CAction map[string]interface{}

type W3CActionType string

const (
	_         W3CActionType = "none"
	ATKey     W3CActionType = "key"
	ATPointer W3CActionType = "pointer"
)

type W3CPointerType string

const (
	PTMouse W3CPointerType = "mouse"
	PTPen   W3CPointerType = "pen"
	PTTouch W3CPointerType = "touch"
)

func NewW3CAction(actionType W3CActionType, gestures *W3CGestures, pointerType ...W3CPointerType) W3CAction {
	w3cAction := make(W3CAction)
	w3cAction["type"] = actionType
	w3cAction["actions"] = gestures
	if actionType != ATPointer {
		return w3cAction
	}

	if len(pointerType) == 0 {
		pointerType = []W3CPointerType{PTTouch}
	}
	type W3CItemParameters struct {
		PointerType W3CPointerType `json:"pointerType"`
	}
	w3cAction["parameters"] = W3CItemParameters{PointerType: pointerType[0]}
	return w3cAction
}

func (d *Driver) PerformW3CActions(action W3CAction, acts ...W3CAction) (err error) {
	var actionId uint64 = 1
	acts = append([]W3CAction{action}, acts...)
	for i := range acts {
		item := acts[i]
		item["id"] = strconv.FormatUint(actionId, 10)
		actionId++
		acts[i] = item
	}
	data := map[string]interface{}{
		"actions": acts,
	}
	// register(postHandler, new W3CActions("/wd/hub/session/:sessionId/actions"))
	_, err = d.executePost(data, "/session", d.sessionId, "/actions")
	return
}

type ClipDataType string

const ClipDataTypePlaintext ClipDataType = "PLAINTEXT"

func (d *Driver) GetClipboard(contentType ...ClipDataType) (content string, err error) {
	if len(contentType) == 0 {
		contentType = []ClipDataType{ClipDataTypePlaintext}
	}
	// register(postHandler, new GetClipboard("/wd/hub/session/:sessionId/appium/device/get_clipboard"))
	data := map[string]interface{}{
		"contentType": contentType[0],
	}
	var rawResp RawResponse
	if rawResp, err = d.executePost(data, "/session", d.sessionId, "appium/device/get_clipboard"); err != nil {
		return "", err
	}
	var reply = new(struct{ Value string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return "", err
	}

	content = reply.Value
	if data, err := base64.StdEncoding.DecodeString(content); err != nil {
		return content, err
	} else {
		content = string(data)
	}
	return
}

func (d *Driver) SetClipboard(contentType ClipDataType, content string, label ...string) (err error) {
	lbl := content
	if len(label) != 0 {
		lbl = label[0]
	}
	const defaultLabelLen = 10
	if len(lbl) > defaultLabelLen {
		lbl = lbl[:defaultLabelLen]
	}

	data := map[string]interface{}{
		"contentType": contentType,
		"label":       lbl,
		"content":     base64.StdEncoding.EncodeToString([]byte(content)),
	}
	// register(postHandler, new SetClipboard("/wd/hub/session/:sessionId/appium/device/set_clipboard"))
	_, err = d.executePost(data, "/session", d.sessionId, "appium/device/set_clipboard")
	return
}

func (d *Driver) AlertAccept(buttonLabel ...string) (err error) {
	data := map[string]interface{}{
		"buttonLabel": nil,
	}
	if len(buttonLabel) != 0 {
		data["buttonLabel"] = buttonLabel[0]
	}
	// register(postHandler, new AcceptAlert("/wd/hub/session/:sessionId/alert/accept"))
	_, err = d.executePost(data, "/session", d.sessionId, "alert/accept")
	return
}

func (d *Driver) AlertDismiss(buttonLabel ...string) (err error) {
	data := map[string]interface{}{
		"buttonLabel": nil,
	}
	if len(buttonLabel) != 0 {
		data["buttonLabel"] = buttonLabel[0]
	}
	// register(postHandler, new DismissAlert("/wd/hub/session/:sessionId/alert/dismiss"))
	_, err = d.executePost(data, "/session", d.sessionId, "alert/dismiss")
	return
}

func (d *Driver) SetAppiumSettings(settings map[string]interface{}) (err error) {
	data := map[string]interface{}{
		"settings": settings,
	}
	// register(postHandler, new UpdateSettings("/wd/hub/session/:sessionId/appium/settings"))
	_, err = d.executePost(data, "/session", d.sessionId, "appium/settings")
	return
}

func (d *Driver) SetOrientation(orientation Orientation) (err error) {
	data := map[string]interface{}{
		"orientation": orientation,
	}
	// register(postHandler, new SetOrientation("/wd/hub/session/:sessionId/orientation"))
	_, err = d.executePost(data, "/session", d.sessionId, "orientation")
	return
}

// SetRotation
//  `x` and `y` are ignored. We only care about `z`
//  0/90/180/270
func (d *Driver) SetRotation(rotation Rotation) (err error) {
	data := map[string]interface{}{
		"z": rotation.Z,
	}
	// register(postHandler, new SetRotation("/wd/hub/session/:sessionId/rotation"))
	_, err = d.executePost(data, "/session", d.sessionId, "rotation")
	return
}

type NetworkType int

const (
	NetworkTypeWifi NetworkType = 2

	// NetworkTypeNone NetworkType = iota
	// NetworkTypeAirplane
	// NetworkTypeWifi
	// _
	// NetworkTypeData
	// _
	// NetworkTypeAll
)

// NetworkConnection always turn on
func (d *Driver) NetworkConnection(networkType NetworkType) (err error) {
	// register(postHandler, new NetworkConnection("/wd/hub/session/:sessionId/network_connection"))
	data := map[string]interface{}{
		"type": networkType,
	}
	_, err = d.executePost(data, "/session", d.sessionId, "network_connection")
	return
}

func (d *Driver) _findElements(method, selector string, elementID ...string) (elements []*Element, err error) {
	// register(postHandler, new FindElements("/wd/hub/session/:sessionId/elements"))
	data := map[string]interface{}{
		"strategy": method,
		"selector": selector,
	}
	if len(elementID) != 0 {
		data["context"] = elementID[0]
	}
	var rawResp RawResponse
	if rawResp, err = d.executePost(data, "/session", d.sessionId, "/elements"); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value []map[string]string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}
	if len(reply.Value) == 0 {
		return nil, fmt.Errorf("no such element: unable to find an element using '%s', value '%s'", method, selector)
	}
	elements = make([]*Element, len(reply.Value))
	for i, elem := range reply.Value {
		var id string
		if id = elementIDFromValue(elem); id == "" {
			return nil, fmt.Errorf("invalid element returned: %+v", reply)
		}
		elements[i] = &Element{parent: d, id: id}
	}
	return
}

func (d *Driver) _findElement(method, selector string, elementID ...string) (elem *Element, err error) {
	// register(postHandler, new FindElement("/wd/hub/session/:sessionId/element"))
	data := map[string]interface{}{
		"strategy": method,
		"selector": selector,
	}
	if len(elementID) != 0 {
		data["context"] = elementID[0]
	}
	var rawResp RawResponse
	if rawResp, err = d.executePost(data, "/session", d.sessionId, "/element"); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value map[string]string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}
	if len(reply.Value) == 0 {
		return nil, fmt.Errorf("no such element: unable to find an element using '%s', value '%s'", method, selector)
	}
	var id string
	if id = elementIDFromValue(reply.Value); id == "" {
		return nil, fmt.Errorf("invalid element returned: %+v", reply)
	}
	elem = &Element{parent: d, id: id}
	return
}

func (d *Driver) FindElements(by BySelector) (elements []*Element, err error) {
	return d._findElements(by.getMethodAndSelector())
}

func (d *Driver) FindElement(by BySelector) (elem *Element, err error) {
	return d._findElement(by.getMethodAndSelector())
}

func (d *Driver) ActiveElement() (elem *Element, err error) {
	// register(getHandler, new ActiveElement("/wd/hub/session/:sessionId/element/active"))
	var rawResp RawResponse
	if rawResp, err = d.executeGet("/session", d.sessionId, "/element/active"); err != nil {
		return nil, err
	}
	var reply = new(struct{ Value map[string]string })
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, err
	}
	if len(reply.Value) == 0 {
		return nil, errors.New("no such element")
	}
	var id string
	if id = elementIDFromValue(reply.Value); id == "" {
		return nil, fmt.Errorf("invalid element returned: %+v", reply)
	}
	elem = &Element{parent: d, id: id}
	return
}

type Condition func(d *Driver) (bool, error)

func (d *Driver) _waitWithTimeoutAndInterval(condition Condition, timeout, interval time.Duration) (err error) {
	startTime := time.Now()
	for {
		done, err := condition(d)
		if err != nil {
			return err
		}
		if done {
			return nil
		}

		if elapsed := time.Since(startTime); elapsed > timeout {
			return fmt.Errorf("timeout after %v", elapsed)
		}
		time.Sleep(interval)
	}
}

// WaitWithTimeoutAndInterval waits for the condition to evaluate to true.
func (d *Driver) WaitWithTimeoutAndInterval(condition Condition, timeout, interval float64) (err error) {
	dTimeout := time.Millisecond * time.Duration(timeout*1000)
	dInterval := time.Millisecond * time.Duration(interval*1000)
	return d._waitWithTimeoutAndInterval(condition, dTimeout, dInterval)
}

// WaitWithTimeout works like WaitWithTimeoutAndInterval, but with default polling interval.
func (d *Driver) WaitWithTimeout(condition Condition, timeout float64) error {
	dTimeout := time.Millisecond * time.Duration(timeout*1000)
	return d._waitWithTimeoutAndInterval(condition, dTimeout, DefaultWaitInterval)
}

// Wait works like WaitWithTimeoutAndInterval, but using the default timeout and polling interval.
func (d *Driver) Wait(condition Condition) error {
	return d._waitWithTimeoutAndInterval(condition, DefaultWaitTimeout, DefaultWaitInterval)
}
