
all: build

build: setup
	cd trequer && wails build

dev: setup
	cd trequer && wails dev

setup: wails

##

wails: _go/bin/wails

##

_go/bin/wails:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

