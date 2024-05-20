package audiocontrol

import (
	"fmt"

	"github.com/raphaelmeyer/trequer/jack"
)

type AudioControl struct {
	client *jack.Client
}

func NewAudioControl() *AudioControl {
	ac := new(AudioControl)

	ac.client, _ = jack.ClientOpen("Trequer", jack.NullOption)

	ac.client.SetPortRegistrationCallback(func() {
		fmt.Println("PortRegistration")
	})

	ac.client.PortRegister("midi-out", jack.DefaultMidiType, jack.PortIsOutput)

	return ac
}

func (ac *AudioControl) ListMidiOutputs() []string {
	return ac.client.GetPorts("", "midi", jack.PortIsInput)
}
