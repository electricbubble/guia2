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
