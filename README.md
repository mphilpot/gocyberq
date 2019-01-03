# Node Library to parse cyberq xml data and serve a pretty UI.

# Random References
- CyberQ XML & POST Guide: http://dl.dropboxusercontent.com/u/466524/GuruWifi.pdf
- Manual: http://www.thebbqguru.com/NEW-PDF/CyberQ-WiFi.pdf

## Development

### Local Fake CyberQ
Start a fake cyberq server up and running on port 10001.

```
node fake_cyberq.js
```

### Launch the App
Launch the proxy application/web server pointed to the fake cyberq environment.

```
node index.js --cyberq:host=localhost:10001
```
