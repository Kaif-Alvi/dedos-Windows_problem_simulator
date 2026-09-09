package main

import (
	"embed"
	"unsafe"

	webview "github.com/webview/webview_go"
	"golang.org/x/sys/windows"
)

//go:embed assets/*.html
var assets embed.FS

var (
	user32 = windows.NewLazySystemDLL("user32.dll")
	gdi32  = windows.NewLazySystemDLL("gdi32.dll")

	procRegisterHotKey   = user32.NewProc("RegisterHotKey")
	procUnregisterHotKey = user32.NewProc("UnregisterHotKey")
	procGetMessageW      = user32.NewProc("GetMessageW")
	procSetWindowLongPtr = user32.NewProc("SetWindowLongPtrW")
	procGetWindowLongPtr = user32.NewProc("GetWindowLongPtrW")
	procSetWindowPos     = user32.NewProc("SetWindowPos")
	procGetSystemMetrics = user32.NewProc("GetSystemMetrics")
	procShowWindow       = user32.NewProc("ShowWindow")
	procIsZoomed         = user32.NewProc("IsZoomed")
	procSetClassLongPtr  = user32.NewProc("SetClassLongPtrW")
	procCreateSolidBrush = gdi32.NewProc("CreateSolidBrush")
)

const (
	MOD_ALT     = 0x0001
	MOD_CONTROL = 0x0002
	MOD_WIN     = 0x0008

	VK_C = 0x43
	VK_P = 0x50
	VK_R = 0x52

	GWL_STYLE = ^uintptr(15)

	WS_CAPTION     = 0x00C00000
	WS_THICKFRAME  = 0x00040000
	WS_MINIMIZEBOX = 0x00020000
	WS_MAXIMIZEBOX = 0x00010000
	WS_SYSMENU     = 0x00080000

	SWP_FRAMECHANGED = 0x0020
	SWP_NOZORDER     = 0x0004
	SWP_NOACTIVATE   = 0x0010
	SWP_NOMOVE       = 0x0002
	SWP_NOSIZE       = 0x0001

	SM_CXSCREEN = 0
	SM_CYSCREEN = 1

	HOTKEY_ID_EXIT   = 1
	HOTKEY_ID_PAUSE  = 2
	HOTKEY_ID_RESUME = 3

	WM_HOTKEY = 0x0312

	SW_MINIMIZE = 6
	SW_MAXIMIZE = 3
	SW_RESTORE  = 9

	GCL_HBRBACKGROUND = ^uintptr(9) // -10, two's complement
)

// titleBar is prepended to normal (non-prank) pages only.
// -webkit-app-region:drag makes the bar draggable with no extra Win32 code.
const titleBar = `<div style="-webkit-app-region:drag;height:34px;width:100%;background:#111318;display:flex;align-items:center;
justify-content:flex-end;position:fixed;top:0;left:0;z-index:99999;font-family:'Segoe UI',sans-serif;font-weight:700;font-size:14px;
color:#eaeaea;user-select:none;border-bottom:1px solid #00aaff"> 
  <span style="-webkit-app-region:drag;flex:1;padding-left:14px;letter-spacing:1px;color:#5ad1ff">DEDOS</span> 
  <span onclick="goMinimize()" style="-webkit-app-region:no-drag;cursor:pointer;padding:8px 16px"
   onmouseover="this.style.background='rgba(90,209,255,0.2)'" onmouseout="this.style.background='transparent'">&#8212;</span>
  <span onclick="goMaximize()" style="-webkit-app-region:no-drag;cursor:pointer;padding:8px 16px" 
  onmouseover="this.style.background='rgba(90,209,255,0.2)'" onmouseout="this.style.background='transparent'">&#9633;</span> 
  <span onclick="goExitApp()" style="-webkit-app-region:no-drag;cursor:pointer;padding:8px 16px" 
  onmouseover="this.style.background='#e81123'" onmouseout="this.style.background='transparent'">&#10005;</span> 
</div>`

