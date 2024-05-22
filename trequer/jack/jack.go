package jack

/*
#cgo pkg-config: jack
#include <jack/jack.h>
#include <jack/midiport.h>
#include <jack/ringbuffer.h>

extern jack_client_t *jack_client_open_go(const char *client_name, jack_options_t options, jack_status_t *status);
extern int jack_set_port_registration_callback_go(jack_client_t *client, uintptr_t arg);
extern int jack_set_client_registration_callback_go(jack_client_t *client, uintptr_t arg);
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type PortRegistrationCallback func()
type ClientRegistrationCallback func()

type callbacks struct {
	portRegistration   cgo.Handle
	clientRegistration cgo.Handle
}

type Client struct {
	handle    *C.jack_client_t
	callbacks callbacks
}

type Port struct {
	handle *C.jack_port_t
}

type Options int
type PortFlags uint
type PortType string

const (
	NullOption Options = Options(C.JackNullOption)
)

const (
	PortIsOutput PortFlags = PortFlags(C.JackPortIsOutput)
	PortIsInput  PortFlags = PortFlags(C.JackPortIsInput)
)

const (
	DefaultMidiType PortType = "8 bit raw midi"
)

func ClientOpen(name string, options Options) (*Client, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var status C.jack_status_t
	cclient := C.jack_client_open_go(cname, (C.jack_options_t)(options), &status)
	if int(status) != 0 {
		return nil, fmt.Errorf("status: %v", int(status))
	}

	return &Client{handle: cclient, callbacks: callbacks{}}, nil
}

func (c *Client) ClientClose() {
	C.jack_client_close(c.handle)
}

func (c *Client) Activate() {
	C.jack_activate(c.handle)
}

func (c *Client) SetPortRegistrationCallback(callback PortRegistrationCallback) error {
	c.callbacks.portRegistration = cgo.NewHandle(callback)
	result := C.jack_set_port_registration_callback_go(c.handle, C.uintptr_t(c.callbacks.portRegistration))
	if result != 0 {
		return fmt.Errorf("error %v", result)
	}
	return nil
}

func (c *Client) SetClientRegistrationCallback(callback ClientRegistrationCallback) error {
	c.callbacks.clientRegistration = cgo.NewHandle(callback)
	result := C.jack_set_port_registration_callback_go(c.handle, C.uintptr_t(c.callbacks.clientRegistration))
	if result != 0 {
		return fmt.Errorf("error %v", result)
	}
	return nil
}

func (c *Client) PortRegister(portName string, portType PortType, flags PortFlags) (*Port, error) {
	cname := C.CString(portName)
	defer C.free(unsafe.Pointer(cname))
	ctype := C.CString(string(portType))
	defer C.free(unsafe.Pointer(ctype))

	port := C.jack_port_register(c.handle, cname, ctype, (C.ulong)(flags), 0)
	if port == nil {
		return nil, fmt.Errorf("failed to register port")
	}
	return &Port{handle: port}, nil

}

func (c *Client) GetPorts(portPattern string, typePattern string, flags PortFlags) []string {
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
	callback := cgo.Handle(arg).Value().(PortRegistrationCallback)
	if callback != nil {
		callback()
	}
}

//export goClientRegistration
func goClientRegistration(name *C.char, reg C.int, arg C.uintptr_t) {
	callback := cgo.Handle(arg).Value().(ClientRegistrationCallback)
	if callback != nil {
		callback()
	}
}
