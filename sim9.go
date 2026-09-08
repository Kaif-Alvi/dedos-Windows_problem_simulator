package main

// sims9.go - Contains HTML strings for simulations 71 to 90 (without duplicates)

const gpuOverheatHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>GPU Overheating</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #3c1a1a; border: 1px solid #662222; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🌡️ GPU Thermal Limit Exceeded</div>
    <p class="desc">Graphics card temperature reached 105°C. Emergency hardware shutoff sequence initiated to prevent silicon damage.</p>
    <button class="btn" onclick="location.reload()">Emergency Exit</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const ramTimingMismatchHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>RAM Timing Mismatch</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">⚡ XMP Profile Training Failed</div>
    <p class="desc">Memory sticks running at unstable latency parameters. Reverting to default JEDEC 2133MHz speeds.</p>
    <button class="btn" onclick="location.reload()">Reset BIOS</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const psuFanFailureHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>PSU Fan Failure</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #3c1a1a; border: 1px solid #662222; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔌 Power Supply Fan Stalled</div>
    <p class="desc">Zero RPM tachometer feedback from PSU internal fan. Risk of overheating and component frying.</p>
    <button class="btn" onclick="location.reload()">Shut Down</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const hddBadSectorsHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Bad Sectors</title>
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
    <div class="title">⚠️ Uncorrectable Sector Cluster</div>
    <p class="desc">Logical block addressing failed on track 45102. Data allocation tables unable to map surface sectors.</p>
    <button class="btn" onclick="location.reload()">Run CHKDSK</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const networkLoopHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Network Loop</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔄 Ethernet Broadcast Storm</div>
    <p class="desc">Switch loop detected. Packet buffer saturation causing massive frame drops across local network interface.</p>
    <button class="btn" onclick="location.reload()">Disconnect</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const dnsPoisoningHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>DNS Poisoning</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #3c1a1a; border: 1px solid #662222; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🚨 Cache Poisoning Threat</div>
    <p class="desc">Rogue DNS responses intercepted. Domain name resolution redirected to unauthorized external IP addresses.</p>
    <button class="btn" onclick="location.reload()">Flush Cache</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const audioCodecCrashHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Audio Codec Crash</title>
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
    <div class="title">🔊 Realtek Audio Subsystem Fault</div>
    <p class="desc">Endpoint buffer underrun. Audio rendering hardware pipeline threw unhandled interrupt exception.</p>
    <button class="btn" onclick="location.reload()">Restart Audio</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const webcamHookHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Webcam Hook</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #3c1a1a; border: 1px solid #662222; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">📷 Unauthorized Camera Stream</div>
    <p class="desc">An untrusted background process hooked into DirectShow capture interface to record video feed.</p>
    <button class="btn" onclick="location.reload()">Block Access</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const registryHiveCorruptHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Registry Hive Corrupt</title>
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
    <div class="title">🛠️ SYSTEM Registry Corrupted</div>
    <p class="desc">The configuration registry database file was damaged and restored from a backup log successfully.</p>
    <button class="btn" onclick="location.reload()">Restore Point</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const groupPolicyConflictHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Group Policy Conflict</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">📜 Group Policy Object Error</div>
    <p class="desc">Domain controller policy processing failed. Security descriptors mismatch found on local template.</p>
    <button class="btn" onclick="location.reload()">Force Update</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const serviceHangHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Service Hang</title>
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
    <div class="title">⚙️ Windows Service Timeout</div>
    <p class="desc">The service did not respond to the start or control request in a timely fashion (Event ID 7000).</p>
    <button class="btn" onclick="location.reload()">Kill Service</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const certificateErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Certificate Error</title>
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
    <div class="title">🔒 SSL Certificate Expired</div>
    <p class="desc">The security certificate for this site has expired or is untrusted by root store authorities.</p>
    <button class="btn" onclick="location.reload()">Advanced</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const proxyErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Proxy Error</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🌐 Proxy Server Refused Connection</div>
    <p class="desc">The configured HTTP proxy server is not responding or rejecting downstream socket requests.</p>
    <button class="btn" onclick="location.reload()">Disable Proxy</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const indexingServiceDeadHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Indexing Service Dead</title>
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
    <div class="title">🔍 Windows Search Indexer Failed</div>
    <p class="desc">The Windows Search service database catalog is corrupted. File querying index has stopped.</p>
    <button class="btn" onclick="location.reload()">Rebuild Index</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const printDriverConflictHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Print Driver Conflict</title>
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
    <div class="title">🖨️ Printer Isolation Fault</div>
    <p class="desc">Kernel mode print driver loaded incompatible DLL module. Print isolation host crashed.</p>
    <button class="btn" onclick="location.reload()">Remove Driver</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const bluetoothStackFaultHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Bluetooth Stack Fault</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #2a2a2a; border: 1px solid #444; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔵 BTHUSB Radio Stack Error</div>
    <p class="desc">Bluetooth radio transceiver failed to handle LMP response timeout. Host controller reset required.</p>
    <button class="btn" onclick="location.reload()">Reset Adapter</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const usbHubPowerFaultHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>USB Hub Power Fault</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #3c1a1a; border: 1px solid #662222; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔌 USB Over-Current Surge Detected</div>
    <p class="desc">A USB device has exceeded the power port limit. Port hub shut down to protect mainboard circuits.</p>
    <button class="btn" onclick="location.reload()">Reset Port</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const kernelSecurityCheckHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Kernel Security Check</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #0078d7; height: 100vh; display: flex; flex-direction: column; justify-content: center; padding: 80px; font-family: 'Segoe UI', sans-serif; color: #fff; }
  h1 { font-size: 40px; font-weight: 300; margin-bottom: 20px; }
  p { font-size: 16px; line-height: 1.6; max-width: 700px; margin-bottom: 30px; }
  .code { font-family: monospace; font-size: 14px; opacity: 0.8; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: rgba(255,255,255,0.6); }
</style>
</head>
<body>
  <h1>:(</h1>
  <p>Your PC ran into a problem and needs to restart. We're just collecting some error info, and then we'll restart for you.</p>
  <p class="code">Stop Code: KERNEL_SECURITY_CHECK_FAILURE</p>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`