package guia2

import (
	"testing"
)

func TestUiSelectorHelper_NewUiSelectorHelper(t *testing.T) {
	uiSelector := NewUiSelectorHelper().Text("a").String()
	if uiSelector != `new UiSelector().text("a");` {
		t.Fatal("[ERROR]", uiSelector)
	}

	uiSelector = NewUiSelectorHelper().Text("a").TextStartsWith("b").String()
	if uiSelector != `new UiSelector().text("a").textStartsWith("b");` {
		t.Fatal("[ERROR]", uiSelector)
	}

	uiSelector = NewUiSelectorHelper().ClassName("android.widget.LinearLayout").Index(6).String()
	if uiSelector != `new UiSelector().className("android.widget.LinearLayout").index(6);` {
		t.Fatal("[ERROR]", uiSelector)
	}

	uiSelector = NewUiSelectorHelper().Focused(false).Instance(6).String()
	if uiSelector != `new UiSelector().focused(false).instance(6);` {
		t.Fatal("[ERROR]", uiSelector)
	}

	uiSelector = NewUiSelectorHelper().ChildSelector(NewUiSelectorHelper().Enabled(true)).String()
	if uiSelector != `new UiSelector().childSelector(new UiSelector().enabled(true));` {
		t.Fatal("[ERROR]", uiSelector)
	}

}

func Test_getFreePort(t *testing.T) {
	freePort, err := getFreePort()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(freePort)
}

func TestDeviceList(t *testing.T) {
	devices, err := DeviceList()
	if err != nil {
		t.Fatal(err)
	}
	for i := range devices {
		t.Log(devices[i].Serial())
	}
}

func TestNewUSBDriver(t *testing.T) {
	SetDebug(true)

	driver, err := NewUSBDriver()
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Dispose()

	ready, err := driver.Status()
	if err != nil {
		t.Fatal(err)
	}
	if !ready {
		t.Fatal("should be 'true'")
	}
}

func TestDriver_ActiveAppPackageName(t *testing.T) {
	devices, err := DeviceList()
	if err != nil {
		t.Fatal(err)
	}
	dev := devices[len(devices)-1]

	// SetDebug(true)

	driver, err := NewUSBDriver(dev)
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Dispose()

	appPackageName, err := driver.ActiveAppPackageName()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(appPackageName)
}

func TestDriver_AppLaunch(t *testing.T) {
	driver, err := NewUSBDriver()
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Dispose()

	// err = driver.AppLaunch("tv.danmaku.bili", BySelector{ResourceIdID: "tv.danmaku.bili:id/action_bar_root"})
	err = driver.AppLaunch("com.android.settings", BySelector{ResourceIdID: "android:id/list"})
	if err != nil {
		t.Fatal(err)
	}

	// screenshot, err := driver.Screenshot()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(ioutil.WriteFile("/Users/hero/Desktop/s1.png", screenshot.Bytes(), 0600))
}

func TestDriver_AppTerminate(t *testing.T) {
	driver, err := NewUSBDriver()
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Dispose()

	err = driver.AppTerminate("tv.danmaku.bili")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewWiFiDriver(t *testing.T) {
	driver, err := NewWiFiDriver("192.168.1.28")
	if err != nil {
		t.Fatal(err)
	}

	// SetDebug(false, true)
	_, err = driver.ActiveAppActivity()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_AppInstall(t *testing.T) {
	driver, err := NewUSBDriver()
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Dispose()

	err = driver.AppInstall("/Users/hero/Desktop/xuexi_android_10002068.apk")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriver_AppUninstall(t *testing.T) {
	driver, err := NewUSBDriver()
	if err != nil {
		t.Fatal(err)
	}
	defer driver.Dispose()

	err = driver.AppUninstall("cn.xuexi.android")
	if err != nil {
		t.Fatal(err)
	}
}
