package jack

/*
#cgo pkg-config: jack
#include <jack/jack.h>
#include <jack/midiport.h>
#include <jack/ringbuffer.h>

extern jack_client_t *jack_client_open_go(const char *client_name, jack_options_t options, jack_status_t *status);
extern int jack_set_port_registration_callback_go(jack_client_t *client, uintptr_t arg);
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type Client struct {
	handle  *C.jack_client_t
	channel *C.jack_port_t

	cbHandle  cgo.Handle
	cbHandler CallbackHandler
}

type options int
type portFlags uint
type portType string

const (
	NullOption options = options(C.JackNullOption)
)

const (
	portIsOutput portFlags = portFlags(C.JackPortIsOutput)
	portIsInput  portFlags = portFlags(C.JackPortIsInput)
)

const (
	midiType portType = "8 bit raw midi"
)

func clientOpen(name string, options options) (*Client, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var status C.jack_status_t
	cclient := C.jack_client_open_go(cname, (C.jack_options_t)(options), &status)
	if int(status) != 0 {
		return nil, fmt.Errorf("status: %v", int(status))
	}

	return &Client{handle: cclient}, nil
}

func (c *Client) clientClose() error {
	status := C.jack_client_close(c.handle)
	if int(status) != 0 {
		return fmt.Errorf("status: %v", int(status))
	}
	return nil
}

func (c *Client) activate() error {
	status := C.jack_activate(c.handle)
	if int(status) != 0 {
		return fmt.Errorf("status: %v", int(status))
	}
	return nil
}

func (c *Client) setPortRegistrationCallback(cbHandle cgo.Handle) error {
	result := C.jack_set_port_registration_callback_go(c.handle, C.uintptr_t(cbHandle))
	if result != 0 {
		return fmt.Errorf("error %v", result)
	}
	return nil
}

func (c *Client) portRegister(portName string, portType portType, flags portFlags) (*C.jack_port_t, error) {
	cname := C.CString(portName)
	defer C.free(unsafe.Pointer(cname))
	ctype := C.CString(string(portType))
	defer C.free(unsafe.Pointer(ctype))

	port := C.jack_port_register(c.handle, cname, ctype, (C.ulong)(flags), 0)
	if port == nil {
		return nil, fmt.Errorf("failed to register port")
	}
	return port, nil

}

func (c *Client) getPorts(portPattern string, typePattern string, flags portFlags) []string {
	cportPattern := C.CString(portPattern)
	defer C.free(unsafe.Pointer(cportPattern))

	ctypePattern := C.CString(typePattern)
	defer C.free(unsafe.Pointer(ctypePattern))

	cports := C.jack_get_ports(c.handle, cportPattern, ctypePattern, (C.ulong)(flags))
	if cports == nil {
		return []string{}
	}

	defer C.jack_free(unsafe.Pointer(cports))

	var ports []string
	cport := cports
	for *cport != nil {
		port := C.GoString(*cport)
		ports = append(ports, port)
		cport = (**C.char)(unsafe.Add(unsafe.Pointer(cport), unsafe.Sizeof(cport)))
	}

	return ports
}

//export goPortRegistration
func goPortRegistration(port C.jack_port_id_t, reg C.int, arg C.uintptr_t) {
	client := cgo.Handle(arg).Value().(*Client)
	if client.cbHandler != nil {
		client.cbHandler.OnRegisterPort()
	}
}
