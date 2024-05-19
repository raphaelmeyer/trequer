package audiocontrol

// #cgo pkg-config: jack
// #include <jack/jack.h>
// #include <jack/midiport.h>
// #include <jack/ringbuffer.h>
/*
jack_client_t * jack_client_open_ (const char *client_name, jack_options_t options, jack_status_t *status) {
	return jack_client_open(client_name, options, status);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type AudioControl struct {
	client *C.jack_client_t
}

func NewAudioControl() *AudioControl {
	ac := new(AudioControl)

	clientName := C.CString("Trequer")
	defer C.free(unsafe.Pointer(clientName))
	ac.client = C.jack_client_open_(clientName, C.JackNullOption, nil)

	return ac
}

func (ac *AudioControl) ListMidiOutputs() []string {
	filter := C.CString("midi")
	defer C.free(unsafe.Pointer(filter))

	var ports []string

	cports, _ := C.jack_get_ports(ac.client, nil, filter, C.JackPortIsInput)
	if cports != nil {
		defer C.jack_free(unsafe.Pointer(cports))
		cport := cports
		for *cport != nil {
			name := C.GoString(*cport)

			fmt.Println(name)
			ports = append(ports, name)

			cport = (**C.char)(unsafe.Add(unsafe.Pointer(cport), unsafe.Sizeof(cport)))
		}
	}

	return ports

	// auto ports = jack_get_ports(client, nullptr, "midi", JackPortIsInput);
	// for (char const **port = ports; *port != nullptr; ++port) {
	//   printf("%s\n", *port);
	// }
	// jack_free(ports);
}
