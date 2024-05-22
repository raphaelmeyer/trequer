package audiocontrol

import (
	"github.com/raphaelmeyer/trequer/jack"
)

type AudioControl struct {
	client *jack.Client
	out    *jack.Port

	onPortsChanged func()
}

func NewAudioControl() *AudioControl {
	client, err := jack.ClientOpen("Trequer", jack.NullOption)
	if err != nil {
		panic(err)
	}

	out, err := client.PortRegister("midi-out", jack.DefaultMidiType, jack.PortIsOutput)
	if err != nil {
		panic(err)
	}

	ac := &AudioControl{client: client, out: out}

	err = client.SetPortRegistrationCallback(func() {
		if ac.onPortsChanged != nil {
			ac.onPortsChanged()
		}
	})
	if err != nil {
		panic(err)
	}

	client.Activate()

	return ac
}

func (ac *AudioControl) Destroy() {
	ac.client.ClientClose()
}

func (ac *AudioControl) OnPortsChanged(callback func()) {
	ac.onPortsChanged = callback
}

func (ac *AudioControl) ListMidiOutputs() []string {
	return ac.client.GetPorts("", "midi", jack.PortIsInput)
}
