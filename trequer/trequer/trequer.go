package trequer

type MidiPort struct {
	Id     string `json:"id"`
	Client string `json:"client"`
	Port   string `json:"port"`
}

type Trequer struct {
	ports []MidiPort
}

func NewTrequer() *Trequer {
	return &Trequer{
		ports: []MidiPort{
			{Id: "14:0", Client: "Midi Through", Port: "Midi Through Port-0"},
			{Id: "128:0", Client: "FLUID Synth (1307)", Port: "Synth input port (1307:0)"},
		},
	}
}

func (t *Trequer) GetMidiPorts() []MidiPort {
	return t.ports
}
