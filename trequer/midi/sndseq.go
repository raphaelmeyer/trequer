package midi

/*
#cgo pkg-config: alsa
#include <alsa/asoundlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type SndSeq struct {
	handle *C.snd_seq_t
}

func convertError(res C.int) error {
	cstr := C.snd_strerror(res)
	defer C.free(unsafe.Pointer(cstr))
	return fmt.Errorf("Alsa: %v", C.GoString(cstr))
}

func NewSndSeq() (*SndSeq, error) {
	seq := &SndSeq{nil}
	if err := seq.open(); err != nil {
		return nil, err
	}

	return seq, nil
}

func (s *SndSeq) open() error {
	name := C.CString("default")
	defer C.free(unsafe.Pointer(name))

	if res := C.snd_seq_open(&s.handle, name, C.SND_SEQ_OPEN_DUPLEX, 0); res != 0 {
		return convertError(res)
	}

	clientName := C.CString("trequer")
	defer C.free(unsafe.Pointer(clientName))

	if res := C.snd_seq_set_client_name(s.handle, clientName); res != 0 {
		s.Close()
		return convertError(res)
	}

	return nil
}

func (s *SndSeq) Close() {
	C.snd_seq_close(s.handle)
	s.handle = nil
}

func (s *SndSeq) ListPorts() ([]MidiPort, error) {
	ports := []MidiPort{}

	var cinfo *C.snd_seq_client_info_t
	if res := C.snd_seq_client_info_malloc(&cinfo); res != 0 {
		return nil, convertError(res)
	}
	defer C.snd_seq_client_info_free(cinfo)

	var pinfo *C.snd_seq_port_info_t
	if res := C.snd_seq_port_info_malloc(&pinfo); res != 0 {
		return nil, convertError(res)
	}
	defer C.snd_seq_port_info_free(pinfo)

	C.snd_seq_client_info_set_client(cinfo, -1)
	for C.snd_seq_query_next_client(s.handle, cinfo) >= 0 {
		client := C.snd_seq_client_info_get_client(cinfo)

		C.snd_seq_port_info_set_client(pinfo, client)
		C.snd_seq_port_info_set_port(pinfo, -1)

		for C.snd_seq_query_next_port(s.handle, pinfo) >= 0 {
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

	return ports, nil

}
