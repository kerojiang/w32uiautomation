package w32uiautomation

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation6 struct {
	*IUIAutomation5
}

type IUIAutomation6Vtbl struct {
	*IUIAutomation5Vtbl
	AddActiveTextPositionChangedEventHandler    uintptr
	AddEventHandlerGroup                        uintptr
	CreateEventHandlerGroup                     uintptr
	Get_CoalesceEvents                          uintptr
	Get_ConnectionRecoveryBehavior              uintptr
	Put_CoalesceEvents                          uintptr
	Put_ConnectionRecoveryBehavior              uintptr
	RemoveActiveTextPositionChangedEventHandler uintptr
	RemoveEventHandlerGroup                     uintptr
}

// IID为AAE072DA-29E3-413D-87A7-192DBF81ED10
var IID_IUIAutomation6 = ole.NewGUID("AAE072DA-29E3-413D-87A7-192DBF81ED10")

func NewUIAutomation6() (*IUIAutomation6, error) {
	result, err := ole.CreateInstance(CLSID_CUIAutomation8, IID_IUIAutomation6)
	if err != nil {
		return nil, err
	}

	auto6 := (*IUIAutomation6)(unsafe.Pointer(result))

	auto6.IUIAutomation5, err = NewUIAutomation5()
	if err != nil {
		return nil, err
	}
	return auto6, nil
}

func (auto6 *IUIAutomation6) VTable() *IUIAutomation6Vtbl {
	// return (*IUIAutomation6Vtbl)(unsafe.Pointer(auto6.RawVTable))
	return nil
}
