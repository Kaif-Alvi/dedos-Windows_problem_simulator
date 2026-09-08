package main

// sims6.go - Contains HTML strings for simulations 31 to 40

const dnsFailureHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>DNS Failure</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1e1e1e; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #252526; border: 1px solid #333; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #bbb; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🌐 DNS Server Not Responding</div>
    <div class="desc">The DNS server might be unavailable or your network configuration is incorrect. Error code: DNS_PROBE_FINISHED_NXDOMAIN.</div>
    <button class="btn" onclick="location.reload()">Flush DNS</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const ipConflictHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>IP Conflict</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1e1e1e; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #252526; border: 1px solid #333; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #f48771; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #bbb; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">⚠️ IP Address Conflict Detected</div>
    <div class="desc">Another device on the local network is using the same IP address assigned to this computer.</div>
    <button class="btn" onclick="location.reload()">Renew IP</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const bluetoothFailHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Bluetooth Failure</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1e1e1e; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2d2d2d; border: 1px solid #3f3f3f; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">📶 Bluetooth Adapter Missing</div>
    <div class="desc">Bluetooth is not available on this device or the driver has stopped responding. Pair connections lost.</div>
    <button class="btn" onclick="location.reload()">Close</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const camMissingHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Camera Missing</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1e1e1e; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2d2d2d; border: 1px solid #3f3f3f; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #f14c4c; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">📷 Webcam Error (0xA00F4244)</div>
    <div class="desc">We can't find your camera. Check if it is securely connected or used by another application.</div>
    <button class="btn" onclick="location.reload()">Close</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const micMutedHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Microphone Error</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1e1e1e; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2d2d2d; border: 1px solid #3f3f3f; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🎙️ Microphone Access Blocked</div>
    <div class="desc">Privacy settings are preventing applications from accessing your default recording device.</div>
    <button class="btn" onclick="location.reload()">Settings</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const batteryLowHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Battery Low</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #221515; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #331a1a; border: 1px solid #552222; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔋 Critical Battery Level (4%)</div>
    <div class="desc">Your computer is about to go into hibernation due to critically low battery power. Plug in charger immediately.</div>
    <button class="btn" onclick="location.reload()">Hibernate Now</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const gpuDriverCrashHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>GPU Driver Crash</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #f48771; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🖥️ Display Driver Stopped Responding</div>
    <div class="desc">NVIDIA / AMD Windows Kernel Mode Driver has crashed and successfully recovered from a fatal exception.</div>
    <button class="btn" onclick="location.reload()">OK</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const directxErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>DirectX Error</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #f14c4c; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🎮 DirectX Fatal Error</div>
    <div class="desc">DX11 feature level 10.0 is required to run the engine. Check your hardware graphics capabilities.</div>
    <button class="btn" onclick="location.reload()">Close</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const dllMissingHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>DLL Missing</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #f48771; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">⚠️ MSVCP140.dll Missing</div>
    <div class="desc">The code execution cannot proceed because MSVCP140.dll was not found. Reinstalling the application may fix this problem.</div>
    <button class="btn" onclick="location.reload()">OK</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const appFaultHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Application Fault</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #f14c4c; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">❌ AppCrash (Fault Module Name)</div>
    <div class="desc">Faulting application name: app.exe, version: 1.0.0.0, faulting module: ntdll.dll, exception code: 0xc0000005.</div>
    <button class="btn" onclick="location.reload()">Close Program</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`