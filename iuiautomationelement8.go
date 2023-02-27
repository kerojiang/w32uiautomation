package w32uiautomation

import (
	"github.com/go-ole/go-ole"
)

type IUIAutomationElement8 struct {
	IUIAutomationElement7
}

type IUIAutomationElement8Vtbl struct {
	IUIAutomationElement7Vtbl
	Get_CachedHeadingLevel  uintptr
	Get_CurrentHeadingLevel uintptr
}

// IID为8C60217D-5411-4CDE-BCC0-1CEDA223830C
var IID_IUIAutomationElement8 = ole.NewGUID("8C60217D-5411-4CDE-BCC0-1CEDA223830C")
