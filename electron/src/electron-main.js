'use strict'

const { app, Menu, BrowserWindow, dialog } = require('electron');

var log = require('electron-log');

const path = require('path');

const childProcess = require('child_process');

const cwd = require('process').cwd();

// This adds refresh and devtools console keybindings
// Page can refresh with cmd+r, ctrl+r, F5
// Devtools can be toggled with cmd+alt+i, ctrl+shift+i, 
require('electron-debug')({ enabled: true, showDevTools: false });


global.eval = function() { throw new Error('bad!!'); }

const defaultURL = 'http://127.0.0.1:8620/';
let currentURL;

// Force everything localhost, in case of a leak, except for the teller API endpoint. We should also allow access to the
// Comodo certificate authority since teller API server does not send a complete certificate chain and Chromium requires
// an extra download of the intermediate certificate "COMODO RSA Domain Validation Secure Server CA" in order to
// complete the chain. See https://bugs.chromium.org/p/chromium/issues/detail?id=799722 for more details.
app.commandLine.appendSwitch('host-resolver-rules', 'MAP * 127.0.0.1, EXCLUDE teller.spaco.io, EXCLUDE *.comodoca.com');
app.commandLine.appendSwitch('ssl-version-fallback-min', 'tls1.2');
app.commandLine.appendSwitch('--no-proxy-server');
app.setAsDefaultProtocolClient('spo');



// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let win;

var spo = null;

function startSpo() {
  console.log('Starting spo from electron');

  if (spo) {
    console.log('Spo already running');
    app.emit('spo-ready');
    return
  }

  var reset = () => {
    spo = null;
  }

  // Resolve spo binary location
  var appPath = app.getPath('exe');
  var exe = (() => {
        switch (process.platform) {
  case 'darwin':
    return path.join(appPath, '../../Resources/app/spo');
  case 'win32':
    // Use only the relative path on windows due to short path length
    // limits
    return './resources/app/spo.exe';
  case 'linux':
    return path.join(path.dirname(appPath), './resources/app/spo');
  default:
    return './resources/app/spo';
  }
})()

  var args = [
    '-launch-browser=false',
    '-gui-dir=' + path.dirname(exe),
    '-color-log=false', // must be disabled or web interface detection
    '-logtofile=true',
    // will break
    // broken (automatically generated certs do not work):
    // '-web-interface-https=true',
  ]
  spo = childProcess.spawn(exe, args);

  spo.on('error', (e) => {
    dialog.showErrorBox('Failed to start spo', e.toString());
  app.quit();
});

spo.stdout.on('data', (data) => {
    console.log(data.toString());

  // Scan for the web URL string
  if (currentURL) {
    return
  }
  const marker = 'Starting web interface on ';
  var i = data.indexOf(marker);
  if (i === -1) {
    return
  }
  // var j = data.indexOf('\n', i);

  // // dialog.showErrorBox('index of newline: ', j);
  // if (j === -1) {
  //     throw new Error('web interface url log line incomplete');
  // }
  // var url = data.slice(i + marker.length, j);
  // currentURL = url.toString();
  currentURL = defaultURL;
  app.emit('spo-ready', { url: currentURL });
});

spo.stderr.on('data', (data) => {
    console.log(data.toString());
});

spo.on('close', (code) => {
    console.log('Spo closed');
  reset();
});

spo.on('exit', (code) => {
    console.log('Spo exited');
  reset();
});
}

function createWindow(url) {
  if (!url) {
    url = defaultURL;
  }

  // Create the browser window.
  win = new BrowserWindow({
    width: 1200,
    height: 900,
    title: 'Spo',
    nodeIntegration: false,
    webPreferences: {
      webgl: false,
      webaudio: false,
    },
  });

  // patch out eval
  win.eval = global.eval;

  const ses = win.webContents.session
  ses.clearCache(function () {
    console.log('Cleared the caching of the skycoin wallet.');
  });

  ses.clearStorageData([],function(){
    console.log('Cleared the stored cached data');
  });

  win.loadURL(url);

  // Open the DevTools.
  // win.webContents.openDevTools();

  // Emitted when the window is closed.
  win.on('closed', () => {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    win = null;
});

  // create application's main menu
  var template = [{
    label: "Spaco",
    submenu: [
      { label: "About Spaco", selector: "orderFrontStandardAboutPanel:" },
      { type: "separator" },
      { label: "Quit", accelerator: "Command+Q", click: function() { app.quit(); } }
    ]
  }, {
    label: "Edit",
    submenu: [
      { label: "Undo", accelerator: "CmdOrCtrl+Z", selector: "undo:" },
      { label: "Redo", accelerator: "Shift+CmdOrCtrl+Z", selector: "redo:" },
      { type: "separator" },
      { label: "Cut", accelerator: "CmdOrCtrl+X", selector: "cut:" },
      { label: "Copy", accelerator: "CmdOrCtrl+C", selector: "copy:" },
      { label: "Paste", accelerator: "CmdOrCtrl+V", selector: "paste:" },
      { label: "Select All", accelerator: "CmdOrCtrl+A", selector: "selectAll:" }
    ]
  }];

  Menu.setApplicationMenu(Menu.buildFromTemplate(template));
}

// Enforce single instance
const alreadyRunning = app.makeSingleInstance((commandLine, workingDirectory) => {
      // Someone tried to run a second instance, we should focus our window.
      if (win) {
        if (win.isMinimized()) {
          win.restore();
        }
        win.focus();
      } else {
        createWindow(currentURL || defaultURL);
}
});

if (alreadyRunning) {
  app.quit();
  return;
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', startSpo);

app.on('spo-ready', (e) => {
  createWindow(e.url);
});

// Quit when all windows are closed.
app.on('window-all-closed', () => {
  // On OS X it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
  app.quit();
}
});

app.on('activate', () => {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (win === null) {
  createWindow();
}
});

app.on('will-quit', () => {
  if (spo) {
    spo.kill('SIGINT');
  }
});

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.