// splashHTML is shown the instant the window exists, before the real page
// loads, so the user sees a themed splash instead of a white box.
const splashHTML = `<body style="margin:0;height:100vh;background:#0d0f14;display:flex;
align-items:center;justify-content:center;font-family:'Segoe UI',sans-serif;color:#5ad1ff;
font-size:28px;font-weight:700;letter-spacing:2px">DEDOS</body>`

var (
	w             webview.WebView
	hwnd          uintptr
	originalStyle uintptr
	isFullscreen  bool
)

type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      struct{ X, Y int32 }
}

func main() {
	w = webview.New(false)
	defer w.Destroy()

	w.SetTitle("Dedos - Windows Problem Simulator")
	w.SetSize(900, 600, webview.HintNone)
	w.SetHtml(splashHTML) // paint instantly, avoids white box on startup

	hwnd = uintptr(w.Window())
	setDarkBackground(hwnd)
	removeCaption(hwnd)

	registerBindings()
	loadPage("home") // replaces the splash with the real page

	go startHotkeyListener()

	w.Run()
}

// setDarkBackground swaps the window class background brush from the
// default white to dark, so any repaint before HTML is ready is dark too.
func setDarkBackground(h uintptr) {
	const darkColorRef = 0x00141110 // 0x00BBGGRR
	brush, _, _ := procCreateSolidBrush.Call(uintptr(darkColorRef))
	procSetClassLongPtr.Call(h, GCL_HBRBACKGROUND, brush)
}

// removeCaption strips the native title bar/system menu once at startup,
// keeping the resize border, so our own HTML title bar takes over.
func removeCaption(h uintptr) {
	style, _, _ := procGetWindowLongPtr.Call(h, uintptr(GWL_STYLE))
	originalStyle = style &^ uintptr(WS_CAPTION)
	procSetWindowLongPtr.Call(h, uintptr(GWL_STYLE), originalStyle)
	procSetWindowPos.Call(h, 0, 0, 0, 0, 0,
		uintptr(SWP_FRAMECHANGED|SWP_NOZORDER|SWP_NOACTIVATE|SWP_NOMOVE|SWP_NOSIZE))
}

func toggleMaximize() {
	zoomed, _, _ := procIsZoomed.Call(hwnd)
	if zoomed != 0 {
		procShowWindow.Call(hwnd, uintptr(SW_RESTORE))
	} else {
		procShowWindow.Call(hwnd, uintptr(SW_MAXIMIZE))
	}
}

func registerBindings() {
	w.Bind("goLoadPage", func(name string) {
		w.Dispatch(func() {
			loadPage(name)
		})
	})

	w.Bind("goTriggerBSOD", func() {
		w.Dispatch(func() { triggerFullscreenSim("bsod") })
	})
	w.Bind("goTriggerBlackScreen", func() {
		w.Dispatch(func() { triggerFullscreenSim("blackscreen") })
	})
	w.Bind("goTriggerCorruptBios", func() {
		w.Dispatch(func() { triggerFullscreenSim("corruptbios") })
	})
	w.Bind("goTriggerCpuOverload", func() {
		w.Dispatch(func() { triggerFullscreenSim("cpuoverload") })
	})
	w.Bind("goTriggerFakeUpdate", func() {
		w.Dispatch(func() { triggerFullscreenSim("fakeupdate") })
	})
	w.Bind("goTriggerRansomware", func() {
		w.Dispatch(func() { triggerFullscreenSim("ransomware") })
	})
	w.Bind("goTriggerDriverFailure", func() {
		w.Dispatch(func() { triggerFullscreenSim("driverfailure") })
	})
	w.Bind("goTriggerDiskBootFailure", func() {
		w.Dispatch(func() { triggerFullscreenSim("diskbootfailure") })
	})
	w.Bind("goTriggerLowDiskSpace", func() {
		w.Dispatch(func() { triggerFullscreenSim("lowdiskspace") })
	})
	w.Bind("goTriggerNetworkDisconnected", func() {
		w.Dispatch(func() { triggerFullscreenSim("networkdisconnected") })
	})

	w.Bind("goTriggerProblem", func(kind string) {
		w.Dispatch(func() {
			triggerFullscreenSim(kind)
		})
	})

	w.Bind("goExitApp", func() {
		w.Dispatch(func() {
			exitApp()
		})
	})

	w.Bind("goMinimize", func() {
		w.Dispatch(func() {
			procShowWindow.Call(hwnd, uintptr(SW_MINIMIZE))
		})
	})

	w.Bind("goMaximize", func() {
		w.Dispatch(func() {
			toggleMaximize()
		})
	})
}

