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
	"unsafe"
)

type AudioControl struct {
	client *C.jack_client_t
}

func NewAudioControl() *AudioControl {
	ac := new(AudioControl)

	clientName := C.CString("trequer")
	defer C.free(unsafe.Pointer(clientName))
	ac.client = C.jack_client_open_(clientName, C.JackNullOption, nil)

	return ac
}
