package main

// sim10.go - Contains HTML strings for simulations 91 to 100

const biosUpdateFailureHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>BIOS Update Failure</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #000; height: 100vh; display: flex; flex-direction: column; justify-content: center; align-items: center; font-family: monospace; color: #ff3333; }
  .box { width: 500px; border: 2px solid #ff3333; padding: 30px; background: #110000; text-align: center; }
  h2 { margin-bottom: 15px; font-size: 20px; }
  p { font-size: 14px; color: #ff9999; line-height: 1.5; margin-bottom: 20px; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #666; font-family: 'Segoe UI', sans-serif; }
</style>
</head>
<body>
  <div class="box">
    <h2>⚠️ BIOS Flash Verification Failed</h2>
    <p>EEPROM checksum mismatch detected during firmware update. Do not turn off your system to prevent permanent bricking.</p>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const batteryCriticalHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Battery Critical</title>
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
    <div class="title">🔋 Battery Critically Low (3%)</div>
    <p class="desc">Your computer will go into hibernation mode soon to prevent total power loss and data corruption.</p>
    <button class="btn" onclick="location.reload()">Plug in Charger</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const displayDriverCrashHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Display Driver Crash</title>
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
    <div class="title">🖥️ Display driver stopped responding</div>
    <p class="desc">Display driver NVIDIA Windows Kernel Mode Driver stopped responding and has successfully recovered.</p>
    <button class="btn" onclick="location.reload()">Dismiss</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const wifiAuthFailedHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Wi-Fi Auth Failed</title>
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
    <div class="title">📶 Wi-Fi Authentication Mismatch</div>
    <p class="desc">WPA3-Personal handshake failed. Pre-shared key authentication rejected by access point.</p>
    <button class="btn" onclick="location.reload()">Re-enter Key</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const gatewayTimeoutHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Gateway Timeout</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #f8f9fa; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #333; }
  .container { text-align: center; max-width: 400px; padding: 20px; }
  h1 { font-size: 48px; color: #e03131; margin-bottom: 10px; }
  h3 { font-size: 18px; margin-bottom: 15px; }
  p { font-size: 14px; color: #666; line-height: 1.5; margin-bottom: 20px; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #999; }
</style>
</head>
<body>
  <div class="container">
    <h1>504</h1>
    <h3>Gateway Timeout</h3>
    <p>The upstream server failed to fulfill the request within the specified time allotment socket window.</p>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const antivirusThreatHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Antivirus Threat</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #1a1a1a; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 440px; background: #1b382b; border: 1px solid #2ea043; border-radius: 8px; padding: 24px; }
  .title { font-size: 16px; font-weight: 600; color: #3fb950; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #d2a8ff; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #2ea043; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #888; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🛡️ Threat Quarantined Successfully</div>
    <p class="desc">Windows Defender detected and blocked Trojan:Win32/Wacatac.H!ml from executing payload.</p>
    <button class="btn" onclick="location.reload()">View Details</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const dbConnectionLostHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>DB Connection Lost</title>
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
    <div class="title">🗄️ Database Connection Refused</div>
    <p class="desc">SQLSTATE[HY000] [2002] Connection refused. Local service socket pool exhausted or offline.</p>
    <button class="btn" onclick="location.reload()">Reconnect</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`
