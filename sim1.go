package main

// ================= FAKE UPDATE LOOP SIMULATION =================
const fakeUpdateHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background:#000;
    height:100vh;
    display:flex;
    flex-direction:column;
    align-items:center;
    justify-content:center;
    font-family:'Segoe UI', sans-serif;
    color:#fff;
    overflow:hidden;
  }
  .spinner {
    width:60px; height:60px;
    border:4px solid #222;
    border-top:4px solid #0078d7;
    border-radius:50%;
    animation:spin 1s linear infinite;
    margin-bottom:30px;
  }
  @keyframes spin { 100% { transform:rotate(360deg); } }
  .msg { font-size:18px; margin-bottom:8px; }
  .percent { font-size:14px; color:#aaa; }
  .warn { position:fixed; bottom:20px; font-size:12px; color:#555; }
</style>
</head>
<body>
  <div class="spinner" id="spinner"></div>
  <div class="msg">Working on updates <span id="pct">14</span>% complete</div>
  <div class="percent">Don't turn off your PC</div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  // The percentage rises and falls randomly without reaching 100%, simulating a stuck update.
  let pct = 14;
  let running = true;
  const pctEl = document.getElementById('pct');

  function tick() {
    if (!running) { setTimeout(tick, 200); return; }
    if (pct < 87) {
      pct += Math.random() * 0.6;
    } else {
      pct = 61; // Fall back after approaching 87% to create the illusion of being stuck.
    }
    pctEl.innerText = pct.toFixed(0);
    setTimeout(tick, 300);
  }
  tick();

  window.dedosPause = function() { running = false; document.getElementById('spinner').style.animationPlayState='paused'; };
  window.dedosResume = function() { running = true; document.getElementById('spinner').style.animationPlayState='running'; };
</script>
</body>
</html>
`

// ================= RANSOMWARE SCREEN SIMULATION (visual mock only; touches no files) =================
const ransomwareHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background:#111;
    height:100vh;
    display:flex;
    flex-direction:column;
    align-items:center;
    justify-content:center;
    font-family:'Consolas', monospace;
    color:#ddd;
    overflow:hidden;
  }
  .box {
    width:520px;
    border:2px solid #cc2222;
    background:#1a0d0d;
    padding:30px 36px;
    text-align:center;
  }
  .lock { font-size:50px; margin-bottom:14px; }
  .title { color:#ff4444; font-size:22px; font-weight:bold; margin-bottom:14px; letter-spacing:1px; }
  .desc { font-size:13px; line-height:1.8; color:#ccc; margin-bottom:20px; }
  .timer { font-size:26px; color:#ff8800; margin-bottom:10px; letter-spacing:2px; }
  .sim-tag { font-size:12px; color:#8fbfff; margin-top:16px; }
  .warn { position:fixed; bottom:18px; font-size:12px; color:#777; }
</style>
</head>
<body>
  <div class="box">
    <div class="lock">🔒</div>
    <div class="title">YOUR FILES ARE LOCKED</div>
    <div class="desc">
      This is an educational simulation - no real files were touched or encrypted.<br>
      Real ransomware uses screens like this to frighten users.
    </div>
    <div class="timer" id="timer">23:59:59</div>
    <div class="sim-tag">⚠ SIMULATION MODE — NO FILES AFFECTED</div>
  </div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  // Visual countdown effect only; no real action is performed.
  let seconds = 24 * 3600 - 1;
  let running = true;
  const timerEl = document.getElementById('timer');

  function format(s) {
    const h = String(Math.floor(s/3600)).padStart(2,'0');
    const m = String(Math.floor((s%3600)/60)).padStart(2,'0');
    const sec = String(s%60).padStart(2,'0');
    return h+':'+m+':'+sec;
  }

  function tick() {
    if (!running) { setTimeout(tick, 200); return; }
    if (seconds > 0) seconds--;
    timerEl.innerText = format(seconds);
    setTimeout(tick, 1000);
  }
  tick();

  window.dedosPause = function() { running = false; };
  window.dedosResume = function() { running = true; };
</script>
</body>
</html>
`

// ================= DRIVER FAILURE SIMULATION (BSOD variant) =================
const driverFailureHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  html, body { width:100%; height:100%; background:#0078d7; font-family:'Segoe UI', sans-serif; color:#fff; overflow:hidden; }
  .wrap { height:100vh; display:flex; flex-direction:column; justify-content:center; padding:80px 100px; }
  .face { font-size:110px; margin-bottom:20px; line-height:1; }
  .msg1 { font-size:26px; max-width:750px; line-height:1.5; margin-bottom:6px; }
  .msg2 { font-size:16px; max-width:700px; line-height:1.6; color:#e6f0fb; margin-bottom:40px; }
  .progress { font-size:16px; margin-bottom:60px; }
  .stop-info { font-size:13px; line-height:1.8; color:#dbeafe; }
  .stop-code { color:#fff; font-weight:600; }
  .footer { position:fixed; bottom:18px; left:0; width:100%; text-align:center; font-size:12px; color:#cfe4fb; }
</style>
</head>
<body>
  <div class="wrap">
    <div class="face">:(</div>
    <div class="msg1">Your PC ran into a problem and needs to restart.</div>
    <div class="msg2">A driver caused an invalid memory access and could not be handled by the system.</div>
    <div class="progress" id="progressText">0% complete</div>
    <div class="stop-info">
      For more information about this issue and possible fixes, visit<br>
      https://www.windows.com/stopcode<br><br>
      <span class="stop-code">STOP CODE: DRIVER_IRQL_NOT_LESS_OR_EQUAL</span>
    </div>
  </div>
  <div class="footer">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  let percent = 0;
  let running = true;
  const progressEl = document.getElementById('progressText');
  function tick() {
    if (!running) { setTimeout(tick, 150); return; }
    if (percent < 100) {
      percent += 1;
      progressEl.innerText = percent + '% complete';
      setTimeout(tick, 55);
    } else {
      percent = 0;
      setTimeout(tick, 1500);
    }
  }
  tick();
  window.dedosPause = function() { running = false; };
  window.dedosResume = function() { running = true; };
</script>
</body>
</html>
`

// ================= DISK BOOT FAILURE SIMULATION =================
const diskBootFailureHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background:#000;
    height:100vh;
    font-family:'Consolas', monospace;
    color:#c0c0c0;
    padding:40px;
    font-size:15px;
    line-height:1.8;
    overflow:hidden;
  }
  .cursor-inline {
    display:inline-block;
    width:9px; height:15px;
    background:#c0c0c0;
    animation:blink 1s infinite;
    vertical-align:middle;
  }
  @keyframes blink { 50% { opacity:0; } }
  .warn { position:fixed; bottom:15px; left:40px; color:#555; font-size:12px; }
</style>
</head>
<body>
  <div>PhoenixBIOS(TM) 4.0 Release 6.0</div>
  <div>Copyright 1985-2026 Phoenix Technologies Ltd.</div>
  <br>
  <div id="lines"></div>
  <br>
  <div>Reboot and Select proper Boot device</div>
  <div>or Insert Boot Media in selected Boot device<span class="cursor-inline"></span></div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  const lines = [
    'Detecting IDE drives... FAILED',
    'Primary Master: NOT DETECTED',
    'Primary Slave: NOT DETECTED',
    'NO BOOTABLE DEVICE -- insert boot disk and press any key'
  ];
  let idx = 0, paused = false;
  const el = document.getElementById('lines');

  function showNext() {
    if (paused) { setTimeout(showNext, 200); return; }
    if (idx >= lines.length) return;
    el.innerHTML += lines[idx] + '<br>';
    idx++;
    setTimeout(showNext, 500);
  }
  showNext();

  window.dedosPause = function() { paused = true; };
  window.dedosResume = function() { paused = false; };
</script>
</body>
</html>
`

// ================= LOW DISK SPACE SIMULATION =================
const lowDiskSpaceHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background:#0d0f14;
    height:100vh;
    display:flex;
    align-items:center;
    justify-content:center;
    font-family:'Segoe UI', sans-serif;
    color:#eaeaea;
    overflow:hidden;
  }
  .dialog {
    width:420px;
    background:#1e1e1e;
    border:1px solid #3a3a3a;
    border-radius:8px;
    padding:24px;
    box-shadow:0 20px 60px rgba(0,0,0,0.5);
  }
  .dtitle { font-size:15px; font-weight:600; margin-bottom:10px; display:flex; align-items:center; gap:8px; }
  .ddesc { font-size:13px; color:#c3c9d6; margin-bottom:18px; line-height:1.6; }
  .bar-bg { background:#333; border-radius:6px; height:16px; overflow:hidden; margin-bottom:8px; }
  .bar { height:100%; background:linear-gradient(90deg,#ff3333,#ff8800); width:5%; transition:width 2s ease; }
  .stat { font-size:12px; color:#999; display:flex; justify-content:space-between; }
  .warn { position:fixed; bottom:18px; font-size:12px; color:#555; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="dtitle">⚠️ Low disk space</div>
    <div class="ddesc">You are running very low on disk space on Local Disk (C:). Free up some space by deleting old files.</div>
    <div class="bar-bg"><div class="bar" id="bar"></div></div>
    <div class="stat"><span id="freeSpace">1.2 GB free</span><span>256 GB total</span></div>
  </div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  let running = true;
  let used = 5;
  const bar = document.getElementById('bar');
  const freeText = document.getElementById('freeSpace');

  function tick() {
    if (!running) { setTimeout(tick, 300); return; }
    used = Math.min(98, used + Math.random() * 3);
    bar.style.width = used + '%';
    const freeGB = (256 * (100 - used) / 100).toFixed(1);
    freeText.innerText = freeGB + ' GB free';
    setTimeout(tick, 1500);
  }
  tick();

  window.dedosPause = function() { running = false; };
  window.dedosResume = function() { running = true; };
</script>
</body>
</html>
`

// ================= NETWORK DISCONNECTED SIMULATION =================
const networkDisconnectedHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background:#f5f5f5;
    height:100vh;
    display:flex;
    flex-direction:column;
    align-items:center;
    justify-content:center;
    font-family:'Segoe UI', sans-serif;
    color:#333;
    overflow:hidden;
  }
  .icon-wrap { position:relative; margin-bottom:24px; }
  .wifi-icon { font-size:70px; color:#888; }
  .cross { position:absolute; top:-6px; right:-10px; font-size:34px; color:#e02020; font-weight:bold; }
  .title { font-size:20px; font-weight:600; margin-bottom:6px; }
  .desc { font-size:13px; color:#666; }
  .warn { position:fixed; bottom:18px; font-size:12px; color:#999; }
</style>
</head>
<body>
  <div class="icon-wrap">
    <div class="wifi-icon">📶</div>
    <div class="cross">✕</div>
  </div>
  <div class="title">No Internet Connection</div>
  <div class="desc">You're not connected to any network right now.</div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  // This simulation is static, so keep Pause/Resume functions empty to satisfy Go bindings.
  window.dedosPause = function() {};
  window.dedosResume = function() {};
</script>
</body>
</html>
`
