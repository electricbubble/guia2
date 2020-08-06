package guia2

import (
	"io/ioutil"
	"testing"
)

func TestElement_Text(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	text, err := elem.Text()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)
}

func TestElement_GetAttribute(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	attribute, err := elem.GetAttribute("class")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(attribute)
}

func TestElement_ContentDescription(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/search"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	name, err := elem.ContentDescription()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(name)
}

func TestElement_Size(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/search"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	size, err := elem.Size()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(size)
}

func TestElement_Rect(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	rect, err := elem.Rect()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rect)
}

func TestElement_Screenshot(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	screenshot, err := elem.Screenshot()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ioutil.WriteFile("/Users/hero/Desktop/e1.png", screenshot.Bytes(), 0600))
}

func TestElement_Location(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	location, err := elem.Location()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(location)
}

func TestElement_Click(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = elem.Click()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Clear(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "android:id/search_src_text"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = elem.Clear()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_SendKeys(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "android:id/search_src_text"})
	if err != nil {
		t.Fatal(err)
	}

	// return

	SetDebug(true)

	// err = elem.SendKeys("abc")
	err = elem.SendKeys("456", false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_FindElements(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	parentElem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/main_content"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	elements, err := parentElem.FindElements(BySelector{ResourceIdID: "com.android.settings:id/category"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(elements))
}

func TestElement_FindElement(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	parentElem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/main_content"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	elem, err := parentElem.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(false)
	t.Log(elem.Text())
}

func TestElement_Swipe(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/category_title"})
	if err != nil {
		t.Fatal(err)
	}

	rect, err := elem.Rect()
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	t.Log(rect)

	var startX, startY, endX, endY int
	startX = rect.X + rect.Width/20
	startY = rect.Y + rect.Height/2
	endX = startX
	endY = startY - startY/2
	err = elem.Swipe(startX, startY, endX, endY)
	if err != nil {
		t.Fatal(err)
	}

	startPoint := PointF{X: float64(rect.X + rect.Width/20 + 30), Y: float64(startY / 2)}
	endPoint := PointF{X: startPoint.X, Y: startPoint.Y + startPoint.Y}
	err = elem.SwipePointF(startPoint, endPoint)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Drag(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elements, err := driver.FindElements(BySelector{ClassName: "android.widget.TextView"})
	if err != nil {
		t.Fatal(err)
	}

	for i, elem := range elements {
		text, _ := elem.Text()
		t.Log(i, text)
	}

	rect, err := elements[0].Rect()
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	// err = elements[0].Drag(300, 450, 256)
	err = elements[0].Drag(300, 450, 256)
	if err != nil {
		t.Fatal(err)
	}

	err = elements[0].DragTo(elements[1], 256)
	if err != nil {
		t.Fatal(err)
	}

	endPoint := PointF{X: float64(rect.X + rect.Width/3*2), Y: float64(rect.Y + rect.Height/2)}
	err = elements[0].DragPointF(endPoint, 256)
	if err != nil {
		t.Fatal()
	}
}

func TestElement_Flick(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	elem, err := driver.FindElement(BySelector{UiAutomator: "new UiSelector().text(\"提示音和通知\");"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = elem.Flick(36, 20, 100)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_ScrollTo(t *testing.T) {
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	// how to make it work?
	// parentElem, err := driver.FindElement(BySelector{ClassName: "android.widget.ScrollView"})
	// parentElem, err := driver.FindElement(BySelector{ResourceIdID: "com.cyanogenmod.filemanager:id/navigation_view_layout"})
	parentElem, err := driver.FindElement(BySelector{ResourceIdID: "com.android.settings:id/dashboard"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = parentElem.ScrollTo(BySelector{ContentDescription: "电池"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_ScrollToElement(t *testing.T) {
	// android.widget.HorizontalScrollView
	driver, err := NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	// how to make it work?
	parentElem, err := driver.FindElement(BySelector{UiAutomator: "new UiSelector().resourceId(\"com.android.settings:id/dashboard\");"})
	if err != nil {
		t.Fatal(err)
	}

	element, err := driver.FindElement(BySelector{UiAutomator: "new UiSelector().text(\"电池\");"})
	if err != nil {
		t.Fatal(err)
	}

	SetDebug(true)

	err = parentElem.ScrollToElement(element)
	if err != nil {
		t.Fatal(err)
	}
}
