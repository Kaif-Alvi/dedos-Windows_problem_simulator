package main

// sims8.go - Contains HTML strings for simulations 61 to 70

const cpuThrottleHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>CPU Thermal Throttling</title>
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
    <div class="title">🔥 CPU Clock Speed Throttled</div>
    <p class="desc">Processor speed reduced to 0.79 GHz due to sustained high temperature and inadequate cooling performance.</p>
    <button class="btn" onclick="location.reload()">Cool Down</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const pagefileCorruptHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Pagefile Corrupt</title>
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
    <div class="title">⚠️ Virtual Memory Pagefile Corrupted</div>
    <p class="desc">Windows created a temporary paging file because of a configuration problem upon startup. Data integrity risk.</p>
    <button class="btn" onclick="location.reload()">Recreate</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const masterBootRecordHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>MBR Damaged</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #000; height: 100vh; display: flex; flex-direction: column; justify-content: center; padding: 60px; font-family: monospace; color: #fff; }
  h2 { margin-bottom: 20px; font-size: 22px; color: #ff5555; }
  p { font-size: 15px; line-height: 1.5; color: #aaa; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #555; }
</style>
</head>
<body>
  <h2>Master Boot Record Error</h2>
  <p>Invalid partition table or missing operating system signature on sector 0. Press any key to reboot.</p>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const secureBootViolationHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Secure Boot Violation</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #900; height: 100vh; display: flex; align-items: center; justify-content: center; font-family: 'Segoe UI', sans-serif; color: #fff; }
  .dialog { width: 480px; background: #400; border: 2px solid #f00; border-radius: 8px; padding: 24px; }
  .title { font-size: 18px; font-weight: 600; color: #ffcccc; margin-bottom: 12px; }
  .desc { font-size: 14px; color: #ff9999; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #fff; color: #900; font-weight: bold; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #ff9999; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔒 Secure Boot Violation</div>
    <p class="desc">System unauthorized image detected. The Secure Boot key validation failed for bootmgr signature database.</p>
    <button class="btn" onclick="location.reload()">Enter BIOS</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const tpmLockoutHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>TPM Lockout</title>
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
    <div class="title">🛡️ TPM 2.0 Device Lockout</div>
    <p class="desc">Trusted Platform Module encountered too many authorization failures. BitLocker recovery key entry required.</p>
    <button class="btn" onclick="location.reload()">Enter Recovery Key</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const bitlockerLockHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>BitLocker Recovery</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background: #004578; height: 100vh; display: flex; flex-direction: column; justify-content: center; padding: 80px; font-family: 'Segoe UI', sans-serif; color: #fff; }
  h1 { font-size: 26px; font-weight: 300; margin-bottom: 20px; }
  p { font-size: 15px; line-height: 1.6; max-width: 600px; margin-bottom: 30px; }
  .btn { padding: 10px 20px; background: #fff; color: #004578; font-weight: bold; border: none; border-radius: 4px; cursor: pointer; width: max-content; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: rgba(255,255,255,0.6); }
</style>
</head>
<body>
  <h1>BitLocker Drive Encryption Recovery</h1>
    <p>Your drive is locked because of security policy changes. You need to enter your 48-digit recovery key to unlock this drive.</p>
    <button class="btn" onclick="location.reload()">Enter Recovery Password</button>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const ghostTouchHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Ghost Touch Error</title>
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
    <div class="title">📱 Touch Screen Calibration Fault</div>
    <p class="desc">Ghost touch inputs detected on hardware digitizer interface. Multi-touch driver input stream jammed.</p>
    <button class="btn" onclick="location.reload()">Recalibrate</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const ambientLightSensorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Sensor Fault</title>
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
    <div class="title">💡 Ambient Light Sensor Failure</div>
    <p class="desc">Automatic brightness control service cannot poll sensor data stream from I2C bus device.</p>
    <button class="btn" onclick="location.reload()">Reset Sensor</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const fingerprintHardwareHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Biometric Failure</title>
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
    <div class="title">👆 Fingerprint Scanner Unresponsive</div>
    <p class="desc">Windows Hello biometric sensor hardware initialization timed out. Sensor unit needs power cycle.</p>
    <button class="btn" onclick="location.reload()">Re-initialize</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`

const nvmeControllerHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>NVMe Controller Error</title>
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
    <div class="title">⚡ NVMe Express Controller Reset</div>
    <p class="desc">PCIe storage controller register status error. Storage command queue aborted unexpectedly.</p>
    <button class="btn" onclick="location.reload()">Restart Controller</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`