package main

import "html"

func problemSimulationHTML(title, message string) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>` + html.EscapeString(title) + `</title>
<style>
  * { margin:0; padding:0; box-sizing:border-box; }
  body { background:#1e1e1e; color:#fff; height:100vh; display:flex; align-items:center; justify-content:center; font-family:'Segoe UI',sans-serif; }
  .dialog { width:440px; background:#2d2d2d; border:1px solid #444; border-radius:8px; padding:24px; }
  .title { color:#f48771; font-size:16px; font-weight:600; margin-bottom:10px; }
  .message { color:#ccc; font-size:13.5px; line-height:1.6; margin-bottom:20px; }
  .button { background:#0078d7; border:0; border-radius:4px; color:#fff; cursor:pointer; padding:8px 18px; }
  .warning { position:fixed; bottom:18px; color:#777; font-size:12px; }
</style>
</head>
<body>
  <div class="dialog">
    <div class="title">` + html.EscapeString(title) + `</div>
    <div class="message">` + html.EscapeString(message) + `</div>
    <button class="button" onclick="location.reload()">Close</button>
  </div>
  <div class="warning">Simulation only. Press Ctrl + Alt + Win + C to exit.</div>
<script>window.dedosPause=function(){};window.dedosResume=function(){};</script>
</body>
</html>
`
}

var malwareDetectedHTML = problemSimulationHTML("Malware Detected", "Windows Defender detected a potentially unwanted application and blocked it from running.")
var firewallBlockHTML = problemSimulationHTML("Firewall Blocked Connection", "Windows Defender Firewall blocked an application from accepting incoming connections.")
var syncFailedHTML = problemSimulationHTML("Sync Failed", "Your files could not be synchronized. Check your connection and try again.")
var pinBlockedHTML = problemSimulationHTML("PIN Temporarily Blocked", "Too many incorrect PIN attempts were detected. Try again later.")
var passwordExpiredHTML = problemSimulationHTML("Password Expired", "Your password has expired and must be changed before you can continue.")
var serviceFailedHTML = problemSimulationHTML("Service Failed", "A required Windows service failed to start correctly.")
var taskSchedulerErrorHTML = problemSimulationHTML("Task Scheduler Error", "The scheduled task could not be started because the task definition is invalid.")
var environmentVariableHTML = problemSimulationHTML("Environment Variable Error", "The requested environment variable could not be found.")
var groupPolicyHTML = problemSimulationHTML("Group Policy Error", "This operation was blocked by your system administrator's group policy.")
var kernelPanicHTML = problemSimulationHTML("Critical System Error", "The system encountered a critical kernel error and must be restarted.")
