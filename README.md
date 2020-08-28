# Golang-UIAutomator2
[![go doc](https://godoc.org/github.com/electricbubble/guia2?status.svg)](https://pkg.go.dev/github.com/electricbubble/guia2?tab=doc)
[![license](https://img.shields.io/github/license/electricbubble/guia2)](https://github.com/electricbubble/guia2/blob/master/LICENSE)

使用 Golang 实现 [appium/appium-uiautomator2-server](https://github.com/appium/appium-uiautomator2-server) 的客户端库

## 扩展库

- [electricbubble/guia2-ext-opencv](https://github.com/electricbubble/guia2-ext-opencv) 直接通过指定图片进行操作

## 安装
```bash
go get -u github.com/electricbubble/guia2
```

## 使用

> 首次使用需要在 `Android` 设备中安装两个 `apk`  
> `appium-uiautomator2-server-debug-androidTest.apk`  
> `appium-uiautomator2-server-vXX.XX.XX.apk`
>
>> `apk` 可以选择通过 [appium/appium-uiautomator2-server](https://github.com/appium/appium-uiautomator2-server#building-project) 进行构建  
>> 也可以直接从这里下载 [electricbubble/appium-uiautomator2-server-apk](https://github.com/electricbubble/appium-uiautomator2-server-apk/releases)
>  
>
> 再通过 `adb` 启动 `appium-uiautomator2-server`  
> ```shell script
> adb shell am instrument -w io.appium.uiautomator2.server.test/androidx.test.runner.AndroidJUnitRunner
> # ⬇️ 后台运行
> adb shell "nohup am instrument -w io.appium.uiautomator2.server.test/androidx.test.runner.AndroidJUnitRunner >/sdcard/uia2server.log 2>&1 &"
> # or
> adb -s $serial shell "nohup am instrument -w io.appium.uiautomator2.server.test/androidx.test.runner.AndroidJUnitRunner >/sdcard/uia2server.log 2>&1 &"
> ```

### `guia2.NewUSBDriver()`
该函数使用期间, `Android` 设备必须一直保持 `USB` 的连接

### `guia2.NewWiFiDriver("192.168.1.28")`
1. 先通过 `USB` 连接 `Android` 设备
2. 让设备在 5555 端口监听 TCP/IP 连接
    ```shell script
    adb tcpip 5555
   # or
    adb -s $serial tcpip 5555
    ```
3. 查询 `Android` 设备的 `IP` (这一步骤开始可选择断开 `USB` 连接)
4. 通过 `IP` 连接 `Android` 设备
    ```shell script
    adb connect $deviceIP
    ```
5. 确认连接状态
    ```shell script
    adb devices
    ```
    看到以下格式的设备, 说明连接成功
    ```shell script
    $deviceIP:5555    device
    ```

```go
package main

import (
	"github.com/electricbubble/guia2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// driver, err := guia2.NewDriver(guia2.NewEmptyCapabilities(), "http://localhost:6790/wd/hub")
	// driver, err := guia2.NewDriver(nil, "http://192.168.1.28:6790/wd/hub")
	driver, err := guia2.NewUSBDriver()
	// driver, err := guia2.NewWiFiDriver("192.168.1.28")
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

	element, err := driver.FindElement(guia2.BySelector{ResourceIdID: "tv.danmaku.bili:id/expand_search"})
	checkErr(err)

	err = element.Click()
	checkErr(err)

	bySelector := guia2.BySelector{UiAutomator: guia2.NewUiSelectorHelper().Focused(true).String()}
	exists := func(d *guia2.Driver) (bool, error) {
		element, err = d.FindElement(bySelector)
		if err == nil {
			return true, nil
		}
		return false, nil
	}

	err = driver.Wait(exists)
	checkErr(err)

	err = element.SendKeys("雾山五行")
	checkErr(err)

	err = driver.PressKeyCode(guia2.KCEnter, guia2.KMEmpty)
	checkErr(err)

	bySelector = guia2.BySelector{UiAutomator: guia2.NewUiSelectorHelper().TextStartsWith("番剧").String()}
	checkErr(driver.Wait(exists))
	checkErr(element.Click())

	bySelector = guia2.BySelector{UiAutomator: guia2.NewUiSelectorHelper().Text("立即观看").String()}
	checkErr(driver.Wait(exists))
	checkErr(element.Click())

	bySelector = guia2.BySelector{ResourceIdID: "tv.danmaku.bili:id/videoview_container_space"}
	checkErr(driver.Wait(exists))

	// time.Sleep(time.Second * 5)

	screenshot, err := element.Screenshot()
	checkErr(err)
	userHomeDir, _ := os.UserHomeDir()
	checkErr(ioutil.WriteFile(userHomeDir+"/Desktop/element.png", screenshot.Bytes(), 0600))

	err = driver.PressKeyCode(guia2.KCMediaPause, guia2.KMEmpty)
	checkErr(err)

	err = driver.PressBack()
	checkErr(err)
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

> 感谢小伙伴提供的 `红米 Note 5A`


![example](https://github.com/electricbubble/ImageHosting/blob/master/img/202008192034_guia2.gif)


## Thanks

Thank you [JetBrains](https://www.jetbrains.com/?from=gwda) for providing free open source licenses
