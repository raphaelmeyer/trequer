package audiocontrol

import (
	"context"
	"log"

	"github.com/raphaelmeyer/trequer/jack"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type AudioControl struct {
	client *jack.Client
	ctx    context.Context
}

func NewAudioControl(ctx context.Context) *AudioControl {
	client, err := jack.NewClient("Trequer")
	if err != nil {
		log.Fatalln(err)
	}

	ac := &AudioControl{client: client, ctx: ctx}

	client.SetCallbackHandler(ac)

	return ac
}

func (ac *AudioControl) Close() {
	ac.client.Close()
}

func (ac *AudioControl) ListMidiOut() ([]string, error) {
	return ac.client.ListMidiOut()
}

func (ac *AudioControl) Play() {
	ac.client.Play()
}

func (ac *AudioControl) Stop() {
	ac.client.Stop()
}

func (ac *AudioControl) OnBeat(beat uint32) {
	runtime.EventsEmit(ac.ctx, "on-beat", beat)
}

func (ac *AudioControl) OnRegisterPort() {
	runtime.EventsEmit(ac.ctx, "ports-changed")
}
