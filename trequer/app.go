package main

import (
	"context"

	"github.com/raphaelmeyer/trequer/audiocontrol"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx   context.Context
	audio *audiocontrol.AudioControl
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.audio = audiocontrol.NewAudioControl()
}

func (a *App) domReady(ctx context.Context) {
	a.audio.OnPortsChanged(func() {
		runtime.EventsEmit(a.ctx, "ports-changed")
	})
}

func (a *App) shutdown(ctx context.Context) {
	a.audio.Destroy()
}

func (a *App) ListMidiOutputs() []string {
	if a.audio == nil {
		return []string{}
	}
	return a.audio.ListMidiOutputs()
}
