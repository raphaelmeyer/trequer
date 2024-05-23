package main

import (
	"context"
	"log"

	"github.com/raphaelmeyer/trequer/audiocontrol"
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
}

func (a *App) domReady(ctx context.Context) {
	a.audio = audiocontrol.NewAudioControl(a.ctx)
}

func (a *App) shutdown(ctx context.Context) {
	a.audio.Close()
}

func (a *App) ListMidiOut() []string {
	if a.audio == nil {
		return []string{}
	}
	outs, err := a.audio.ListMidiOut()
	if err != nil {
		log.Println(err)
		return []string{}
	}

	return outs
}
