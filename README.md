# Taste reverse proxy

* Taste with [Caddy v2.6.4](https://github.com/caddyserver/caddy/releases/tag/v2.6.4)

## Windows Powershell Scripts
* `ExecutionPolicy` should be set to `RemoteSigned` and unblock `ps1` files
```powershell
# Check
ExecutionPolicy
# Set as RemoteSigned
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned

# Unblock ps1 files
Unblock-File *.ps1
```

## Run
* Windows
```powershell
.\caddy_download.ps1
cd caddy
.\caddy.exe run --config Caddyfile
cd ..

go build -o bin/
cd bin
dummy-server.exe -port [4416|5525|6636]
```
* Linux
```bash
.\caddy_download.sh1
cd caddy
./caddy run --config Caddyfile
cd ..

go build -o bin/
cd bin
./dummy-server.exe -port [4416|5525|6636]
```

## Requests
* Send via [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for VS Code
```http
@uri = http://localhost

### 4416 Hello
GET {{uri}}/api/teacher/hello HTTP/1.1

### 4416 Health
GET {{uri}}/api/teacher/health HTTP/1.1

### 5525 Hello
GET {{uri}}/api/id/hello HTTP/1.1

### 5525 Health
GET {{uri}}/api/id/health HTTP/1.1

### 6636 Hello
GET {{uri}}/api/id/hello HTTP/1.1

### 6636 Health
GET {{uri}}/api/id/health HTTP/1.1
```
