//go:build windows

package main

import "syscall"

func hideConsole() {
	getConsoleWindow := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	showWindow := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	if getConsoleWindow.Find() == nil && showWindow.Find() == nil {
		hwnd, _, _ := getConsoleWindow.Call()
		if hwnd != 0 {
			// SW_HIDE = 0
			// 0 を渡すとウィンドウが非表示になる
			showWindow.Call(hwnd, 0)
		}
	}
}
