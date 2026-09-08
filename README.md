<div align="center">

<img src="dedos_logo.ico" alt="Dedos Logo" width="120"/>

# Dedos

**A lightweight Windows problem/prank simulator — built with Go + WebView**

[![Platform](https://img.shields.io/badge/platform-Windows-0078D6?logo=windows)](https://github.com)
[![Go](https://img.shields.io/badge/built%20with-Go-00ADD8?logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

</div>

---

## About

**Dedos** is a native Windows desktop app that simulates realistic system problems — BSODs, fake updates, ransomware screens, driver failures, and 100+ other scenarios — for pranks, demos, or testing reactions in a safe, harmless environment.

Built entirely in Go using [`webview_go`](https://github.com/webview/webview_go), with a custom borderless UI and zero heavy dependencies.

---

## Screenshots

<table>
  <tr>
    <td><img src="assets/screenshots/screenshot1.png" alt="Home screen" width="400"/></td>
    <td><img src="assets/screenshots/screenshot2.png" alt="Problem selection" width="400"/></td>
  </tr>
  <tr>
    <td><img src="assets/screenshots/screenshot3.png" alt="Simulation running" width="400"/></td>
    <td><img src="assets/screenshots/screenshot4.png" alt="BSOD simulation" width="400"/></td>
  </tr>
</table>

---

## Features

- 🖥️ **100+ simulations** — BSOD, ransomware, driver failures, disk errors, network issues, and more
- 🎨 **Custom title bar** — clean minimize / maximize / close controls, no native chrome
- ⚡ **Instant load** — themed splash screen, no white-screen flash on startup
- ⌨️ **Global hotkeys** — `Ctrl+Alt+Win+C` to exit any simulation instantly
- 🪶 **Lightweight** — single `.exe`, no installer bloat, pure Go (no CGO runtime dependency at use-time)

---

## Installation

1. Download the latest `dedos.exe` from [Releases](../../releases)
2. Run it — no installation required
3. Press **Ctrl+Alt+Win+C** anytime to exit a running simulation

---

## Building from Source

```bash
git clone https://github.com/<your-username>/dedos.git
cd dedos/dedos
go mod tidy
go build -ldflags "-H=windowsgui" -o dedos.exe .
```

> Requires Go 1.21+ and a MinGW-w64 toolchain (for `webview_go`'s CGO WebView2 bindings) on Windows, or the provided GitHub Actions workflow (MSYS2/MinGW64) for CI builds.

### Embedding the App Icon

```bash
windres resource.rc -O coff -o resource.syso
go build -ldflags "-H=windowsgui" -o dedos.exe .
```

---

## Project Structure

```
dedos/
├── main.go              # entry point, window/title-bar logic, Win32 bindings
├── sim1.go – sim10.go    # simulation HTML constants
├── assets/*.html         # simulation screens (embedded via embed.FS)
├── dedos_logo.png        # app logo
├── dedos.ico             # compiled app icon
├── resource.rc           # Windows resource script (icon)
└── go.mod / go.sum
```

---

## Controls

| Action              | Shortcut / Method        |
|---------------------|---------------------------|
| Exit simulation     | `Ctrl + Alt + Win + C`    |
| Pause simulation     | `Ctrl + Alt + Win + P`    |
| Resume simulation    | `Ctrl + Alt + Win + R`    |
| Minimize / Maximize / Close | Custom title bar buttons |

---

## Disclaimer

Dedos is intended for **harmless pranks, demos, and educational purposes only**. It does not modify, damage, or access any system files or data. Use responsibly and only on devices you have permission to use.

---

## License

MIT © [Kaif Alvi](https://github.com/Kaif-Alvi)
