/**
直接使用api
*/
package main

import (
	"fmt"
	xcgui "github.com/xicheng/xcgui/xc"
	//"syscall"
)

func main() {
	xcgui.XInitXCGUI()
	hwnd := xcgui.XWndCreate(400, 200, 300, 200, "标题",
		xcgui.NULL, xcgui.XC_WINDOW_STYLE_DEFAULT)

	parent := xcgui.HXCGUI(hwnd)
	//button
	btn := xcgui.XBtnCreate(10, 5, 80, 22, "关闭", parent)
	xcgui.XBtnSetType(btn, xcgui.BUTTON_TYPE_CLOSE)
	//label
	lb := xcgui.XShapeTextCreate(50, 100, 100, 22, "hello world!", parent)
	xcgui.XShapeTextSetText(lb, "hello 世界!")
	xcgui.XShapeTextSetTextColor(lb, 0xff0000, 255)

	//取text及长度
	str := xcgui.XShapeTextGetTextGo(lb)
	fmt.Println(str)
	fmt.Println(xcgui.XShapeTextGetTextLength(lb))

	xcgui.XWndShowWindow(hwnd, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
