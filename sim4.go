package main

// sims4.go - Contains HTML strings for simulations 11 to 20

// 11. Critical Process Died Simulation
const criticalProcessHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Critical Process Died</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #0078d7;
    height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 80px;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .emoji { font-size: 80px; margin-bottom: 20px; }
  h1 { font-size: 28px; font-weight: 300; line-height: 1.4; margin-bottom: 30px; }
  .info { display: flex; gap: 30px; align-items: center; }
  .qr { width: 120px; height: 120px; background: #fff; padding: 5px; }
  .details { font-size: 15px; line-height: 1.6; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: rgba(255,255,255,0.5); }
</style>
</head>
<body>
  <div class="emoji">:(</div>
  <h1>Your PC ran into a problem and needs to restart. We're just collecting some error info, and then we'll restart for you.</h1>
  <div class="info">
    <div class="qr" style="background:#fff; display:flex; align-items:center; justify-content:center; color:#000; font-size:11px; text-align:center;">[ QR CODE ]</div>
    <div class="details">
      100% complete<br><br>
      What failed: CRITICAL_PROCESS_DIED
    </div>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 12. Boot Loop Simulation
const bootLoopHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Boot Loop</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #000;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .logo { font-size: 48px; color: #0078d7; margin-bottom: 40px; animation: pulse 1.5s infinite; }
  @keyframes pulse { 0% { opacity: 0.3; } 50% { opacity: 1; } 100% { opacity: 0.3; } }
  .spinner {
    width: 40px;
    height: 40px;
    border: 4px solid rgba(255,255,255,0.1);
    border-top: 4px solid #fff;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }
  @keyframes spin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #444; }
</style>
</head>
<body>
  <div class="logo">🪟</div>
  <div class="spinner"></div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script.dedosResume = function() {};
</script>
</body>
</html>>
  window.dedosPause = function() {};
  window
`

// 13. Out of Memory Simulation
const outOfMemoryHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Out of Memory</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #202020;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
  }
  .dialog {
    width: 420px;
    background: #2d2d2d;
    border: 1px solid #3f3f3f;
    border-radius: 6px;
    padding: 24px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.5);
  }
  .title { font-size: 16px; font-weight: 600; color: #f48771; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #cccccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">⚠️ Out of Memory Error</div>
    <div class="desc">Your system has run out of available RAM and virtual memory. Close some applications immediately to prevent data loss.</div>
    <button class="btn" onclick="location.reload()">Close Apps</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 14. Explorer Crash Simulation
const explorerCrashHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Explorer Crash</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #008080;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .msg { font-size: 20px; background: rgba(0,0,0,0.4); padding: 20px; border-radius: 8px; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #ddd; }
</style>
</head>
<body>
  <div class="msg">Windows Explorer has stopped working (Taskbar and Desktop hidden)</div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 15. No Internet Simulation
const noInternetHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>No Internet</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #1e1e1e;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
  }
  .dialog {
    width: 440px;
    background: #252526;
    border: 1px solid #333;
    border-radius: 8px;
    padding: 24px;
  }
  .title { font-size: 16px; font-weight: 600; color: #cca700; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #bbb; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🌐 No Internet Connection</div>
    <div class="desc">Ethernet / Wi-Fi network cable is unplugged or the router is offline. Limited or no connectivity.</div>
    <button class="btn" onclick="location.reload()">Diagnose Network</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 16. Overheating Thermal Warning
const overheatingHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Overheating Warning</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #2c1212;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
  }
  .dialog {
    width: 440px;
    background: #3c1a1a;
    border: 1px solid #662222;
    border-radius: 8px;
    padding: 24px;
  }
  .title { font-size: 16px; font-weight: 600; color: #ff5555; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ffcccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #cc3333; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #885555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔥 CPU Overheating Alert</div>
    <div class="desc">CPU temperature reached 98°C. Thermal throttling active. System will shut down automatically to prevent hardware damage.</div>
    <button class="btn" onclick="location.reload()">Shut Down Now</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 17. Automatic Repair Loop
const autoRepairHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Preparing Automatic Repair</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #000;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  h1 { font-size: 22px; font-weight: 300; margin-bottom: 30px; }
  .spinner {
    width: 50px;
    height: 50px;
    border: 5px solid rgba(255,255,255,0.2);
    border-top: 5px solid #fff;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }
  @keyframes spin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #555; }
</style>
</head>
<body>
  <h1>Preparing Automatic Repair</h1>
  <div class="spinner"></div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 18. Permanent Watermark Overlay
const watermarkHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Activation Watermark</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #111;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .watermark {
    position: fixed;
    bottom: 30px;
    right: 30px;
    color: rgba(255,255,255,0.25);
    font-size: 16px;
    font-family: monospace;
    pointer-events: none;
    text-align: right;
    line-height: 1.4;
  }
  .center-text { font-size: 24px; color: #888; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #555; }
</style>
</head>
<body>
  <div class="center-text">Windows Unactivated Build 26100</div>
  <div class="watermark">
    Activate Windows<br>
    Go to Settings to activate Windows.
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 19. Registry Corruption Error
const registryErrorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Registry Error</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #1a1a1a;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
  }
  .dialog {
    width: 440px;
    background: #2a2a2a;
    border: 1px solid #444;
    border-radius: 8px;
    padding: 24px;
  }
  .title { font-size: 16px; font-weight: 600; color: #ff9900; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">⚙️ Registry Hive Corruption</div>
    <div class="desc">System registry files are missing or corrupted. Windows could not load SYSTEM hive configuration safely.</div>
    <button class="btn" onclick="location.reload()">Restore Registry</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 20. USB Device Malfunctioning
const usbMalfunctionHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>USB Malfunction</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #1e1e1e;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
  }
  .dialog {
    width: 440px;
    background: #2d2d2d;
    border: 1px solid #3f3f3f;
    border-radius: 8px;
    padding: 24px;
  }
  .title { font-size: 16px; font-weight: 600; color: #f14c4c; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .btn { padding: 8px 18px; background: #0078d7; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🔌 USB Device Not Recognized</div>
    <div class="desc">The last USB device you connected to this computer malfunctioned, and Windows does not recognize it.</div>
    <button class="btn" onclick="location.reload()">Close</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`