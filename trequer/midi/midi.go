package midi

/*
#cgo pkg-config: alsa
#include <alsa/asoundlib.h>

snd_seq_client_info_t * midi_snd_seq_client_info_malloc();
snd_seq_port_info_t * midi_snd_seq_port_info_malloc();
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type MidiPort struct {
	Id     string
	Client string
	Port   string
}

type Sequencer struct {
	seq *C.snd_seq_t
}

func NewSequencer() *Sequencer {
	return new(Sequencer)
}

func (s *Sequencer) ListPorts() []MidiPort {
	ports := []MidiPort{}

	name := C.CString("default")
	defer C.free(unsafe.Pointer(name))

	res, err := C.snd_seq_open(&s.seq, name, C.SND_SEQ_OPEN_DUPLEX, 0)
	verify(res, err)

	clientName := C.CString("trequer")
	defer C.free(unsafe.Pointer(clientName))
	res, err = C.snd_seq_set_client_name(s.seq, clientName)
	verify(res, err)

	cinfo := C.midi_snd_seq_client_info_malloc()
	defer C.snd_seq_client_info_free(cinfo)

	pinfo := C.midi_snd_seq_port_info_malloc()
	defer C.snd_seq_port_info_free(pinfo)

	C.snd_seq_client_info_set_client(cinfo, -1)
	for C.snd_seq_query_next_client(s.seq, cinfo) >= 0 {
		client := C.snd_seq_client_info_get_client(cinfo)

		C.snd_seq_port_info_set_client(pinfo, client)
		C.snd_seq_port_info_set_port(pinfo, -1)

		for C.snd_seq_query_next_port(s.seq, pinfo) >= 0 {

			/* port must understand MIDI messages */
			if (C.snd_seq_port_info_get_type(pinfo) &
				C.SND_SEQ_PORT_TYPE_MIDI_GENERIC) == 0 {
				continue
			}

			/* we need both WRITE and SUBS_WRITE */
			if (C.snd_seq_port_info_get_capability(pinfo) &
				(C.SND_SEQ_PORT_CAP_WRITE | C.SND_SEQ_PORT_CAP_SUBS_WRITE)) !=
				(C.SND_SEQ_PORT_CAP_WRITE | C.SND_SEQ_PORT_CAP_SUBS_WRITE) {
				continue
			}

			clientId := C.snd_seq_port_info_get_client(pinfo)
			clientName := C.GoString(C.snd_seq_client_info_get_name(cinfo))
			portId := C.snd_seq_port_info_get_port(pinfo)
			portName := C.GoString(C.snd_seq_port_info_get_name(pinfo))

			ports = append(ports, MidiPort{
				Id:     fmt.Sprintf("%d:%d", clientId, portId),
				Client: clientName,
				Port:   portName})
		}
	}

	return ports
}

func verify(res C.int, err error) {
	if res < 0 {
		println(C.GoString(C.snd_strerror(res)))
	}
	if err != nil {
		println(err)
	}
}
