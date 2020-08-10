package guia2

import (
	"io/ioutil"
	"testing"
	"time"
)

var uiaServerURL = "http://localhost:6790/wd/hub"

func TestDriver_NewSession(t *testing.T) {
	SetDebug(true)

	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	firstMatchEntry := make(map[string]interface{})
	firstMatchEntry["package"] = "com.android.settings"
	firstMatchEntry["activity"] = "com.android.settings/.Settings"
	caps := Capabilities{
		"firstMatch":  []interface{}{firstMatchEntry},
		"alwaysMatch": struct{}{},
	}
	sessionID, err := driver.NewSession(caps)
	if err != nil {
		t.Fatal(err)
	}
	if len(sessionID) == 0 {
		t.Fatal("should not be empty")
	}
}

func TestNewDriver(t *testing.T) {
	SetDebug(true)

	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(driver.sessionId)
}

func TestDriver_Quit(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	if err = driver.Quit(); err != nil {
		t.Fatal(err)
	}
}

func TestDriver_Status(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	_, err = driver.Status()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_SessionIDs(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	sessions, err := driver.SessionIDs()
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) == 0 {
		t.Fatal("should have at least one")
	}
	t.Log(len(sessions), sessions)
}

func TestDriver_SessionDetails(t *testing.T) {
	// firstMatchEntry := make(map[string]interface{})
	// firstMatchEntry["package"] = "com.android.settings"
	// firstMatchEntry["activity"] = "com.android.settings/.Settings"
	// caps = Capabilities{
	// 	"firstMatch":  []interface{}{firstMatchEntry},
	// 	"alwaysMatch": struct{}{},
	// }
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	scrollData, err := driver.SessionDetails()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(scrollData)
}

func TestDriver_Screenshot(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	screenshot, err := driver.Screenshot()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ioutil.WriteFile("/Users/hero/Desktop/s1.png", screenshot.Bytes(), 0600))
}

func TestDriver_Orientation(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	orientation, err := driver.Orientation()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(orientation)
}

func TestDriver_Rotation(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	rotation, err := driver.Rotation()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("x = %d\ty = %d\tz = %d", rotation.X, rotation.Y, rotation.Z)
}

func TestDriver_DeviceSize(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	deviceSize, err := driver.DeviceSize()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("width = %d\theight = %d", deviceSize.Width, deviceSize.Height)
}

func TestDriver_Source(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	source, err := driver.Source()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(source)
}

func TestDriver_StatusBarHeight(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	statusBarHeight, err := driver.StatusBarHeight()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(statusBarHeight)
}

func TestDriver_BatteryInfo(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	batteryInfo, err := driver.BatteryInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(batteryInfo)
}

func TestDriver_GetAppiumSettings(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	appiumSettings, err := driver.GetAppiumSettings()
	if err != nil {
		t.Fatal(err)
	}

	for k := range appiumSettings {
		t.Logf("key: %s\tvalue: %v", k, appiumSettings[k])
	}
	// t.Log(appiumSettings)
}

func TestDriver_DeviceScaleRatio(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	scaleRatio, err := driver.DeviceScaleRatio()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(scaleRatio)
}

func TestDriver_DeviceInfo(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	devInfo, err := driver.DeviceInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("api version: %s", devInfo.APIVersion)
	t.Logf("platform version: %s", devInfo.PlatformVersion)
	t.Logf("bluetooth state: %s", devInfo.Bluetooth.State)
}

func TestDriver_AlertText(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	alertText, err := driver.AlertText()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(alertText)
}

func TestDriver_Tap(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.Tap(150, 340)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)

	err = driver.TapFloat(60.5, 125.5)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)

	err = driver.TapPoint(Point{X: 150, Y: 340})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)

	err = driver.TapPointF(PointF{X: 60.5, Y: 125.5})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_Swipe(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.Swipe(400, 1000, 400, 500, 10)
	if err != nil {
		t.Fatal(err)
	}

	err = driver.SwipeFloat(400, 555.5, 400, 1255.5)
	if err != nil {
		t.Fatal(err)
	}

	startPoint := Point{400, 1000}
	endPoint := Point{400, 500}
	err = driver.SwipePoint(startPoint, endPoint)
	if err != nil {
		t.Fatal(err)
	}

	startPointF := PointF{400, 555.5}
	endPointF := PointF{400, 1255.5}
	err = driver.SwipePointF(startPointF, endPointF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_Drag(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.Drag(400, 260, 400, 500, 10)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 200)

	err = driver.DragFloat(400, 501.5, 400, 261.5, 10)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 200)

	startPoint := Point{400, 260}
	endPoint := Point{400, 500}
	err = driver.DragPoint(startPoint, endPoint)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 200)

	startPointF := PointF{400.5, 501.5}
	endPointF := PointF{400.5, 261.5}
	err = driver.DragPointF(startPointF, endPointF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_TouchLongClick(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.TouchLongClick(400, 260, 1.2222)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 200)

	err = driver.TouchLongClickPoint(Point{X: 400, Y: 260})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_SendKeys(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.SendKeys("abc")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 2)

	err = driver.SendKeys("def", false)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 2)

	err = driver.SendKeys("\\n")
	// err = driver.SendKeys(`\n`, false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_PressBack(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.PressBack()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_PressKeyCode(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.PressKeyCodeAsync(KCx)
	if err != nil {
		t.Fatal(err)
	}
	err = driver.PressKeyCodeAsync(KCx, KMCapLocked)
	if err != nil {
		t.Fatal(err)
	}
	err = driver.PressKeyCodeAsync(KCExplorer)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_TouchDown(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	doTouchUp := func() {
		err = driver.TouchUp(400, 260)
		if err != nil {
			t.Fatal(err)
		}
	}

	SetDebug(true)

	err = driver.TouchDown(400, 260)
	if err != nil {
		t.Fatal(err)
	}

	// _ = driver.TapPoint(Point{400, 500})
	doTouchUp()

	err = driver.TouchDownPoint(Point{400, 260})
	if err != nil {
		t.Fatal(err)
	}

	doTouchUp()
}

