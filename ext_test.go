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

	err = driver.AppLaunch("tv.danmaku.bili")
	if err != nil {
		t.Fatal(err)
	}
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
