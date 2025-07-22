//go:build !windows

package main

// Windows以外のOSでは何もしない
func hideConsole() {}
