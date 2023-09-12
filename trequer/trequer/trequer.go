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
			Id: port.Id, Client: port.Client, Port: port.Port,
		})
	}

	return ports
}
