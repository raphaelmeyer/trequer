package trequer

import (
	"github.com/raphaelmeyer/trequer/midi"
)

type MidiPort struct {
	Address string `json:"address"`
	Client  string `json:"client"`
	Port    string `json:"port"`
}

type Trequer struct {
}

func NewTrequer() *Trequer {
	return &Trequer{}
}

func (t *Trequer) GetMidiPorts() []MidiPort {
	ports := []MidiPort{}

	snd, err := midi.NewSndSeq()
	if err != nil {
		return ports
	}
	defer snd.Close()

	ps, err := snd.ListPorts()
	if err != nil {
		return ports
	}
	for _, port := range ps {
		ports = append(ports, MidiPort{
			Address: port.Address, Client: port.Client, Port: port.Port,
		})
	}

	return ports
}