func TestDriver_TouchUp(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	err = driver.TouchDown(400, 260)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	// err = driver.TouchUp(400, 260)
	err = driver.TouchUpPoint(Point{400, 260})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_TouchMove(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	doTouchDown := func(x, y int) {
		err = driver.TouchDown(x, y)
		if err != nil {
			t.Fatal(err)
		}
	}

	doTouchUp := func(x, y int) {
		err = driver.TouchUp(x, y)
		if err != nil {
			t.Fatal(err)
		}
	}

	doTouchDown(400, 260)

	SetDebug(true)

	err = driver.TouchMove(400, 500)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(false)

	doTouchUp(400, 500)

	doTouchDown(400, 500)

	SetDebug(true)

	err = driver.TouchMove(400, 260)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(false)

	doTouchUp(400, 260)

}

func TestDriver_OpenNotification(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.OpenNotification()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_Flick(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.Flick(50, -100)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_ScrollTo(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.ScrollTo(BySelector{ClassName: "android.widget.SeekBar"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_GetClipboard(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	text, err := driver.GetClipboard()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(text)
}

func TestDriver_SetClipboard(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	content := "test123"
	err = driver.SetClipboard(ClipDataTypePlaintext, content)
	if err != nil {
		t.Fatal(err)
	}

	text, err := driver.GetClipboard()
	if err != nil {
		t.Fatal(err)
	}
	if text != content {
		t.Fatal("should be the same")
	}
}

func TestDriver_AlertAccept(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.AlertAccept()
	// err = driver.AlertAccept("是")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_AlertDismiss(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	// err = driver.AlertDismiss()
	err = driver.AlertDismiss("否")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_SetAppiumSettings(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	appiumSettings, err := driver.GetAppiumSettings()
	if err != nil {
		t.Fatal(err)
	}
	sdopd := appiumSettings["shutdownOnPowerDisconnect"]
	t.Log("shutdownOnPowerDisconnect:", sdopd)

	SetDebug(true)

	err = driver.SetAppiumSettings(map[string]interface{}{"shutdownOnPowerDisconnect": !sdopd.(bool)})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(false)

	appiumSettings, err = driver.GetAppiumSettings()
	if err != nil {
		t.Fatal(err)
	}
	if appiumSettings["shutdownOnPowerDisconnect"] == sdopd.(bool) {
		t.Fatal("should not be equal")
	}
	t.Log("shutdownOnPowerDisconnect:", appiumSettings["shutdownOnPowerDisconnect"])
}

func TestDriver_SetOrientation(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.SetOrientation(OrientationLandscape)
	// err = driver.SetOrientation(OrientationPortrait)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_SetRotation(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	// err = driver.SetRotation(Rotation{Z: 0})
	err = driver.SetRotation(Rotation{Z: 270})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_NetworkConnection(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = driver.NetworkConnection(NetworkTypeWifi)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_FindElement(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	elem, err := driver.FindElement(BySelector{ResourceIdID: "android:id/content"})
	if err != nil {
		t.Fatal(err)
	}
	SetDebug(false)
	t.Log(elem.GetAttribute("class"))
}

func TestDriver_FindElements(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	// elements, err := driver.FindElements(BySelector{ResourceIdID: "com.android.settings:id/title"})
	elements, err := driver.FindElements(BySelector{UiAutomator: "new UiSelector().textStartsWith(\"应\");"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(elements))
}

func TestDriver_WaitWithTimeoutAndInterval(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}
	element, err := driver.FindElement(BySelector{UiAutomator: "new UiSelector().className(\"android.view.ViewGroup\");"})
	if err != nil {
		t.Fatal(err)
	}

	elem, err := element.FindElement(BySelector{UiAutomator: "new UiSelector().className(\"android.widget.LinearLayout\").index(6);"})
	if err != nil {
		t.Fatal(err)
	}

	rect, err := elem.Rect()
	if err != nil {
		t.Fatal(err)
	}

	x := rect.X + int(float64(rect.Width)*2)
	y := rect.Y + rect.Height/2
	err = driver.Tap(x, y)
	if err != nil {
		t.Fatal(err)
	}

	by := BySelector{UiAutomator: "new UiSelector().text(\"科技\");"}
	exists := func(d *Driver) (bool, error) {
		element, err = d.FindElement(by)
		if err == nil {
			return true, nil
		}
		return false, nil
	}

	err = driver.WaitWithTimeoutAndInterval(exists, 1, 0.1)
	if err != nil {
		t.Fatal(err)
	}

	// element, err = driver.FindElement(by)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	err = element.Click()
	if err != nil {
		t.Fatal(err)
	}

}
