

setup: wails

wails: _go/bin/wails

##

_go/bin/wails:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

