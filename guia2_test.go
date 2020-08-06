package guia2

import "testing"

func TestBySelector_getMethodAndSelector(t *testing.T) {
	testVal := "test id"
	bySelector := BySelector{ResourceIdID: testVal}
	method, selector := bySelector.getMethodAndSelector()
	if method != "id" || selector != testVal {
		t.Fatal(method, "=", selector)
	}

	bySelector = BySelector{ContentDescription: testVal}
	method, selector = bySelector.getMethodAndSelector()
	if method != "accessibility id" || selector != testVal {
		t.Fatal(method, "=", selector)
	}
}