func loadPage(name string) {
	restoreWindow()

	data, err := assets.ReadFile("assets/" + name + ".html")
	if err != nil {
		w.SetHtml(titleBar + "<h1 style='font-family:sans-serif;color:red'>File not found: " + name + "</h1>")
		return
	}
	// titleBar only goes on normal app pages; prank/fullscreen sim screens
	// (triggerFullscreenSim) intentionally stay without it.
	w.SetHtml(titleBar + string(data))
}

func triggerFullscreenSim(kind string) {
	makeFullscreen()

	switch kind {
	case "bsod":
		data, err := assets.ReadFile("assets/bsod.html")
		if err == nil {
			w.SetHtml(string(data))
		}
	case "blackscreen":
		w.SetHtml(blackScreenHTML)
	case "corruptbios":
		w.SetHtml(corruptBiosHTML)
	case "cpuoverload":
		w.SetHtml(cpuOverloadHTML)
	case "fakeupdate":
		w.SetHtml(fakeUpdateHTML)
	case "ransomware":
		w.SetHtml(ransomwareHTML)
	case "driverfailure":
		w.SetHtml(driverFailureHTML)
	case "diskbootfailure":
		w.SetHtml(diskBootFailureHTML)
	case "lowdiskspace":
		w.SetHtml(lowDiskSpaceHTML)
	case "networkdisconnected":
		w.SetHtml(networkDisconnectedHTML)

	// sims3.go simulations (3 to 10)
	case "updatestuck":
		w.SetHtml(updateStuckHTML)
	case "activationexpiry":
		w.SetHtml(activationExpiryHTML)
	case "printeroffline":
		w.SetHtml(printerOfflineHTML)
	case "audioservice":
		w.SetHtml(audioServiceHTML)
	case "ssdwarning":
		w.SetHtml(ssdWarningHTML)
	case "apphang":
		w.SetHtml(appHangHTML)
	case "gpuvector":
		w.SetHtml(gpuVectorHTML)
	case "defenderalert":
		w.SetHtml(defenderAlertHTML)
	case "criticalprocess":
		w.SetHtml(criticalProcessHTML)
	case "bootloop":
		w.SetHtml(bootLoopHTML)
	case "outofmemory":
		w.SetHtml(outOfMemoryHTML)
	case "explorercrash":
		w.SetHtml(explorerCrashHTML)
	case "nointernet":
		w.SetHtml(noInternetHTML)
	case "overheating":
		w.SetHtml(overheatingHTML)
	case "autorepair":
		w.SetHtml(autoRepairHTML)
	case "watermark":
		w.SetHtml(watermarkHTML)
	case "registryerror":
		w.SetHtml(registryErrorHTML)
	case "usbmalfunction":
		w.SetHtml(usbMalfunctionHTML)
	case "dnsfailure":
		w.SetHtml(dnsFailureHTML)
	case "ipconflict":
		w.SetHtml(ipConflictHTML)
	case "bluetoothfail":
		w.SetHtml(bluetoothFailHTML)
	case "cammissing":
		w.SetHtml(camMissingHTML)
	case "micmuted":
		w.SetHtml(micMutedHTML)
	case "batterylow":
		w.SetHtml(batteryLowHTML)
	case "gpudrivercrash":
		w.SetHtml(gpuDriverCrashHTML)
	case "directxerror":
		w.SetHtml(directxErrorHTML)
	case "dllmissing":
		w.SetHtml(dllMissingHTML)
	case "appfault":
		w.SetHtml(appFaultHTML)
	case "malwaredetected":
		w.SetHtml(malwareDetectedHTML)
	case "firewallblock":
		w.SetHtml(firewallBlockHTML)
	case "syncfailed":
		w.SetHtml(syncFailedHTML)
	case "pinblocked":
		w.SetHtml(pinBlockedHTML)
	case "passwordexpired":
		w.SetHtml(passwordExpiredHTML)
	case "servicefailed":
		w.SetHtml(serviceFailedHTML)
	case "taskschedulererror":
		w.SetHtml(taskSchedulerErrorHTML)
	case "environmentvariable":
		w.SetHtml(environmentVariableHTML)
	case "grouppolicy":
		w.SetHtml(groupPolicyHTML)
	case "kernelpanic":
		w.SetHtml(kernelPanicHTML)
	case "ramparity":
		w.SetHtml(ramParityErrorHTML)
	case "smartfailure":
		w.SetHtml(smartFailureHTML)
	case "powersurge":
		w.SetHtml(powerSurgeHTML)
	case "biosflashfail":
		w.SetHtml(biosFlashFailHTML)
	case "opticalerror":
		w.SetHtml(opticalErrorHTML)
	case "spoolerdeadlock":
		w.SetHtml(spoolerDeadlockHTML)
	case "gatewayunreachable":
		w.SetHtml(gatewayUnreachableHTML)
	case "vpndropped":
		w.SetHtml(vpnDroppedHTML)
	case "cloudsyncconflict":
		w.SetHtml(cloudsyncConflictHTML)
	case "zerodaywarning":
		w.SetHtml(zerodayWarningHTML)
	case "cputhrottle":
		w.SetHtml(cpuThrottleHTML)
	case "pagefilecorrupt":
		w.SetHtml(pagefileCorruptHTML)
	case "masterbootrecord":
		w.SetHtml(masterBootRecordHTML)
	case "securebootviolation":
		w.SetHtml(secureBootViolationHTML)
	case "tplockout":
		w.SetHtml(tpmLockoutHTML)
	case "bitlockerlock":
		w.SetHtml(bitlockerLockHTML)
	case "ghosttouch":
		w.SetHtml(ghostTouchHTML)
	case "ambientlightsensor":
		w.SetHtml(ambientLightSensorHTML)
	case "fingerprinthardware":
		w.SetHtml(fingerprintHardwareHTML)
	case "nvmecontroller":
		w.SetHtml(nvmeControllerHTML)
	case "gpuoverheat":
		w.SetHtml(gpuOverheatHTML)
	case "ramtimingmismatch":
		w.SetHtml(ramTimingMismatchHTML)
	case "psufanfailure":
		w.SetHtml(psuFanFailureHTML)
	case "hddbadsectors":
		w.SetHtml(hddBadSectorsHTML)
	case "networkloop":
		w.SetHtml(networkLoopHTML)
	case "dnspoisoning":
		w.SetHtml(dnsPoisoningHTML)
	case "audiocodeccrash":
		w.SetHtml(audioCodecCrashHTML)
	case "webcamhook":
		w.SetHtml(webcamHookHTML)
	case "registryhivecorrupt":
		w.SetHtml(registryHiveCorruptHTML)
	case "grouppolicyconflict":
		w.SetHtml(groupPolicyConflictHTML)
	case "servicehang":
		w.SetHtml(serviceHangHTML)
	case "certificateerror":
		w.SetHtml(certificateErrorHTML)
	case "proxyerror":
		w.SetHtml(proxyErrorHTML)
	case "indexingservicedead":
		w.SetHtml(indexingServiceDeadHTML)
	case "printdriverconflict":
		w.SetHtml(printDriverConflictHTML)
	case "bluetoothstackfault":
		w.SetHtml(bluetoothStackFaultHTML)
	case "usbhubpowerfault":
		w.SetHtml(usbHubPowerFaultHTML)
	case "kernelsecuritycheck":
		w.SetHtml(kernelSecurityCheckHTML)
	case "biosupdatefailure":
		w.SetHtml(biosUpdateFailureHTML)
	case "batterycritical":
		w.SetHtml(batteryCriticalHTML)
	case "displaydrivercrash":
		w.SetHtml(displayDriverCrashHTML)
	case "wifiauthfailed":
		w.SetHtml(wifiAuthFailedHTML)
	case "gatewaytimeout":
		w.SetHtml(gatewayTimeoutHTML)
	case "antivirusthreat":
		w.SetHtml(antivirusThreatHTML)
	case "dbconnectionlost":
		w.SetHtml(dbConnectionLostHTML)
	case "appcrash", "memoryerror", "systemfiles", "loginloop", "webcam", "bluetooth", "printer", "sound", "displaydriver", "explorer", "searchindex", "permissions", "date", "storage", "battery", "usb", "restore", "activation":
		data, err := assets.ReadFile("assets/" + kind + ".html")
		if err == nil {
			w.SetHtml(string(data))
		}
	}
}

