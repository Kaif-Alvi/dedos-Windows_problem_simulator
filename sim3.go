package main

// sims3.go - Contains HTML strings for new simulations (3 to 10)

// 3. Update Stuck at 99% Simulation
const updateStuckHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Windows Update Stuck</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #0078d7;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .spinner {
    width: 50px;
    height: 50px;
    border: 5px solid rgba(255,255,255,0.3);
    border-top: 5px solid #fff;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin-bottom: 30px;
  }
  @keyframes spin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
  h1 { font-size: 24px; font-weight: 400; margin-bottom: 10px; }
  p { font-size: 15px; color: rgba(255,255,255,0.8); }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: rgba(255,255,255,0.5); }
</style>
</head>
<body>
  <div class="spinner"></div>
  <h1>Working on updates 99%</h1>
  <p>Don't turn off your PC. This will take a while.</p>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 4. Activation Expiry Warning Simulation
const activationExpiryHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Activation Expiry Warning</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #121619;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .dialog {
    width: 450px;
    background: #1e252b;
    border: 1px solid #323d47;
    border-radius: 8px;
    padding: 28px;
    box-shadow: 0 15px 40px rgba(0,0,0,0.5);
    text-align: center;
  }
  .icon { font-size: 42px; color: #ffbc42; margin-bottom: 14px; }
  .title { font-size: 17px; font-weight: 600; margin-bottom: 10px; }
  .desc { font-size: 13.5px; color: #a0aec0; line-height: 1.6; margin-bottom: 24px; }
  .btn {
    padding: 10px 22px;
    font-size: 13.5px;
    background: #0078d7;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .btn:hover { background: #1084e3; }
  .watermark {
    position: fixed;
    bottom: 40px;
    right: 40px;
    color: rgba(255,255,255,0.15);
    font-size: 14px;
    font-family: monospace;
    pointer-events: none;
    text-align: right;
    line-height: 1.4;
  }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="icon">⚡</div>
    <div class="title">Your Windows license will expire soon</div>
    <div class="desc">You need to activate Windows in Settings before your current license period officially ends.</div>
    <button class="btn" onclick="alert('Redirecting to Windows Activation Settings...')">Go to Settings</button>
  </div>
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

// 5. Printer Offline Simulation
const printerOfflineHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Printer Offline</title>
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
    overflow: hidden;
  }
  .dialog {
    width: 440px;
    background: #2d2d2d;
    border: 1px solid #3f3f3f;
    border-radius: 8px;
    padding: 24px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.6);
  }
  .title { font-size: 16px; font-weight: 600; margin-bottom: 12px; color: #ff9800; }
  .desc { font-size: 13.5px; color: #ccc; line-height: 1.6; margin-bottom: 20px; }
  .printer-box {
    background: #252525;
    border: 1px solid #444;
    border-radius: 6px;
    padding: 12px;
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 20px;
  }
  .p-icon { font-size: 28px; color: #9cdcfe; }
  .p-name { font-size: 14px; font-weight: 600; }
  .p-status { font-size: 12px; color: #ce9178; }
  .btn-row { display: flex; justify-content: flex-end; gap: 10px; }
  .btn {
    padding: 8px 18px;
    font-size: 13px;
    background: #0078d7;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .btn:hover { background: #1084e3; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🖨️ Printer is Offline</div>
    <div class="desc">Windows cannot communicate with the printer. Check if the printer is powered on and connected to the network or USB.</div>
    <div class="printer-box">
      <div class="p-icon">⎙</div>
      <div>
        <div class="p-name">HP LaserJet Pro M15w</div>
        <div class="p-status">Status: Offline (3 documents pending)</div>
      </div>
    </div>
    <div class="btn-row">
      <button class="btn" onclick="location.reload()">Open Print Queue</button>
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

// 6. Audio Service Simulation
const audioServiceHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Audio Service Error</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #1a1528;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .dialog {
    width: 440px;
    background: #262035;
    border: 1px solid #3d3254;
    border-radius: 8px;
    padding: 24px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.6);
    text-align: center;
  }
  .icon { font-size: 48px; color: #db8cff; margin-bottom: 12px; }
  .title { font-size: 16px; font-weight: 600; margin-bottom: 8px; }
  .desc { font-size: 13.5px; color: #bca8d3; line-height: 1.6; margin-bottom: 20px; }
  .btn {
    padding: 8px 18px;
    font-size: 13px;
    background: #8b5cf6;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .btn:hover { background: #7c3aed; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="icon">🔇</div>
    <div class="title">No Audio Output Device is Installed</div>
    <div class="desc">Windows cannot find your speakers, headphones, or audio hardware drivers. Audio services have stopped responding.</div>
    <button class="btn" onclick="location.reload()">Troubleshoot Sound</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 7. SSD Warning Simulation
const ssdWarningHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>SSD Health Warning</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #221515;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .dialog {
    width: 460px;
    background: #331f1f;
    border: 1px solid #553333;
    border-radius: 8px;
    padding: 24px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.6);
  }
  .title { font-size: 16px; font-weight: 600; margin-bottom: 12px; color: #ff6b6b; display: flex; align-items: center; gap: 8px; }
  .desc { font-size: 13.5px; color: #dcdcdc; line-height: 1.6; margin-bottom: 20px; }
  .health-box { background: #261515; border: 1px solid #442222; padding: 12px; border-radius: 6px; font-size: 13px; margin-bottom: 20px; color: #ff9999; }
  .btn { padding: 8px 18px; font-size: 13px; background: #d9534f; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .btn:hover { background: #c9302c; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #777; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">⚠️ Hardware Health Critical Warning</div>
    <div class="desc">S.M.A.R.T. error detected on your primary storage drive. Immediate backup of personal files is strongly recommended.</div>
    <div class="health-box">
      Drive: C: (NVMe SSD)<br>
      Predicted Failure Risk: HIGH (Health: 12%)
    </div>
    <button class="btn" onclick="location.reload()">Back Up Now</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 8. App Hang Simulation
const appHangHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>App Not Responding</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: rgba(0,0,0,0.6);
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #000;
  }
  .dialog {
    width: 420px;
    background: #f0f0f0;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-shadow: 0 5px 20px rgba(0,0,0,0.4);
    padding: 20px;
  }
  .title { font-size: 15px; font-weight: 600; margin-bottom: 10px; color: #333; }
  .desc { font-size: 13px; color: #555; line-height: 1.5; margin-bottom: 20px; }
  .btn-row { display: flex; justify-content: flex-end; gap: 8px; }
  .btn { padding: 6px 16px; font-size: 12.5px; background: #e1e1e1; border: 1px solid #adadad; border-radius: 2px; cursor: pointer; }
  .btn:hover { background: #e5f1fb; border-color: #0078d7; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #aaa; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">Google Chrome is not responding</div>
    <div class="desc">Windows can check online for a solution and try to restart the program, or you can close the program right now.</div>
    <div class="btn-row">
      <button class="btn" onclick="location.reload()">Close the program</button>
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

// 9. GPU TDR Crash Simulation
const gpuVectorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Display Driver Recovered</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #111;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #00ffcc;
    font-size: 22px;
    overflow: hidden;
  }
  .box { text-align: center; }
  .sub { font-size: 14px; color: #888; margin-top: 10px; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #555; }
</style>
</head>
<body>
  <div class="box">
    <div>Display driver stopped responding and has recovered</div>
    <div class="sub">NVIDIA Windows Kernel Mode Driver (WDDM) recovered successfully.</div>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`

// 10. Defender Threat Alert Simulation
const defenderAlertHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Security Alert</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background: #0f172a;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Segoe UI', sans-serif;
    color: #fff;
    overflow: hidden;
  }
  .dialog {
    width: 440px;
    background: #1e293b;
    border: 1px solid #334155;
    border-radius: 8px;
    padding: 24px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.6);
  }
  .title { font-size: 16px; font-weight: 600; margin-bottom: 10px; color: #ef4444; }
  .desc { font-size: 13.5px; color: #cbd5e1; line-height: 1.6; margin-bottom: 20px; }
  .threat { background: #0f172a; padding: 10px; border-radius: 6px; font-family: monospace; font-size: 12px; color: #f87171; margin-bottom: 20px; }
  .btn { padding: 8px 18px; font-size: 13px; background: #ef4444; color: #fff; border: none; border-radius: 4px; cursor: pointer; }
  .btn:hover { background: #dc2626; }
  .warn { position: fixed; bottom: 18px; font-size: 12px; color: #64748b; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">🛡️ Threat Detected!</div>
    <div class="desc">Microsoft Defender Antivirus has found active malware running in your system memory.</div>
    <div class="threat">Trojan:Win32/Wacatac.H!ml<br>Status: Severe Active Threat</div>
    <button class="btn" onclick="location.reload()">Take Actions</button>
  </div>
  <div class="warn">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`