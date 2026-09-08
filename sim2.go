package main

// These three simulations are kept as Go string constants instead of separate
// asset files, following the Sparrow project's style.

// ================= BLACK SCREEN SIMULATION =================
const blackScreenHTML = `
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
    color:#111;
    overflow:hidden;
  }
  .cursor {
    width:2px; height:20px; background:#333;
    animation: blink 1s infinite;
  }
  @keyframes blink { 50% { opacity:0; } }
  .warn {
    position:fixed; bottom:20px; left:20px;
    color:#222; font-size:12px; font-family:monospace;
  }
</style>
</head>
<body>
  <div class="cursor"></div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>
</body>
</html>
`

// ================= CORRUPT BIOS SIMULATION =================
const corruptBiosHTML = `
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
    padding:30px;
    font-size:15px;
    line-height:1.6;
    overflow:hidden;
  }
  .red { color:#ff4444; font-weight:bold; }
  .yellow { color:#ffcc00; }
  .green { color:#33ff33; }
  .warn {
    position:fixed; bottom:15px; left:20px;
    color:#666; font-size:12px;
  }
</style>
</head>
<body>
  <div class="red">CMOS CHECKSUM ERROR - DEFAULTS LOADED</div>
  <br>
  <div>BIOS ROM CHECKSUM FAILED</div>
  <div class="yellow">System halted. Verify integrity of BIOS image.</div>
  <br>
  <div>Primary Boot Device: <span class="red">NOT FOUND</span></div>
  <div>Secondary Boot Device: <span class="red">NOT FOUND</span></div>
  <br>
  <div class="green">Press F1 to enter SETUP</div>
  <div class="green">Press F2 to load defaults and continue</div>
  <br><br>
  <div id="typing"></div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  // Show error codes with a typing effect.
  const lines = [
    "AMIBIOS(C) 2026 American Megatrends, Inc.",
    "Checking NVRAM...FAILED",
    "Initializing USB Controllers... ERROR",
    "Fatal Error: NO BOOTABLE DEVICE FOUND"
  ];
  let idx = 0, char = 0, paused = false;
  const el = document.getElementById('typing');

  function typeLoop() {
    if (paused) { setTimeout(typeLoop, 200); return; }
    if (idx >= lines.length) return;
    const line = lines[idx];
    if (char <= line.length) {
      el.innerText = lines.slice(0, idx).join('\\n') + '\\n' + line.slice(0, char);
      char++;
      setTimeout(typeLoop, 40);
    } else {
      idx++; char = 0;
      setTimeout(typeLoop, 300);
    }
  }
  typeLoop();

  // Global functions called by Go to pause and resume.
  window.dedosPause = function() { paused = true; };
  window.dedosResume = function() { paused = false; };
</script>
</body>
</html>
`

// ================= CPU OVERLOAD SIMULATION =================
const cpuOverloadHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body {
    background:#0a0a0a;
    height:100vh;
    display:flex;
    flex-direction:column;
    align-items:center;
    justify-content:center;
    font-family:'Segoe UI', sans-serif;
    color:#fff;
    overflow:hidden;
  }
  .title { font-size:22px; margin-bottom:30px; color:#ff5555; }
  .core-grid {
    display:grid;
    grid-template-columns:repeat(4, 1fr);
    gap:12px;
    width:420px;
  }
  .core {
    background:#1a1a1a;
    border-radius:6px;
    padding:10px;
    text-align:center;
    border:1px solid #333;
  }
  .core-label { font-size:11px; color:#888; margin-bottom:6px; }
  .core-bar-bg { background:#222; border-radius:4px; height:10px; overflow:hidden; }
  .core-bar { height:100%; background:linear-gradient(90deg,#ff3333,#ff9900); width:0%; transition:width .3s; }
  .core-val { font-size:13px; margin-top:6px; color:#ff8888; }
  .warn { position:fixed; bottom:20px; font-size:12px; color:#555; }
  .temp { margin-top:25px; font-size:14px; color:#ffaa00; }
</style>
</head>
<body>
  <div class="title">⚠ CPU USAGE CRITICAL</div>
  <div class="core-grid" id="grid"></div>
  <div class="temp" id="temp">CPU TEMP: 45°C</div>
  <div class="warn">This is a simulation. Press Ctrl + Alt + Win + C to exit.</div>

<script>
  const grid = document.getElementById('grid');
  const cores = [];
  for (let i = 0; i < 8; i++) {
    const div = document.createElement('div');
    div.className = 'core';
    div.innerHTML = '<div class="core-label">CORE ' + i + '</div>' +
      '<div class="core-bar-bg"><div class="core-bar" id="bar' + i + '"></div></div>' +
      '<div class="core-val" id="val' + i + '">0%</div>';
    grid.appendChild(div);
    cores.push(i);
  }

  let running = true;
  let temp = 45;

  function tick() {
    if (!running) { setTimeout(tick, 200); return; }
    cores.forEach(i => {
      const usage = Math.min(100, 60 + Math.random() * 40).toFixed(0);
      document.getElementById('bar' + i).style.width = usage + '%';
      document.getElementById('val' + i).innerText = usage + '%';
    });
    temp = Math.min(99, temp + Math.random() * 2);
    document.getElementById('temp').innerText = 'CPU TEMP: ' + temp.toFixed(0) + '°C';
    setTimeout(tick, 400);
  }
  tick();

  // Global functions called by Go to pause and resume.
  window.dedosPause = function() { running = false; };
  window.dedosResume = function() { running = true; };
</script>
</body>
</html>
`