func makeFullscreen() {
	if isFullscreen {
		return
	}

	style, _, _ := procGetWindowLongPtr.Call(hwnd, uintptr(GWL_STYLE))
	originalStyle = style

	newStyle := style &^ uintptr(WS_CAPTION|WS_THICKFRAME|WS_MINIMIZEBOX|WS_MAXIMIZEBOX|WS_SYSMENU)
	procSetWindowLongPtr.Call(hwnd, uintptr(GWL_STYLE), newStyle)

	screenW, _, _ := procGetSystemMetrics.Call(uintptr(SM_CXSCREEN))
	screenH, _, _ := procGetSystemMetrics.Call(uintptr(SM_CYSCREEN))

	procSetWindowPos.Call(
		hwnd, 0,
		0, 0, screenW, screenH,
		uintptr(SWP_FRAMECHANGED|SWP_NOZORDER|SWP_NOACTIVATE),
	)

	isFullscreen = true
}

func restoreWindow() {
	if !isFullscreen {
		return
	}

	procSetWindowLongPtr.Call(hwnd, uintptr(GWL_STYLE), originalStyle)
	procSetWindowPos.Call(
		hwnd, 0,
		100, 100, 900, 600,
		uintptr(SWP_FRAMECHANGED|SWP_NOZORDER|SWP_NOACTIVATE),
	)

	isFullscreen = false
}

