package main

// sims7.go - Contains HTML strings for simulations 51 to 60

const ramParityErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>RAM Parity Error</title>
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
    <div class="title">⚠️ RAM Parity Check Failure</div>
    <div class="desc">Hardware memory address exception encountered. Data corruption detected in physical memory modules.</div>
    <button class="btn" onclick="location.reload()">Diagnostic</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const smartFailureHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>SMART Failure</title>
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
    <div class="title">💽 S.M.A.R.T. Hard Drive Warning</div>
    <div class="desc">A hard disk issue has been detected. Backup your files immediately to prevent permanent data loss.</div>
    <button class="btn" onclick="location.reload()">Backup Now</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const powerSurgeHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Power Surge</title>
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
    <div class="title">⚡ ASUS Anti-Surge Triggered</div>
    <div class="desc">Power Surge was detected on the previous power supply unit. System was unstable to protect mainboard components.</div>
    <button class="btn" onclick="location.reload()">Power Off</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const biosFlashFailHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>BIOS Flash Failed</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #000; height: 100vh; display: flex; flex-direction: column; justify-content: center; padding: 60px; font-family: monospace; color: #00ff00; }
  h2 { margin-bottom: 20px; font-size: 20px; }
  p { font-size: 15px; line-height: 1.5; color: #ccc; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #555; }
</style>
</head>
<body>
  <h2>EEPROM Flash Block Verification Error</h2>
  <p>BIOS image checksum mismatch! Do not turn off your system or remove the update media, or your motherboard will be bricked.</p>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const opticalErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Optical Drive Error</title>
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
    <div class="title">💿 CD/DVD Read Error</div>
    <p class="desc">The disc in drive D: is unreadable, scratched, or uses an unsupported file system format (UDF/ISO).</p>
    <button class="btn" onclick="location.reload()">Eject</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const spoolerDeadlockHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Spooler Deadlock</title>
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
    <div class="title">🖨️ Print Spooler Queue Stuck</div>
    <p class="desc">Multiple jobs are blocked in the local print queue directory. Spooler service deadlock occurred.</p>
    <button class="btn" onclick="location.reload()">Clear Queue</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const gatewayUnreachableHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Gateway Unreachable</title>
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
    <div class="title">🌐 Default Gateway Unreachable</div>
    <p class="desc">Packets dropped between host and router interface. Ping request could not find host 192.168.1.1.</p>
    <button class="btn" onclick="location.reload()">Troubleshoot</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const vpnDroppedHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>VPN Dropped</title>
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
    <div class="title">🔒 VPN Connection Lost</div>
    <p class="desc">Secure tunnel encryption interrupted. Kill switch enabled to block all unsecured network traffic.</p>
    <button class="btn" onclick="location.reload()">Reconnect</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const cloudsyncConflictHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Cloud Sync Conflict</title>
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
    <div class="title">☁️ OneDrive Sync Conflict</div>
    <p class="desc">A file version mismatch was found between local storage and cloud storage. Manual merge required.</p>
    <button class="btn" onclick="location.reload()">Resolve</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const zerodayWarningHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Zero-Day Warning</title>
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
    <div class="title">🚨 Zero-Day Vulnerability Alert</div>
    <p class="desc">An unpatched remote code execution vulnerability is targeting system binaries. Isolate device immediately.</p>
    <button class="btn" onclick="location.reload()">Isolate</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`