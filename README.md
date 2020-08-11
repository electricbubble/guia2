# Golang-UIAutomator2

使用 Golang 实现 [appium/appium-uiautomator2-server](https://github.com/appium/appium-uiautomator2-server) 的客户端库

## 安装
```bash
go get -u github.com/electricbubble/guia2
```

## 使用

> 首次使用需要在 `Android` 设备中安装两个 `apk`  
> `appium-uiautomator2-server-debug-androidTest.apk`  
> `appium-uiautomator2-server-vXX.XX.XX.apk`
>  
> 再通过 `adb` 启动 `appium-uiautomator2-server`  
> ```bash
> adb shell am instrument -w io.appium.uiautomator2.server.test/androidx.test.runner.AndroidJUnitRunner
> ```

```go
package main

import (
	"github.com/electricbubble/guia2"
	"log"
)

func main() {
	// driver, err := guia2.NewDriver(guia2.NewEmptyCapabilities(), "http://localhost:6790/wd/hub")
	driver, err := guia2.NewDriver(nil, "http://localhost:6790/wd/hub")
	checkErr(err)

	// fmt.Println(driver.Source())
	// return

	deviceSize, err := driver.DeviceSize()
	checkErr(err)

	var startX, startY, endX, endY int
	startX = deviceSize.Width / 2
	startY = deviceSize.Height / 2
	endX = startX
	endY = startY / 2
	err = driver.Swipe(startX, startY, endX, endY)
	checkErr(err)

	var startPoint, endPoint guia2.PointF
	startPoint = guia2.PointF{X: float64(startX), Y: float64(startY)}
	endPoint = guia2.PointF{X: startPoint.X, Y: startPoint.Y * 1.6}
	err = driver.SwipePointF(startPoint, endPoint)
	checkErr(err)

	element, err := driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().textStartsWith(\"MIUI\");"})
	checkErr(err)

	err = element.Click()
	checkErr(err)

	element, err = driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().textStartsWith(\"查看更多\");"})
	checkErr(err)

	checkErr(element.Click())

	exists := func(d *guia2.Driver) (bool, error) {
		element, err = driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().text(\"关注\");"})
		if err == nil {
			return true, nil
		}
		return false, nil
	}
	err = driver.Wait(exists)
	checkErr(err)

	element, err = driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().textContains(\" 图像\");"})
	checkErr(err)

	checkErr(element.Click())

	err = driver.ScrollTo(guia2.BySelector{UiAutomator: "new UiSelector().textContains(\"全部评论\");"})
	checkErr(err)

	// element, err = driver.FindElement(guia2.BySelector{ResourceIdID: "cn.xuexi.android:id/TOP_LAYER_VIEW_ID"})
	// checkErr(err)
	// elemBack, err := element.FindElement(guia2.BySelector{ClassName: "android.widget.ImageView"})
	// checkErr(err)

	// screenshot, err := elem.Screenshot()
	// checkErr(err)
	// ioutil.WriteFile("/path/Desktop/e1.png", screenshot.Bytes(), 0600)

}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}

```

> 以上代码仅使用了 `小米8` 进行了测试。


![example](https://github.com/electricbubble/ImageHosting/blob/master/img/202008091517_guia2.gif)


## TODO

- [x] LongPressKeyCode
- [x] PressKeyCode
- [x] MultiPointerGesture
- [ ] W3CActions

## Thanks

Thank you [JetBrains](https://www.jetbrains.com/?from=gwda) for providing free open source licenses