func startHotkeyListener() {
	procRegisterHotKey.Call(0, uintptr(HOTKEY_ID_EXIT), uintptr(MOD_CONTROL|MOD_ALT|MOD_WIN), uintptr(VK_C))
	procRegisterHotKey.Call(0, uintptr(HOTKEY_ID_PAUSE), uintptr(MOD_CONTROL|MOD_ALT|MOD_WIN), uintptr(VK_P))
	procRegisterHotKey.Call(0, uintptr(HOTKEY_ID_RESUME), uintptr(MOD_CONTROL|MOD_ALT|MOD_WIN), uintptr(VK_R))

	defer procUnregisterHotKey.Call(0, uintptr(HOTKEY_ID_EXIT))
	defer procUnregisterHotKey.Call(0, uintptr(HOTKEY_ID_PAUSE))
	defer procUnregisterHotKey.Call(0, uintptr(HOTKEY_ID_RESUME))

	var msg MSG
	for {
		ret, _, _ := procGetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if int32(ret) <= 0 {
			break
		}
		if msg.Message == WM_HOTKEY {
			switch msg.WParam {
			case HOTKEY_ID_EXIT:
				exitApp()
			case HOTKEY_ID_PAUSE:
				pauseSim()
			case HOTKEY_ID_RESUME:
				resumeSim()
			}
		}
	}
}

func exitApp() {
	w.Dispatch(func() {
		restoreWindow()
		w.Terminate()
	})
}

func pauseSim() {
	w.Dispatch(func() {
		w.Eval("if (typeof window.dedosPause === 'function') { window.dedosPause(); }")
	})
}

func resumeSim() {
	w.Dispatch(func() {
		w.Eval("if (typeof window.dedosResume === 'function') { window.dedosResume(); }")
	})
}
