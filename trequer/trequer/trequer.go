package trequer

import (
	"github.com/raphaelmeyer/trequer/midi"
)

type MidiPort struct {
	Id     string `json:"id"`
	Client string `json:"client"`
	Port   string `json:"port"`
}

type Trequer struct {
}

func NewTrequer() *Trequer {
	return &Trequer{}
}

func (t *Trequer) GetMidiPorts() []MidiPort {
	seq := midi.NewSequencer()

	ports := []MidiPort{}
	for _, port := range seq.ListPorts() {
		ports = append(ports, MidiPort{
			Id:     port.Id,
			Client: port.Client,
			Port:   port.Port,
		})
	}

	return ports
}
